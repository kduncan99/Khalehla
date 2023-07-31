// Khalehla Project
// Copyright © 2023 by Kurt Duncan, BearSnake LLC
// All Rights Reserved

package ipEngine

import "khalehla/pkg"

const (
	NoTransfer uint = iota
	BasicToBasicTransfer
	BasicToExtendedTransfer
	ExtendedToBasicTransfer
	ExtendedToExtendedTransfer
)

type BankManipulator struct {
	interrupt       pkg.Interrupt
	instructionType int
	operands        []pkg.Word36
	engine          *InstructionEngine

	baseRegisterIndex         uint // Base register to be loaded, determined at New* or in step 10
	gate                      *Gate
	isCallOperation           bool
	isLoadInstruction         bool
	isLXJInstruction          bool
	isReturnOperation         bool
	lxjBankSelector           uint
	lxjInterfaceSpec          uint
	lxjXRegisterIndex         uint
	nextStep                  int
	priorBankDescriptorIndex  uint
	priorBankLevel            uint
	sourceBankDescriptorIndex uint
	sourceBankLevel           uint
	sourceBankDescriptor      *pkg.BankDescriptor
	sourceBankOffset          uint64
	targetBankDescriptor      *pkg.BankDescriptor
	targetBankDescriptorIndex uint
	targetBankLevel           uint
	targetBankOffset          uint64
	transferMode              uint
	returnControlStackFrame   *ReturnControlStackFrame
}

// step1 does a sanity check for a couple of instructions
//
//	returns true if it completed successfully, else false indicating that an interrupt has been posted
//	and processing should be discontinued.
func step1(bm *BankManipulator) bool {
	if bm.instructionType != InvalidInstruction {
		//	post InvalidInstructionInterrupt if B0 or B1 are the target for an LBU instructionType
		if (bm.instructionType == LBUInstruction) &&
			(bm.engine.activityStatePacket.currentInstruction.GetA() < 2) {
			bm.engine.PostInterrupt(pkg.NewInvalidInstructionInterrupt(pkg.InvalidInstructionLBUUsesB0OrB1))
			return false
		}

		//	post AddressingExceptionInterrupt if IS is 3 for LxJ instructions
		if bm.isLXJInstruction && (bm.lxjInterfaceSpec == 3) {
			bm.engine.PostInterrupt(pkg.NewAddressingExceptionInterrupt(pkg.AddressingExceptionInvalidISValue, 0, 0))
			return false
		}
	}

	bm.nextStep++
	return true
}

// step2 retrieves prior L,BDI for any instruction which will result in acquiring a return address/bank
//
//	returns true if it completed successfully, else false indicating that an interrupt has been posted
//	and processing should be discontinued.
func step2(bm *BankManipulator) bool {
	if bm.instructionType == CALLInstruction {
		par := bm.engine.activityStatePacket.programAddressRegister
		bm.priorBankLevel = par.GetLevel()
		bm.priorBankDescriptorIndex = par.GetBankDescriptorIndex()
	} else if bm.isLXJInstruction && (bm.lxjInterfaceSpec < 2) {
		//  We're supposed to be here for normal LxJ and for LxJ/CALL, but we also catch LxJ/GOTO
		//  (interfaceSpec == 1 and target BD is extended with enter access, or gate)...
		//  Because we must do this for IS == 1 and source BD is basic, and it is too early in
		//  the algorithm to know the source BD bank type.
		abtx := uint(0) //	active base table index
		dr := bm.engine.activityStatePacket.designatorRegister

		if bm.instructionType == LBJInstruction {
			abtx = bm.lxjBankSelector + 12
		} else if bm.instructionType == LDJInstruction {
			if dr.basicModeBaseRegisterSelection {
				abtx = 15
			} else {
				abtx = 14
			}
		} else { // LIJInstruction
			if dr.basicModeBaseRegisterSelection {
				abtx = 13
			} else {
				abtx = 12
			}
		}

		abte := bm.engine.GetActiveBaseTableEntry(abtx)
		bm.priorBankLevel = abte.bankLevel
		bm.priorBankDescriptorIndex = abte.bankDescriptorIndex
	}

	bm.nextStep++
	return true
}

