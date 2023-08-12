// Khalehla Project
// Copyright © 2023 by Kurt Duncan, BearSnake LLC
// All Rights Reserved

package pkg

import "fmt"

type InterruptClass uint64
type InterruptShortStatus uint64
type InterruptSync uint64
type InterruptPoint uint64

const (
	HardwareDefaultInterruptClass              InterruptClass = 0
	HardwareCheckInterruptClass                InterruptClass = 1
	ReferenceViolationInterruptClass           InterruptClass = 8
	AddressingExceptionInterruptClass          InterruptClass = 9
	TerminalAddressingExceptionInterruptClass  InterruptClass = 10
	RCSGenericStackUnderOverflowInterruptClass InterruptClass = 11
	SignalInterruptClass                       InterruptClass = 12
	TestAndSetInterruptClass                   InterruptClass = 13
	InvalidInstructionInterruptClass           InterruptClass = 14
	ArithmeticExceptionInterruptClass          InterruptClass = 15
	DataExceptionInterruptClass                InterruptClass = 17
	OperationTrapInterruptClass                InterruptClass = 18
	BreakpointInterruptClass                   InterruptClass = 19
	QuantumTimerInterruptClass                 InterruptClass = 20
	SoftwareBreakInterruptClass                InterruptClass = 24
	JumpHistoryFullInterruptClass              InterruptClass = 25
	DayClockInterruptClass                     InterruptClass = 27
	IPLInterruptClass                          InterruptClass = 29
	UPIInitialInterruptClass                   InterruptClass = 30
	UPINormalInterruptClass                    InterruptClass = 31
)

const (
	ReferenceViolationGRS           InterruptShortStatus = 0
	ReferenceViolationStorageLimits InterruptShortStatus = 1
	ReferenceViolationReadAccess    InterruptShortStatus = 2
	ReferenceViolationWriteAccess   InterruptShortStatus = 3
)

const (
	AddressingExceptionFatal                            InterruptShortStatus = 00
	AddressingExceptionGateGBitSet                      InterruptShortStatus = 01
	AddressingExceptionEnterAccessDenied                InterruptShortStatus = 02
	AddressingExceptionInvalidSourceLBDI                InterruptShortStatus = 03
	AddressingExceptionGateBankBoundaryViolation        InterruptShortStatus = 04
	AddressingExceptionInvalidISValue                   InterruptShortStatus = 05
	AddressingExceptionGOTOInhibit                      InterruptShortStatus = 06
	AddressingExceptionGeneralQueuingViolation          InterruptShortStatus = 07
	AddressingExceptionMaxCountEnq                      InterruptShortStatus = 010
	AddressingExceptionIndirectGBitSet                  InterruptShortStatus = 011
	AddressingExceptionInactiveQueueDListEmpty          InterruptShortStatus = 013
	AddressingExceptionUpdateInProgress                 InterruptShortStatus = 014
	AddressingExceptionQueueBankRepositoryFull          InterruptShortStatus = 015
	AddressingExceptionBDTypeInvalid                    InterruptShortStatus = 016
	AddressingExceptionAccessDeniedPosternOrDataExpanse InterruptShortStatus = 017
	//	There are others...
)

const (
	RCSGenericStackOverflow  InterruptShortStatus = 00
	RCSGenericStackUnderflow InterruptShortStatus = 01
)

const (
	InvalidInstructionBadFunctionCode  InterruptShortStatus = 00
	InvalidInstructionX0Linkage        InterruptShortStatus = 00
	InvalidInstructionLBUUsesB0OrB1    InterruptShortStatus = 00
	InvalidInstructionLBUDUsesB0       InterruptShortStatus = 00
	InvalidInstructionBadPP            InterruptShortStatus = 01
	InvalidInstructionEXRInvalidTarget InterruptShortStatus = 03
)

const (
	InterruptBetweenInstruction InterruptPoint = 0
	InterruptMidExecution       InterruptPoint = 1
	InterruptIndirectExecute    InterruptPoint = 2
)

const (
	InterruptSynchronous  InterruptSync = 0
	InterruptAsynchronous InterruptSync = 1
	InterruptBroadcast    InterruptSync = 2
	InterruptPended       InterruptSync = 3
)