// step3 determines source level, BDI, and offset.
//
//	For transfers, this is the address to which we jump.
//	For loads, L,BDI is the bank and offset is a subset.
//	(mostly... there's actually no requirement for LxJ instructions (for example) to jump to the bank which they base).
//	This is a little tricky as the algorithm seemingly requires us to know (in some cases) whether the
//	target bank is extended or basic mode, and we've not yet begun to determine the target bank.
//	In point of fact, the decision tree here does not actually require knowledge of the target bank type.
//
//	returns true if it completed successfully, else false indicating that an interrupt has been posted
//	and processing should be discontinued.
func step3(bm *BankManipulator) bool {
	if bm.interrupt != nil {
		//  source L,BDI,Offset comes from the interrupt vector...
		//  The bank described by B16 begins with 64 contiguous words, indexed by interrupt class (of which there are 64).
		//  Each word is a Program Address Register word, containing the L,BDI,Offset of the interrupt handling routine
		//  Make sure B16 is valid before dereferencing through it.
		bReg := bm.engine.GetBaseRegister(L0BDTBaseRegister)
		if bReg.IsVoid() {
			bm.engine.Stop(L0BaseRegisterInvalidStop, 0)
			return false
		}

		//  intOffset is the offset from the start of the level 0 BDT, to the vector we're interested in.
		bdtLevel0 := bm.engine.GetBaseRegister(L0BDTBaseRegister).GetStorage()
		intOffset := bm.interrupt.GetClass()
		if intOffset >= uint(len(bdtLevel0)) {
			bm.engine.Stop(InterruptHandlerOffsetOutOfRangeStop, 0)
			return false
		}

		lbdiOffset := bdtLevel0[intOffset]
		bm.sourceBankLevel = uint(lbdiOffset >> 33)
		bm.sourceBankDescriptorIndex = uint(lbdiOffset>>18) & 077777
		bm.sourceBankOffset = uint64(lbdiOffset) & 0777777
	} else if bm.instructionType == URInstruction {
		//  source L,BDI comes from operand L,BDI
		//  offset comes from operand.PAR.PC
		bm.sourceBankLevel = uint(bm.operands[0] >> 33)
		bm.sourceBankDescriptorIndex = uint(bm.operands[0]>>18) & 077777
		bm.sourceBankOffset = uint64(bm.operands[0] & 0777777)
	} else if bm.isReturnOperation {
		//  source L,BDI,Offset comes from RCS L,BDI and offset fields
		//  This is where we pop an RCS frame and grab the relevant fields therefrom.
		rcsBReg := bm.engine.GetBaseRegister(RCSBaseRegister)
		if rcsBReg.IsVoid() {
			i := pkg.NewRCSGenericStackUnderOverflowInterrupt(pkg.RCSGenericStackOverflow, RCSBaseRegister, 0)
			bm.engine.PostInterrupt(i)
			return false
		}

		rcsXReg := bm.engine.GetExecOrUserXRegister(RCSIndexRegister)
		if rcsXReg.GetXM() > rcsBReg.GetUpperLimitNormalized() {
			i := pkg.NewRCSGenericStackUnderOverflowInterrupt(pkg.RCSGenericStackUnderflow, RCSIndexRegister, uint(rcsXReg.GetXM()))
			bm.engine.PostInterrupt(i)
			return false
		}

		framePointer := rcsXReg.GetXM()
		offset := framePointer - rcsBReg.GetLowerLimitNormalized()
		frame := rcsBReg.GetStorage()[offset : offset+2]
		bm.returnControlStackFrame = NewReturnControlStackFrameFromBuffer(frame)
		rcsXReg.SetXM(framePointer + 2)

		bm.sourceBankLevel = bm.returnControlStackFrame.bankLevel
		bm.sourceBankDescriptorIndex = bm.returnControlStackFrame.bankDescriptorIndex
		bm.sourceBankOffset = uint64(bm.returnControlStackFrame.offset)
	} else if bm.isLXJInstruction {
		//  source L,BDI comes from basic mode X(a) E,LS,BDI
		//  offset comes from operand
		bmSpec := bm.engine.GetExecOrUserXRegister(bm.lxjXRegisterIndex).GetW()
		execFlag := (bmSpec & 0_400000_000000) != 0
		levelSpec := (bmSpec & 0_040000_000000) != 0
		if execFlag {
			if levelSpec {
				bm.sourceBankLevel = 0
			} else {
				bm.sourceBankLevel = 2
			}
		} else {
			if levelSpec {
				bm.sourceBankLevel = 6
			} else {
				bm.sourceBankLevel = 4
			}
		}
		bm.sourceBankDescriptorIndex = uint((bmSpec >> 18) & 077777)
		bm.sourceBankOffset = uint64(bm.operands[0] & 0777777)
	} else {
		//  source L,BDI,Offset comes from operand
		bm.sourceBankLevel = uint(bm.operands[0]>>33) & 07
		bm.sourceBankDescriptorIndex = uint(bm.operands[0]>>18) & 077777
		bm.sourceBankOffset = uint64(bm.operands[0] & 0777777)
	}

	bm.nextStep++
	return true
}

// step4 ensures L,BDI is valid.  If L,BDI is in the range of 0,1:0,31 we throw an AddressingException.
// If we are handling an interrupt, we stop the ipEngine instead of throwing, and discard further processing.
//
//	returns true if it completed successfully, else false indicating that an interrupt has been posted
//	and processing should be discontinued.
func step4(bm *BankManipulator) bool {
	if (bm.sourceBankLevel == 0) &&
		(bm.sourceBankDescriptorIndex > 0) &&
		(bm.sourceBankDescriptorIndex < 32) {
		if bm.interrupt != nil {
			detail := pkg.Word36((bm.sourceBankLevel << 15) | (bm.sourceBankDescriptorIndex))
			bm.engine.Stop(InterruptHandlerInvalidLevelBDIStop, detail)
			return false
		} else {
			i := pkg.NewAddressingExceptionInterrupt(pkg.AddressingExceptionInvalidSourceLBDI, bm.sourceBankLevel, bm.sourceBankDescriptorIndex)
			bm.engine.PostInterrupt(i)
			return false
		}
	}

	bm.nextStep++
	return true
}

// step5 does void bank handling.
// IF void:
//
//	For loads, we skip to step 10.
//	For interrupt handling we stop the ipEngine
//	For transfers, we either throw an addressing exception or skip to step 10.
//
//	returns true if it completed successfully, else false indicating that an interrupt has been posted
//	and processing should be discontinued.
func step5(bm *BankManipulator) bool {
	if (bm.sourceBankLevel == 0) && (bm.sourceBankDescriptorIndex == 0) {
		if bm.interrupt != nil {
			bm.engine.Stop(InterruptHandlerInvalidLevelBDIStop, 0)
			return false
		} else if bm.isLoadInstruction {
			bm.nextStep = 10
			return true
		} else if bm.isReturnOperation {
			//	pull basic mode enabled flag from rcs frame DB12-17 field
			if bm.returnControlStackFrame.designatorRegister.basicModeEnabled {
				//  return to basic mode - void bank
				bm.nextStep = 10
				return true
			} else {
				//  return to extended mode - addressing exception
				i := pkg.NewAddressingExceptionInterrupt(pkg.AddressingExceptionInvalidSourceLBDI, bm.sourceBankLevel, bm.sourceBankDescriptorIndex)
				bm.engine.PostInterrupt(i)
				return false
			}
		} else if bm.instructionType == URInstruction {
			drReturn := DesignatorRegister{}
			drReturn.SetComposite(bm.operands[1].GetW())
			if drReturn.basicModeEnabled {
				//  return to extended mode - addressing exception
				i := pkg.NewAddressingExceptionInterrupt(pkg.AddressingExceptionInvalidSourceLBDI, bm.sourceBankLevel, bm.sourceBankDescriptorIndex)
				bm.engine.PostInterrupt(i)
				return false
			} else {
				//  return to basic mode - void bank
				bm.nextStep = 10
				return true
			}
		}
	}

	bm.nextStep++
	return true
}

// step6 retrieves the bank descriptor corresponding to the source L,BDI which, at this point, is greater than 0,31.
//
//	returns true if it completed successfully, else false indicating that an interrupt has been posted
//	and processing should be discontinued.
func step6(bm *BankManipulator) bool {
	sbd, ok := bm.engine.findBankDescriptor(bm.sourceBankLevel, bm.sourceBankDescriptorIndex)
	if !ok {
		//	this is worse than the already-posted interrupt... we need to STOP now
		bm.engine.Stop(InterruptHandlerInvalidLevelBDIStop, pkg.Word36(bm.sourceBankLevel<<15|bm.sourceBankDescriptorIndex))
		return false
	}
	bm.sourceBankDescriptor = sbd
	bm.nextStep++
	return true
}

// step7 examines the source bank type to determine what should be done next
//
//	returns true if it completed successfully, else false indicating that an interrupt has been posted
//	and processing should be discontinued.
func step7(bm *BankManipulator) bool {
	if bm.sourceBankDescriptor.GetBankType() == pkg.ExtendedModeBankDescriptor {
		//	In all cases, drop through
	} else if bm.sourceBankDescriptor.GetBankType() == pkg.BasicModeBankDescriptor {
		//  Per PRM, interrupt processing always transfers to B0...
		//  implying that the interrupt handler must be extended mode.
		if bm.interrupt != nil {
			bm.engine.Stop(
				InterruptHandlerInvalidBankTypeStop,
				pkg.Word36((bm.sourceBankLevel<<15)|bm.sourceBankDescriptorIndex))
			return false
		} else if bm.instructionType == LBUInstruction &&
			(bm.engine.activityStatePacket.designatorRegister.processorPrivilege > 1) &&
			!bm.sourceBankDescriptor.GetGeneralAccessPermissions().CanEnter() &&
			!bm.sourceBankDescriptor.GetSpecialAccessPermissions().CanEnter() {
			bm.targetBankDescriptor = nil
		} else if ((bm.instructionType == RTNInstruction) || bm.isLXJInstruction) &&
			!bm.returnControlStackFrame.designatorRegister.basicModeEnabled {
			bm.engine.PostInterrupt(pkg.NewAddressingExceptionInterrupt(pkg.AddressingExceptionBDTypeInvalid, bm.sourceBankLevel, bm.sourceBankDescriptorIndex))
			return false
		}
	} else if bm.sourceBankDescriptor.GetBankType() == pkg.GateBankDescriptor {
		if bm.interrupt != nil {
			bm.engine.Stop(
				InterruptHandlerInvalidBankTypeStop,
				pkg.Word36((bm.sourceBankLevel<<15)|bm.sourceBankDescriptorIndex))
			return false
		} else if bm.isCallOperation || (bm.instructionType == GOTOInstruction) {
			bm.nextStep = 9
		} else if bm.isReturnOperation || (bm.instructionType == URInstruction) {
			bm.engine.PostInterrupt(pkg.NewAddressingExceptionInterrupt(pkg.AddressingExceptionBDTypeInvalid, bm.sourceBankLevel, bm.sourceBankDescriptorIndex))
			return false
		}
	} else if bm.sourceBankDescriptor.GetBankType() == pkg.IndirectBankDescriptor {
		if bm.interrupt != nil {
			bm.engine.Stop(
				InterruptHandlerInvalidBankTypeStop,
				pkg.Word36((bm.sourceBankLevel<<15)|bm.sourceBankDescriptorIndex))
			return false
		} else if bm.isCallOperation || bm.isLoadInstruction {
			bm.nextStep = 8
		} else if bm.isReturnOperation || bm.instructionType == LAEInstruction || bm.instructionType == URInstruction {
			bm.engine.PostInterrupt(pkg.NewAddressingExceptionInterrupt(pkg.AddressingExceptionBDTypeInvalid, bm.sourceBankLevel, bm.sourceBankDescriptorIndex))
			return false
		}
	} else {
		//	Undefined (or unimplemented) bank types produce addressing exception interrupts
		bm.engine.PostInterrupt(pkg.NewAddressingExceptionInterrupt(pkg.AddressingExceptionBDTypeInvalid, bm.sourceBankLevel, bm.sourceBankDescriptorIndex))
		return false
	}

	//	source becomes target, and we go to step 10.
	//	any exceptions to this are handled in the if/then blocks which follow
	bm.targetBankLevel = bm.sourceBankLevel
	bm.targetBankDescriptorIndex = bm.sourceBankDescriptorIndex
	bm.targetBankOffset = bm.sourceBankOffset
	bm.targetBankDescriptor = bm.sourceBankDescriptor
	bm.nextStep = 10

	return true
}

// step8 processes indirect banks
//
//	returns true if it completed successfully, else false indicating that an interrupt has been posted
//	and processing should be discontinued.
func step8(bm *BankManipulator) bool {
	if bm.sourceBankDescriptor.IsGeneralFault() {
		bm.engine.PostInterrupt(pkg.NewAddressingExceptionInterrupt(pkg.AddressingExceptionIndirectGBitSet, bm.sourceBankLevel, bm.sourceBankDescriptorIndex))
		return false
	}

	if (bm.sourceBankLevel == 0) && (bm.sourceBankDescriptorIndex < 32) {
		bm.engine.PostInterrupt(pkg.NewAddressingExceptionInterrupt(pkg.AddressingExceptionInvalidSourceLBDI, bm.sourceBankLevel, bm.sourceBankDescriptorIndex))
		return false
	}

	//	Assume indirected-to bank becomes target, and move on to step 10
	targetLBDI := bm.sourceBankDescriptor.GetIndirectLevelAndBDI()
	bm.targetBankLevel = targetLBDI >> 15
	bm.targetBankDescriptorIndex = targetLBDI & 077777
	bm.targetBankOffset = bm.sourceBankOffset
	ok := false
	bm.targetBankDescriptor, ok = bm.engine.findBankDescriptor(bm.targetBankLevel, bm.targetBankDescriptorIndex)
	if !ok {
		bm.engine.PostInterrupt(pkg.NewAddressingExceptionInterrupt(pkg.AddressingExceptionFatal, bm.targetBankLevel, bm.targetBankDescriptorIndex))
		return false
	}

	bm.nextStep = 10
	if bm.targetBankDescriptor.GetBankType() == pkg.BasicModeBankDescriptor {
		//	When PP>1 and GAP.E == 0 and SAP.E == 0, do void bank (set target bd null)
		if (bm.engine.activityStatePacket.designatorRegister.processorPrivilege > 1) &&
			!bm.targetBankDescriptor.GetGeneralAccessPermissions().CanEnter() &&
			!bm.targetBankDescriptor.GetSpecialAccessPermissions().CanEnter() {
			bm.targetBankDescriptor = nil
		}
	} else if bm.targetBankDescriptor.GetBankType() == pkg.GateBankDescriptor {
		//	Do gate processing?
		if bm.isLXJInstruction || bm.isCallOperation {
			bm.nextStep = 9
		}
	} else {
		bm.engine.PostInterrupt(pkg.NewAddressingExceptionInterrupt(pkg.AddressingExceptionFatal, bm.targetBankLevel, bm.targetBankDescriptorIndex))
		return false
	}

	return true
}