var InterruptNames = map[InterruptClass]string{
	HardwareDefaultInterruptClass:              "Hardware Default",
	HardwareCheckInterruptClass:                "Hardware Check",
	ReferenceViolationInterruptClass:           "Reference Violation",
	AddressingExceptionInterruptClass:          "Addressing Exception",
	TerminalAddressingExceptionInterruptClass:  "Terminal Addressing Exception",
	RCSGenericStackUnderOverflowInterruptClass: "RCS Generic Stack Under/Overflow",
	SignalInterruptClass:                       "Signal",
	TestAndSetInterruptClass:                   "Test And Set",
	InvalidInstructionInterruptClass:           "Invalid Instruction",
	ArithmeticExceptionInterruptClass:          "Arithmetic Exception",
	DataExceptionInterruptClass:                "Data Exception",
	OperationTrapInterruptClass:                "Operation Trap",
	BreakpointInterruptClass:                   "Breakpoint",
	QuantumTimerInterruptClass:                 "Quantum Timer",
	SoftwareBreakInterruptClass:                "Software Break",
	JumpHistoryFullInterruptClass:              "Jump History Full",
	DayClockInterruptClass:                     "DayClock",
	IPLInterruptClass:                          "IPL",
	UPIInitialInterruptClass:                   "UPI Initial",
	UPINormalInterruptClass:                    "UPI Normal",
}

type Interrupt interface {
	GetClass() InterruptClass
	GetInterruptPoint() InterruptPoint
	GetShortStatusField() InterruptShortStatus
	GetStatusWord0() Word36
	GetStatusWord1() Word36
	GetSynchrony() InterruptSync
	IsDeferrable() bool
	IsFault() bool
}

// Class 1 Hardware Check ----------------------------------------------------------------------------------------------

type HardwareCheckInterrupt struct {
	absoluteAddress AbsoluteAddress
}

func (i *HardwareCheckInterrupt) GetClass() InterruptClass {
	return HardwareCheckInterruptClass
}

func (i *HardwareCheckInterrupt) GetInterruptPoint() InterruptPoint {
	return InterruptMidExecution
}

func (i *HardwareCheckInterrupt) GetShortStatusField() InterruptShortStatus {
	return 0
}

func (i *HardwareCheckInterrupt) GetStatusWord0() Word36 {
	isw0 := i.absoluteAddress.GetComposite()[0]
	isw0 &= 0_001777_777777 // clear RA, RI, and Res
	return Word36(isw0)
}

func (i *HardwareCheckInterrupt) GetStatusWord1() Word36 {
	return Word36(i.absoluteAddress.GetComposite()[1])
}

func (i *HardwareCheckInterrupt) GetSynchrony() InterruptSync {
	return InterruptSynchronous
}

func (i *HardwareCheckInterrupt) IsDeferrable() bool {
	return false
}

func (i *HardwareCheckInterrupt) IsFault() bool {
	return true
}

func NewHardwareCheckInterrupt(absAddr *AbsoluteAddress) *HardwareCheckInterrupt {
	return &HardwareCheckInterrupt{
		absoluteAddress: *absAddr,
	}
}

// Class 8 Reference Violation -----------------------------------------------------------------------------------------

// ssf values:
//	bits 0-1: Entry Type
//				0: GRS violation with insufficient PP (see 2.3.7)
//					JGD j-field concatenated with a-field is a GRS location
//					SRS, LRS a-field indicates a GRS address
//					All other GRS locations developed as an instructionType operand caused by any instructions
//						other than JGD, SRS, or LRS
//				1: Storage Limits violation
//				2: Read Access violation
//				3: Write Access violation
//  bits 2-4: reserved
//	bits 5: true if this occurred during an instructionType fetch

type ReferenceViolationInterrupt struct {
	shortStatusField InterruptShortStatus
}

func (i *ReferenceViolationInterrupt) GetClass() InterruptClass {
	return ReferenceViolationInterruptClass
}

func (i *ReferenceViolationInterrupt) GetInterruptPoint() InterruptPoint {
	return InterruptMidExecution
}

func (i *ReferenceViolationInterrupt) GetShortStatusField() InterruptShortStatus {
	return i.shortStatusField
}