// step9 processes gate banks
//
//	returns true if it completed successfully, else false indicating that an interrupt has been posted
//	and processing should be discontinued.
func step9(bm *BankManipulator) bool {
	if bm.sourceBankDescriptor.IsGeneralFault() {
		bm.engine.PostInterrupt(pkg.NewAddressingExceptionInterrupt(pkg.AddressingExceptionIndirectGBitSet, bm.sourceBankLevel, bm.sourceBankDescriptorIndex))
		return false
	}

	gateBankPerms := bm.sourceBankDescriptor.GetAccessLock().GetEffectivePermissions(
		bm.engine.activityStatePacket.indicatorKeyRegister.accessKey,
		bm.sourceBankDescriptor.GetGeneralAccessPermissions(),
		bm.sourceBankDescriptor.GetSpecialAccessPermissions())
	if !gateBankPerms.CanEnter() {
		bm.engine.PostInterrupt(pkg.NewAddressingExceptionInterrupt(pkg.AddressingExceptionEnterAccessDenied, bm.sourceBankLevel, bm.sourceBankDescriptorIndex))
		return false
	}

	//	Check limits of offset against gate bank to ensure the gate offset is within limits,
	//	and is a multiple of 8 words.
	if (bm.sourceBankOffset < bm.sourceBankDescriptor.GetLowerLimitNormalized()) ||
		(bm.sourceBankOffset > bm.sourceBankDescriptor.GetUpperLimitNormalized()) ||
		(bm.sourceBankOffset&07) != 0 {
		bm.engine.PostInterrupt(pkg.NewAddressingExceptionInterrupt(pkg.AddressingExceptionGateBankBoundaryViolation, bm.sourceBankLevel, bm.sourceBankDescriptorIndex))
		return false
	}

	//	Gate is found at the source offset from the start of the gate bank.
	//	Create gate struct and load it from the packet at the offset.
	gateAddr := bm.sourceBankDescriptor.GetBaseAddress()
	buffer, ok := bm.engine.mainStorage.GetSlice(gateAddr.GetSegment(), gateAddr.GetOffset(), 8)
	if !ok {
		bm.engine.PostInterrupt(pkg.NewAddressingExceptionInterrupt(pkg.AddressingExceptionFatal, bm.sourceBankLevel, bm.sourceBankDescriptorIndex))
		return false
	}
	bm.gate = NewGateFromStorage(buffer)

	//	Compare our key to the gate's lock to ensure we have enter access to the gate
	gatePerms := bm.gate.accessLock.GetEffectivePermissions(
		bm.engine.activityStatePacket.indicatorKeyRegister.accessKey,
		bm.gate.specialAccessPermissions,
		bm.gate.generalAccessPermissions)
	if !gatePerms.CanEnter() {
		bm.engine.PostInterrupt(pkg.NewAddressingExceptionInterrupt(pkg.AddressingExceptionEnterAccessDenied, bm.sourceBankLevel, bm.sourceBankDescriptorIndex))
		return false
	}

	//	If GOTO or (LxJ with X(a).IS ==1 and Gate.GI is set), post GOTO Inhibit interrupt
	if (bm.instructionType == GOTOInstruction) ||
		(bm.isLXJInstruction && (bm.lxjInterfaceSpec == 1)) {
		if bm.gate.gotoInhibit {
			bm.engine.PostInterrupt(pkg.NewAddressingExceptionInterrupt(pkg.AddressingExceptionGateGBitSet, bm.sourceBankLevel, bm.sourceBankDescriptorIndex))
			return false
		}
	}

	//	If target L,BDI is less than 0,32, post interrupt
	if (bm.gate.targetLevel == 0) && (bm.gate.targetBDI < 32) {
		bm.engine.PostInterrupt(pkg.NewAddressingExceptionInterrupt(pkg.AddressingExceptionFatal, bm.sourceBankLevel, bm.sourceBankDescriptorIndex))
		return false
	}

	//	This is where we would do library gate processing. If we did. But we don't.

	//	Fetch target BD
	bm.targetBankLevel = bm.gate.targetLevel
	bm.targetBankDescriptorIndex = bm.gate.targetBDI
	bm.targetBankOffset = uint64(bm.gate.targetOffset)
	bm.targetBankDescriptor, ok = bm.engine.findBankDescriptor(bm.targetBankLevel, bm.targetBankDescriptorIndex)
	if !ok {
		bm.engine.PostInterrupt(pkg.NewAddressingExceptionInterrupt(pkg.AddressingExceptionFatal, bm.sourceBankLevel, bm.sourceBankDescriptorIndex))
		return false
	}

	bm.nextStep = 10
	return true
}

// step10 defines the source and destination execution modes for transfers, then determines the base register to be loaded
//
//	returns true if it completed successfully, else false indicating that an interrupt has been posted
//	and processing should be discontinued.
func step10(bm *BankManipulator) bool {
	bm.nextStep++

	if bm.instructionType == LAEInstruction {
		//	base register index was established a long time ago,
		//	and the execution instruction will load all of B1-B15
		bm.nextStep = 18
	} else if bm.instructionType == LBEInstruction {
		bm.baseRegisterIndex = uint(bm.engine.activityStatePacket.currentInstruction.GetA()) + 16
		bm.nextStep = 18
	} else if bm.instructionType == LBUInstruction {
		bm.baseRegisterIndex = uint(bm.engine.activityStatePacket.currentInstruction.GetA())
		bm.nextStep = 18
	} else if bm.instructionType == URInstruction {
		bm.baseRegisterIndex = 0
		bm.nextStep = 16
	} else if bm.interrupt != nil {
		//	Per PRM, interrupts are always to B0
		bm.baseRegisterIndex = 0
		bm.nextStep = 16
	} else {
		//	So what is left, are transfer operations. What kind of transfer are we doing?
		var destModeBasic bool
		sourceModeBasic := bm.engine.activityStatePacket.designatorRegister.basicModeEnabled
		if bm.isReturnOperation {
			//	Destination mode is defined by DB16 in the RCS frame
			destModeBasic = bm.returnControlStackFrame.designatorRegister.basicModeEnabled
		} else {
			//	call or GOTO op - destination mode is defined by target bank type
			if bm.targetBankDescriptor == nil {
				destModeBasic = sourceModeBasic
			} else {
				destModeBasic = bm.targetBankDescriptor.GetBankType() == pkg.BasicModeBankDescriptor
			}
		}

		bm.baseRegisterIndex = 0
		if destModeBasic {
			if sourceModeBasic {
				bm.transferMode = BasicToBasicTransfer
				if bm.isReturnOperation {
					bm.baseRegisterIndex = bm.returnControlStackFrame.basicModeBaseRegister + 12
				} else if bm.instructionType == LBJInstruction {
					bm.baseRegisterIndex = bm.lxjBankSelector + 12
				} else if bm.instructionType == LDJInstruction {
					if bm.engine.activityStatePacket.designatorRegister.basicModeBaseRegisterSelection {
						bm.baseRegisterIndex = 15
					} else {
						bm.baseRegisterIndex = 14
					}
				} else if bm.instructionType == LIJInstruction {
					if bm.engine.activityStatePacket.designatorRegister.basicModeBaseRegisterSelection {
						bm.baseRegisterIndex = 13
					} else {
						bm.baseRegisterIndex = 12
					}
				}
			} else {
				bm.transferMode = ExtendedToBasicTransfer
				if bm.isReturnOperation {
					bm.baseRegisterIndex = bm.returnControlStackFrame.basicModeBaseRegister + 12
				} else {
					if bm.gate == nil {
						bm.baseRegisterIndex = 12
					} else {
						bm.baseRegisterIndex = bm.gate.targetBDI + 12
					}
				}
			}
		} else {
			if sourceModeBasic {
				bm.transferMode = BasicToExtendedTransfer
			} else {
				bm.transferMode = ExtendedToExtendedTransfer
			}
		}
	}

	return true
}

// step11 deals with prior bank - only for transfers.
//
//		For EM to EM or BM to BM, we do nothing.
//		For EM to BM, set B0 to void base register, and PAR.L,BDI to 0,0
//		For BM to EM LxJ/GOTO and LxJ/CALL B(_baseRegisterIndex).V is set,
//	              LxJ/RTN B(RCS.B+12).V is set
//
//	returns true if it completed successfully, else false indicating that an interrupt has been posted
//	and processing should be discontinued.
func step11(bm *BankManipulator) bool {
	if bm.transferMode == ExtendedToBasicTransfer {
		bm.engine.SetBaseRegister(0, pkg.NewVoidBaseRegister())
		par := bm.engine.activityStatePacket.programAddressRegister
		par.SetLevel(0)
		par.SetBankDescriptorIndex(0)

	} else if bm.transferMode == BasicToBasicTransfer {
		bm.engine.SetBaseRegister(bm.baseRegisterIndex, pkg.NewVoidBaseRegister())
	}

	bm.nextStep++
	return true
}

// step12 creates an entry on the RCS, after checking for RCS overflow.
//
// Only executed for transfers.
// returns true if it completed successfully, else false indicating that an interrupt has been posted
// and processing should be discontinued.
func step12(bm *BankManipulator) bool {
	if bm.isCallOperation {
		rcsBReg := bm.engine.baseRegisters[RCSBaseRegister]
		if rcsBReg.IsVoid() {
			bm.engine.PostInterrupt(pkg.NewRCSGenericStackUnderOverflowInterrupt(pkg.RCSGenericStackOverflow, RCSBaseRegister, 0))
			return false
		}

		rcsXReg := (*IndexRegister)(bm.engine.generalRegisterSet.GetRegister(RCSIndexRegister))
		framePointer := uint64(rcsXReg.GetXM()) - 2
		if framePointer < rcsBReg.GetLowerLimitNormalized() {
			bm.engine.PostInterrupt(pkg.NewRCSGenericStackUnderOverflowInterrupt(pkg.RCSGenericStackOverflow, RCSBaseRegister, uint(framePointer)))
			return false
		}

		rtnAddr := bm.engine.activityStatePacket.programAddressRegister.programCounter + 1
		bValue := uint(0) //	basic mode register is this value + 12
		if (bm.transferMode == ExtendedToBasicTransfer) && (bm.gate != nil) {
			bValue = bm.gate.basicModeBaseRegister
		} else if bm.transferMode == BasicToExtendedTransfer {
			if bm.instructionType == LBJInstruction {
				bValue = bm.lxjBankSelector
			} else if bm.instructionType == LDJInstruction {
				if bm.engine.activityStatePacket.designatorRegister.basicModeBaseRegisterSelection {
					bValue = 15
				} else {
					bValue = 14
				}
			} else if bm.instructionType == LIJInstruction {
				if bm.engine.activityStatePacket.designatorRegister.basicModeBaseRegisterSelection {
					bValue = 13
				} else {
					bValue = 12
				}
			}
		}

		rcsFrame := NewReturnControlStackFrameFromComponents(
			bm.priorBankLevel,
			bm.priorBankDescriptorIndex,
			rtnAddr,
			false,
			bValue,
			bm.engine.activityStatePacket.designatorRegister,
			bm.engine.activityStatePacket.indicatorKeyRegister.accessKey)

		offset := framePointer - rcsBReg.GetLowerLimitNormalized()
		buffer := rcsBReg.GetStorage()[offset : offset+2]
		rcsFrame.WriteToBuffer(buffer)

		xReg := (*IndexRegister)(bm.engine.generalRegisterSet.GetRegister(RCSIndexRegister))
		xReg.SetXM(framePointer)
	}

	bm.nextStep++
	return true
}