func (i *ReferenceViolationInterrupt) GetStatusWord0() Word36 {
	return 0
}

func (i *ReferenceViolationInterrupt) GetStatusWord1() Word36 {
	return 0
}

func (i *ReferenceViolationInterrupt) GetSynchrony() InterruptSync {
	return InterruptSynchronous
}

func (i *ReferenceViolationInterrupt) IsDeferrable() bool {
	return false
}

func (i *ReferenceViolationInterrupt) IsFault() bool {
	return true
}

func NewReferenceViolationInterrupt(entryType InterruptShortStatus, fetchOperation bool) *ReferenceViolationInterrupt {
	ssf := (entryType & 03) << 4
	if fetchOperation {
		ssf |= 01
	}
	return &ReferenceViolationInterrupt{
		shortStatusField: ssf,
	}
}

// Class 9 Addressing Exception ----------------------------------------------------------------------------------------

// ssf values:
//	000 Fatal addressing exception
//	001 G-bit set in gate bank descriptor
//	002 Enter access denied by gate bank descriptor or by gate, or queuing instruction access denied
//	003 invalid source L,BDI or BDT limit error for L,BDI supplied by user instruction
//  004 gate bank boundary violation or gate input offset not within gate bd limits
//	005 invalid IS value
//	006 GOTO inhibit set in gate
//	007 General queuing instruction violation
//	010 MaxCount exceeded on ENQ/ENQF
//	011 G-bit set in indirect bank descriptor
//	013 Inactive QBD list empty on DEQ/DEQW
//	014 Update in progress set in queue structure

type AddressingExceptionInterrupt struct {
	shortStatusField     InterruptShortStatus
	interruptStatusWord1 Word36
}

func (i *AddressingExceptionInterrupt) GetClass() InterruptClass {
	return AddressingExceptionInterruptClass
}

func (i *AddressingExceptionInterrupt) GetInterruptPoint() InterruptPoint {
	return InterruptIndirectExecute
}

func (i *AddressingExceptionInterrupt) GetShortStatusField() InterruptShortStatus {
	return i.shortStatusField
}

func (i *AddressingExceptionInterrupt) GetStatusWord0() Word36 {
	return 0
}

func (i *AddressingExceptionInterrupt) GetStatusWord1() Word36 {
	return i.interruptStatusWord1
}

func (i *AddressingExceptionInterrupt) GetSynchrony() InterruptSync {
	return InterruptSynchronous
}

func (i *AddressingExceptionInterrupt) IsDeferrable() bool {
	return false
}

func (i *AddressingExceptionInterrupt) IsFault() bool {
	return true
}

func NewAddressingExceptionInterrupt(
	shortStatusField InterruptShortStatus,
	sourceBankLevel uint64,
	sourceBankDescriptorIndex uint64) *AddressingExceptionInterrupt {

	isw1 := Word36(sourceBankLevel&07) << 33
	isw1 |= Word36(sourceBankDescriptorIndex&077777) << 18
	return &AddressingExceptionInterrupt{
		shortStatusField:     shortStatusField,
		interruptStatusWord1: isw1,
	}
}

// Class 11 RCS/Generic Stack Under/Overflow ---------------------------------------------------------------------------

// ssf values:
//
//	0 Generic stack or RCS overflow
//	1 Generic stack or RCS underrflow
//
// ISW0:
//	Bits 0-5 (S1): BREG (base register causing trouble) - when the RCS causes the interrupt, BREG will be 25
//  Bits 12-35:    Relative address (n/a for BREG 25)
//                  When BREG != 25 and ssf == 0, this field contains Xm - Xi - d of the X register specified
//                      by the BUY instruction
//                  When BREG != 25 and ssf == 1, this field contains Xm of the X register specified
//                      by the SELL instruction

type RCSGenericStackUnderOverflowInterrupt struct {
	shortStatusField     InterruptShortStatus
	interruptStatusWord0 Word36
}

func (i *RCSGenericStackUnderOverflowInterrupt) GetClass() InterruptClass {
	return RCSGenericStackUnderOverflowInterruptClass
}

func (i *RCSGenericStackUnderOverflowInterrupt) GetInterruptPoint() InterruptPoint {
	return InterruptIndirectExecute
}