// step13 updates X(a) or X11
//
// For LxJ normal,
//
//	translate prior L,BDI to E,LS,BDI,
//	BDR field is _baseRegisterIndex & 03,
//	IS is zero,
//	PAR.PC + 1 -> X(18:35)
//
// For CALL to BM, X11.IS is set to 2, remaining fields undefined
//
//	Designator Register DB17 determines whether X(a) is exec or user register
//
// returns true if it completed successfully, else false indicating that an interrupt has been posted
// and processing should be discontinued.
func step13(bm *BankManipulator) bool {
	if bm.isLXJInstruction && (bm.transferMode == BasicToBasicTransfer) {
		parPCNext := bm.engine.activityStatePacket.programAddressRegister.programCounter + 1
		value := pkg.TranslateToBasicMode(bm.priorBankLevel, bm.priorBankDescriptorIndex, parPCNext).GetComposite()
		value |= pkg.Word36(bm.baseRegisterIndex) << 33
		bm.engine.SetExecOrUserXRegister(bm.lxjXRegisterIndex, IndexRegister(value))
	} else if (bm.instructionType == CALLInstruction) && (bm.transferMode == ExtendedToBasicTransfer) {
		bm.engine.SetExecOrUserXRegister(11, 2<<30)
	}

	bm.nextStep++
	return true
}

// step14 update X(0) - not invoked for non-transfers
//
// For certain transfers, User X0 contains DB16 in Bit 0, and AccessKey in Bits 17:35
//
//	EM to EM GOTO, CALL
//	BM to BM LxJ normal
//	EM to BM GOTO, CALL
//	BM to EM LxJ/GOTO, LxJ/CALL
//
//	returns true if it completed successfully, else false indicating that an interrupt has been posted
//	and processing should be discontinued.
func step14(bm *BankManipulator) bool {
	if bm.isCallOperation {
		asp := bm.engine.activityStatePacket

		value := uint(0)
		if asp.designatorRegister.basicModeEnabled {
			value = 0_400000_000000
		}

		key := asp.indicatorKeyRegister.accessKey
		value |= key.GetComposite()
		bm.engine.GetGeneralRegisterSet().SetRegisterValue(X0, pkg.Word36(value))
	}

	bm.nextStep++
	return true
}

// step15 handles Gate fields transfer
//
//	If a gate is processed:
//	    If Gate.DBI is clear, DR.DB12:15 <- Gate.DB12:15, DB17 <- Gate.DB17
//	    If Gate.AKI is clear, Indicator/Key.AccessKey <= Gate.AccessKey
//	    If Gate.LP0I is clear, UR0 or ER0 <- Gate.LatentParameter0
//	    If Gate.LP1I is clear, UR1 or ER1 <- Gate.LatentParameter1
//	    Selection of user/exec register set is controlled by Gate.DB17 if DBI is clear, else by DR.DB17
//	    Move on to step 17 (steps 15 and 16 are mutually exclusive)
//	Else move on to step 16
//
//	returns true if it completed successfully, else false indicating that an interrupt has been posted
//	and processing should be discontinued.
func step15(bm *BankManipulator) bool {
	if bm.gate != nil {
		asp := bm.engine.activityStatePacket

		if !bm.gate.designatorInhibit {
			temp := asp.designatorRegister.GetComposite() & 0_777702_777777
			temp |= bm.gate.designatorRegisterValue.GetComposite() & 0_000075_000000
			asp.designatorRegister.SetComposite(temp)
		}

		if !bm.gate.accessKeyInhibit {
			asp.indicatorKeyRegister.accessKey = bm.gate.newAccessKey
		}

		if !bm.gate.latentParameter0Inhibit {
			bm.engine.SetExecOrUserRRegister(0, pkg.Word36(bm.gate.latentParameterValue0))
		}

		if !bm.gate.latentParameter1Inhibit {
			bm.engine.SetExecOrUserRRegister(1, pkg.Word36(bm.gate.latentParameterValue1))
		}

		bm.nextStep = 17
	} else {
		bm.nextStep++
	}

	return true
}

// step16 updates ASP for certain transfer instructions (not invoked for non-transfers)
//
//	EM to EM    RTN Replace AccessKey and DB12:17 with RCS fields
//	BM to BM    LxJ/RTN as above
//	EM to BM    GOTO, CALL set DB16
//	            RTN AccessKey / DB12:17 as above
//	BM to EM    LxJ/GOTO, LxJ/CALL clear DB16
//	UR          Entire ASP is replaced with operand contents
//	Interrupt   New ASP formed by hardware
//
//	returns true if it completed successfully, else false indicating that an interrupt has been posted
//	and processing should be discontinued.
func step16(bm *BankManipulator) bool {
	bm.nextStep++
	if bm.interrupt != nil {
		//  PAR is loaded from the values in _targetBankLevel, _targetBankDescriptorIndex,
		//      and _targetOffset which were established previously in this algorithm
		//  Designator Register is cleared excepting the following bits:
		//      DB17 (Exec Register Set Selection) set to 1
		//      DB29 (Arithmetic Exception Enable) set to 1
		//      DB1 (Performance Monitoring Counter Enabled) is set to DB2 - not supported here
		//      DB2 (PerfMon Counter Interrupt Control) and DB31 (Basic Mode BaseRegister Selection) are not changed
		//			Since we're not supporting PerfMon, we clear that to zero (by ignoring the original value)
		//      DB6 is set if this is a HardwareCheck interrupt
		//  Indicator/Key register is zeroed out.
		//  Quantum timer is undefined, and the rest of the ASP is not relevant.
		asp := bm.engine.activityStatePacket
		asp.programAddressRegister =
			NewProgramAddressRegister(bm.targetBankLevel, bm.targetBankDescriptorIndex, uint(bm.targetBankOffset))

		dr := &DesignatorRegister{}
		dr.execRegisterSetSelected = true
		dr.arithmeticExceptionEnabled = true
		dr.basicModeEnabled = bm.targetBankDescriptor.GetBankType() == pkg.BasicModeBankDescriptor
		dr.basicModeBaseRegisterSelection = asp.designatorRegister.basicModeBaseRegisterSelection
		dr.faultHandlingInProgress = bm.interrupt.GetClass() == pkg.HardwareCheckInterruptClass
		asp.designatorRegister = dr

		asp.indicatorKeyRegister.clear()
	} else if bm.instructionType == URInstruction {
		//	Entire ASP is loaded form 7 consecutive operand words,
		//	excepting the short status field of the indicator/key register and the interrupt status words.
		bm.engine.activityStatePacket.ReadFromBuffer(bm.operands)
	} else if bm.instructionType == RTNInstruction {
		bm.engine.activityStatePacket.indicatorKeyRegister.accessKey = bm.returnControlStackFrame.accessKey
		dr := bm.engine.activityStatePacket.designatorRegister
		dr.quantumTimerEnabled = bm.returnControlStackFrame.designatorRegister.quantumTimerEnabled
		dr.deferrableInterruptEnabled = bm.returnControlStackFrame.designatorRegister.deferrableInterruptEnabled
		dr.processorPrivilege = bm.returnControlStackFrame.designatorRegister.processorPrivilege
		dr.basicModeEnabled = bm.returnControlStackFrame.designatorRegister.basicModeEnabled
		dr.execRegisterSetSelected = bm.returnControlStackFrame.designatorRegister.execRegisterSetSelected
	} else if ((bm.instructionType == GOTOInstruction) || (bm.instructionType == CALLInstruction)) &&
		(bm.transferMode == ExtendedToBasicTransfer) {
		bm.engine.activityStatePacket.designatorRegister.basicModeEnabled = true
	} else if bm.isLXJInstruction && (bm.transferMode == BasicToExtendedTransfer) {
		bm.engine.activityStatePacket.designatorRegister.basicModeEnabled = false
	}

	return true
}

// step17 copies the offset from step 3 (or step 9, if gated) to PAR.PC for transfer instructions
//
//	returns true if it completed successfully, else false indicating that an interrupt has been posted
//	and processing should be discontinued.
func step17(bm *BankManipulator) bool {
	if bm.transferMode != NoTransfer {
		bm.engine.SetProgramCounter(uint(bm.targetBankOffset), true)
	}

	bm.nextStep++
	return true
}

// step18 updates the hard-held PAR.L,BDI if we loaded into B0
// or the appropriate ABT entry to zero for a void bank, or L,BDI otherwise.
//
//	returns true if it completed successfully, else false indicating that an interrupt has been posted
//	and processing should be discontinued.
func step18(bm *BankManipulator) bool {
	if bm.baseRegisterIndex == 0 {
		//	This is already done for interrupt handling and for the UR instruction
		if (bm.interrupt == nil) && (bm.instructionType != URInstruction) {
			bm.engine.activityStatePacket.programAddressRegister.SetLevel(bm.targetBankLevel)
			bm.engine.activityStatePacket.programAddressRegister.SetBankDescriptorIndex(bm.targetBankDescriptorIndex)
		} else if bm.baseRegisterIndex < 16 {
			if bm.targetBankDescriptor == nil {
				bm.engine.activeBaseTable[bm.baseRegisterIndex].SetComposite(0)
			} else {
				var offset uint64
				if bm.isLoadInstruction {
					offset = bm.targetBankOffset
				}
				bm.engine.activeBaseTable[bm.baseRegisterIndex].SetBankLevel(bm.targetBankLevel)
				bm.engine.activeBaseTable[bm.baseRegisterIndex].SetBankDescriptorIndex(bm.targetBankDescriptorIndex)
				bm.engine.activeBaseTable[bm.baseRegisterIndex].SetSubsetSpecification(uint(offset))
			}
		}
	}

	bm.nextStep++
	return true
}

// step19 loads the appropriate base register
//
//	returns true if it completed successfully, else false indicating that an interrupt has been posted
//	and processing should be discontinued.
func step19(bm *BankManipulator) bool {
	if bm.targetBankDescriptor == nil {
		//  There is no bank descriptor - set up a void base register
		bm.engine.baseRegisters[bm.baseRegisterIndex].MakeVoid()
	} else if bm.isLoadInstruction && (bm.targetBankOffset != 0) {
		//  we have sub-setting info (in targetBankOffset) -- set up a real Base Register with sub-setting.
		storage, _ := bm.engine.mainStorage.GetBlock(bm.targetBankDescriptor.GetBaseAddress().GetSegment())
		bm.engine.baseRegisters[bm.baseRegisterIndex].
			FromBankDescriptorWithSubsetting(bm.targetBankDescriptor, bm.targetBankOffset, storage)
	} else {
		//  A normal non-sub-setting base register - make it so.
		storage, _ := bm.engine.mainStorage.GetBlock(bm.targetBankDescriptor.GetBaseAddress().GetSegment())
		bm.engine.baseRegisters[bm.baseRegisterIndex].FromBankDescriptor(bm.targetBankDescriptor, storage)
	}

	bm.nextStep++
	return true
}

// step20 toggles DB31 on transfers to basic mode
//
//	returns true if it completed successfully, else false indicating that an interrupt has been posted
//	and processing should be discontinued.
func step20(bm *BankManipulator) bool {
	if (bm.transferMode == BasicToBasicTransfer) || (bm.transferMode == ExtendedToBasicTransfer) {
		bm.engine.FindBasicModeBank(bm.targetBankOffset, true)
	}

	bm.nextStep++
	return true
}