func (i *RCSGenericStackUnderOverflowInterrupt) GetShortStatusField() InterruptShortStatus {
	return i.shortStatusField
}

func (i *RCSGenericStackUnderOverflowInterrupt) GetStatusWord0() Word36 {
	return i.interruptStatusWord0
}

func (i *RCSGenericStackUnderOverflowInterrupt) GetStatusWord1() Word36 {
	return 0
}

func (i *RCSGenericStackUnderOverflowInterrupt) GetSynchrony() InterruptSync {
	return InterruptSynchronous
}

func (i *RCSGenericStackUnderOverflowInterrupt) IsDeferrable() bool {
	return false
}

func (i *RCSGenericStackUnderOverflowInterrupt) IsFault() bool {
	return true
}

func NewRCSGenericStackUnderOverflowInterrupt(
	shortStatusField InterruptShortStatus,
	baseRegister uint64,
	relativeAddress uint64) *RCSGenericStackUnderOverflowInterrupt {

	isw0 := (Word36(baseRegister) << 30) | Word36(relativeAddress)
	return &RCSGenericStackUnderOverflowInterrupt{
		shortStatusField:     shortStatusField,
		interruptStatusWord0: isw0,
	}
}

// Class 14 Invalid Instruction ----------------------------------------------------------------------------------------

// ssf values:
//
//	0 function code not defined, direct execution or as a target of EXR
//		or LBJ/LIJ/LDJ uses X0
//		or LBU uses B0 or B1
//		or LBUD uses B0
//	1 insufficient processor privilege
//	3 EXR target invalid (other than as above for value 0)
//	4 compatibility trap (we don't do this)

type InvalidInstructionInterrupt struct {
	shortStatusField InterruptShortStatus
}

func (i *InvalidInstructionInterrupt) GetClass() InterruptClass {
	return InvalidInstructionInterruptClass
}

func (i *InvalidInstructionInterrupt) GetInterruptPoint() InterruptPoint {
	return InterruptIndirectExecute
}

func (i *InvalidInstructionInterrupt) GetShortStatusField() InterruptShortStatus {
	return i.shortStatusField
}

func (i *InvalidInstructionInterrupt) GetStatusWord0() Word36 {
	return 0
}

func (i *InvalidInstructionInterrupt) GetStatusWord1() Word36 {
	return 0
}

func (i *InvalidInstructionInterrupt) GetSynchrony() InterruptSync {
	return InterruptSynchronous
}

func (i *InvalidInstructionInterrupt) IsDeferrable() bool {
	return false
}

func (i *InvalidInstructionInterrupt) IsFault() bool {
	return true
}

func NewInvalidInstructionInterrupt(shortStatusField InterruptShortStatus) *InvalidInstructionInterrupt {
	return &InvalidInstructionInterrupt{
		shortStatusField: shortStatusField,
	}
}

func GetInterruptString(i Interrupt) string {
	return fmt.Sprintf("%s(%03o) SSF:%03o ISW0=%012o ISW1=%012o",
		InterruptNames[i.GetClass()],
		i.GetClass(),
		i.GetShortStatusField(),
		i.GetStatusWord0(),
		i.GetStatusWord1())
}

// Class 19 Breakpoint -------------------------------------------------------------------------------------------------

type BreakpointInterrupt struct{}

func (i *BreakpointInterrupt) GetClass() InterruptClass {
	return BreakpointInterruptClass
}

func (i *BreakpointInterrupt) GetInterruptPoint() InterruptPoint {
	return InterruptBetweenInstruction
}

func (i *BreakpointInterrupt) GetShortStatusField() InterruptShortStatus {
	return 0
}

func (i *BreakpointInterrupt) GetStatusWord0() Word36 {
	return Word36(0)
}

func (i *BreakpointInterrupt) GetStatusWord1() Word36 {
	return Word36(0)
}

func (i *BreakpointInterrupt) GetSynchrony() InterruptSync {
	return InterruptPended
}

func (i *BreakpointInterrupt) IsDeferrable() bool {
	return true
}

func (i *BreakpointInterrupt) IsFault() bool {
	return false
}

func NewBreakpointInterrupt() *BreakpointInterrupt {
	return &BreakpointInterrupt{}
}