// step21 handles final exception checks
// returns true if it completed successfully, else false indicating that an interrupt has been posted
// and processing should be discontinued.
func step21(bm *BankManipulator) bool {
	if bm.targetBankDescriptor != nil {
		if (bm.instructionType == LBEInstruction) ||
			(bm.instructionType == LBUInstruction) ||
			(bm.transferMode != NoTransfer) {
			if bm.targetBankDescriptor.IsGeneralFault() {
				bm.engine.PostInterrupt(pkg.NewAddressingExceptionInterrupt(pkg.AddressingExceptionFatal, bm.targetBankLevel, bm.targetBankDescriptorIndex))
				return false
			}
		}

		perms := bm.targetBankDescriptor.GetAccessLock().GetEffectivePermissions(
			bm.engine.activityStatePacket.indicatorKeyRegister.accessKey,
			bm.targetBankDescriptor.GetGeneralAccessPermissions(),
			bm.targetBankDescriptor.GetSpecialAccessPermissions())

		//	Non RTN transfer to extended mode bank with no enter access,
		//	non-gated (of course - targets of gate banks should always have no enter access)
		if (bm.transferMode != NoTransfer) &&
			(bm.gate == nil) &&
			!bm.isReturnOperation {
			if (bm.transferMode == BasicToExtendedTransfer) || (bm.transferMode == ExtendedToExtendedTransfer) {
				if !perms.CanEnter() {
					bm.engine.PostInterrupt(pkg.NewAddressingExceptionInterrupt(pkg.AddressingExceptionFatal, bm.targetBankLevel, bm.targetBankDescriptorIndex))
					return false
				}
			}
		}

		//  Did we attempt a non-gated transfer to a basic mode bank with enter access denied,
		//  with relative address not set to the lower limit of the target BD?
		if (bm.transferMode != NoTransfer) &&
			(bm.gate == nil) &&
			(bm.targetBankDescriptor.GetBankType() == pkg.BasicModeBankDescriptor) {
			if (!perms.CanEnter()) &&
				(bm.targetBankOffset != bm.targetBankDescriptor.GetLowerLimitNormalized()) {
				bm.engine.PostInterrupt(pkg.NewAddressingExceptionInterrupt(pkg.AddressingExceptionFatal, bm.targetBankLevel, bm.targetBankDescriptorIndex))
				return false
			}
		}

		//  Did we do gated transfer, or non-gated with no enter access, to a basic mode bank,
		//  while the new PAR.PC does not refer to that bank?
		if (bm.transferMode != NoTransfer) && ((bm.gate != nil) || !perms.CanEnter()) {
			if bm.targetBankDescriptor.GetBankType() == pkg.BasicModeBankDescriptor {
				bReg := bm.engine.baseRegisters[bm.baseRegisterIndex]
				relAddr := uint64(bm.engine.activityStatePacket.programAddressRegister.programCounter)
				interrupt := bReg.CheckAccessLimits(relAddr, false)
				if interrupt != nil {
					interrupt = pkg.NewAddressingExceptionInterrupt(pkg.AddressingExceptionFatal, bm.targetBankLevel, bm.targetBankDescriptorIndex)
					bm.engine.PostInterrupt(interrupt)
					return false
				}
			}
		}

		//	Check for RCS.Trap (only if there is an RCS frame)
		if (bm.returnControlStackFrame != nil) && (bm.returnControlStackFrame.trapFlag) {
			bm.engine.PostInterrupt(pkg.NewAddressingExceptionInterrupt(pkg.AddressingExceptionFatal, bm.targetBankLevel, bm.targetBankDescriptorIndex))
			return false
		}
	}

	bm.nextStep = 0
	return true
}

var handlerSteps = map[int]func(*BankManipulator) bool{
	1:  step1,
	2:  step2,
	3:  step3,
	4:  step4,
	5:  step5,
	6:  step6,
	7:  step7,
	8:  step8,
	9:  step9,
	10: step10,
	11: step11,
	12: step12,
	13: step13,
	14: step14,
	15: step15,
	16: step16,
	17: step17,
	18: step18,
	19: step19,
	20: step20,
	21: step21,
}

func NewBankManipulator(e *InstructionEngine, instructionType int, operand pkg.Word36) *BankManipulator {
	return &BankManipulator{
		engine:            e,
		interrupt:         nil,
		instructionType:   instructionType,
		operands:          []pkg.Word36{operand},
		baseRegisterIndex: 0,
	}
}

func NewBankManipulatorForUR(e *InstructionEngine, instructionType int, operands []pkg.Word36) *BankManipulator {
	return &BankManipulator{
		engine:            e,
		interrupt:         nil,
		instructionType:   instructionType,
		operands:          operands,
		baseRegisterIndex: 0,
	}
}

func NewBankManipulatorForLAE(e *InstructionEngine, baseRegisterIndex uint, operand pkg.Word36) *BankManipulator {
	return &BankManipulator{
		engine:            e,
		interrupt:         nil,
		instructionType:   LAEInstruction,
		operands:          []pkg.Word36{operand},
		baseRegisterIndex: baseRegisterIndex,
	}
}

func NewBankManipulatorForInterrupt(e *InstructionEngine, interrupt pkg.Interrupt) *BankManipulator {
	return &BankManipulator{
		engine:            e,
		interrupt:         interrupt,
		instructionType:   InvalidInstruction,
		operands:          []pkg.Word36{},
		baseRegisterIndex: 0,
	}
}

func (bm *BankManipulator) process() {
	bm.targetBankOffset = 0
	bm.transferMode = NoTransfer
	bm.gate = nil

	if bm.interrupt != nil {
		bm.isCallOperation = true
	} else {
		bm.isLoadInstruction = (bm.instructionType == LAEInstruction) ||
			(bm.instructionType == LBEInstruction) ||
			(bm.instructionType == LBUInstruction)
		bm.isLXJInstruction = (bm.instructionType == LBJInstruction) ||
			(bm.instructionType == LDJInstruction) ||
			(bm.instructionType == LIJInstruction)

		bm.lxjXRegisterIndex = uint(bm.engine.activityStatePacket.currentInstruction.GetA())

		if bm.isLXJInstruction {
			xRegister := bm.engine.GetExecOrUserXRegister(bm.lxjXRegisterIndex)
			bm.lxjInterfaceSpec = uint(xRegister.GetW()>>30) & 03
			bm.lxjBankSelector = uint(xRegister.GetW()>>33) & 03
		}

		bm.isCallOperation = (bm.instructionType == CALLInstruction) ||
			(bm.instructionType == LOCLInstruction) ||
			(bm.isLXJInstruction && (bm.lxjInterfaceSpec < 2))

		//  Note that UR is not considered a return operation
		bm.isReturnOperation = (bm.instructionType == RTNInstruction) ||
			(bm.isLXJInstruction && (bm.lxjInterfaceSpec == 2))
	}

	for bm.nextStep != 0 {
		if !handlerSteps[bm.nextStep](bm) {
			break
		}
	}
}
