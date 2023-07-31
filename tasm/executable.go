// Khalehla Project
// tiny assembler
// Copyright © 2023 by Kurt Duncan, BearSnake LLC
// All Rights Reserved

package tasm

import (
	"fmt"
	"khalehla/pkg"
	"strings"
)

type Executable struct {
	//	map of BDIs to the bank for that BDI
	banks map[uint]*Bank

	//	map of BDIs to the base register upon which the bank should be registered at run time.
	//  key is base register index (0 to 15) and the value is the BDI of the bank.
	initiallyBasedBanks map[uint]uint

	//  stuff for setting the designator register
	arithmeticExceptionEnable bool
	baseRegisterSelection     bool
	basicMode                 bool
	execRegisterSet           bool
	exec24BitIndexing         bool
	operationTrapEnable       bool
	processorPrivilege        uint
	quarterWordMode           bool
	startingAddress           uint
}

func (e *Executable) GetBanks() map[uint]*Bank {
	return e.banks
}

func (e *Executable) GetBaseRegisterSelection() bool {
	return e.baseRegisterSelection
}

func (e *Executable) GetInitiallyBasedBanks() map[uint]uint {
	return e.initiallyBasedBanks
}

func (e *Executable) GetProcessorPrivilege() uint {
	return e.processorPrivilege

}
func (e *Executable) GetStartingAddress() uint {
	return e.startingAddress
}

func (e *Executable) IsArithmeticExceptionEnabled() bool {
	return e.arithmeticExceptionEnable
}

func (e *Executable) IsBasicMode() bool {
	return e.basicMode
}

func (e *Executable) IsExecRegisterSetEnabled() bool {
	return e.execRegisterSet
}

func (e *Executable) IsExec24BitIndexingEnabled() bool {
	return e.exec24BitIndexing
}

func (e *Executable) IsOperationTrapEnabled() bool {
	return e.operationTrapEnable
}

func (e *Executable) IsQuarterWordMode() bool {
	return e.quarterWordMode
}

// LinkSimple links the given segments into a single bank, all accessLock allowed, ring/domain == 0.
// the BDI for the bank will be 0_600004 (level 6, BDI 00004)
func (e *Executable) LinkSimple(segments map[int]*Segment) {
	fmt.Printf("\nLink Simple...\n")

	bdi := uint(0_600004)
	e.banks = make(map[uint]*Bank)
	e.initiallyBasedBanks = make(map[uint]uint)

	//	Find the offsets of all the segments relative to the start of the bank
	//	key is the segment number, value is the offset
	offsets := make(map[int]uint)
	var offset uint
	var bankLength uint
	for segmentNumber, segment := range segments {
		offsets[segmentNumber] = offset
		for _, codeBlock := range segment.generatedCode {
			blockLen := uint(len(codeBlock.code))
			offset += blockLen
			bankLength += blockLen
		}
	}

	fmt.Printf("  Segment Table:\n")
	for segmentNumber, offset := range offsets {
		fmt.Printf("    Seg %03o is at offset %08o\n", segmentNumber, offset)
	}

	bankCode := make([]uint64, bankLength)
	lowerLimit := uint(01000)

	//	Resolve undefined references for the segments
	resolved := make(map[string]uint64)
	for segmentNumber, segment := range segments {
		for symbol, offset := range segment.labels {
			//	offset is from the start of the segment -
			//  we need to also include the offset of the segment from the start of the bank,
			//  and the lower limit (base address) of the bank.
			resolved[symbol] = uint64(uint(offset) + offsets[segmentNumber] + lowerLimit)
		}
	}

	fmt.Printf("  References:\n")
	for symbol, value := range resolved {
		fmt.Printf("    %-12s: %012o\n", symbol, value)
	}

	//	Load code one segment at a time (unresolved)
	cx := 0
	for _, segment := range segments {
		for _, codeBlock := range segment.generatedCode {
			for _, code := range codeBlock.code {
				bankCode[cx] = code
				cx++
			}
		}
	}

	//	Now resolve references
	for segNumber, segment := range segments {
		segOffset := offsets[segNumber]
		for _, ref := range segment.references {
			newValue := resolved[strings.ToUpper(ref.symbol)]
			targetIndex := int(segOffset) + ref.offset
			baseValue := bankCode[targetIndex]
			var err error
			bankCode[targetIndex], err = addFractional(baseValue, newValue, ref.startingBit, ref.bitCount)
			if err != nil {
				fmt.Printf("E: BDI:%06o Offset:%012o: %s\n", bdi, targetIndex, err.Error())
			}
		}
	}

	bd := pkg.NewExtendedModeBankDescriptor(
		pkg.NewAccessLock(0, 0),
		pkg.NewAccessPermissions(true, true, true),
		pkg.NewAccessPermissions(true, true, true),
		nil, // this has to be filled in when the bank is loaded
		false,
		lowerLimit,
		bankLength,
		0)
	e.banks[bdi] = &Bank{
		bankDescriptor:      bd,
		bankDescriptorIndex: bdi,
		code:                bankCode,
	}

	e.initiallyBasedBanks[0] = bdi // the bank should be based on B0
	e.startingAddress = 01000      // TODO pull this from .OPT command
}

func addFractional(baseValue uint64, addend2 uint64, startingBit int, bitCount int) (uint64, error) {
	mask := uint64(1<<bitCount) - 1
	shift := 36 - startingBit - bitCount
	shiftedMask := uint64(mask << shift)
	shiftedNotMask := (^shiftedMask) & pkg.NegativeZero

	addend1 := (baseValue & shiftedMask) >> shift
	sum := addend1 + addend2
	if (sum & mask) != sum {
		return 0, fmt.Errorf("value %012o truncated startingBit:%v length:%v", sum, startingBit, bitCount)
	}

	shiftedSum := sum << shift
	return (baseValue & shiftedNotMask) | (shiftedSum & shiftedMask), nil
}

func (e *Executable) Show() {
	for _, bank := range e.banks {
		bd := bank.GetBankDescriptor()
		fmt.Printf("  Bank BDI:%06o  Access:%v  GAP %s  SAP %s  Lower:%012o\n",
			bank.bankDescriptorIndex,
			bd.GetAccessLock().GetString(),
			bd.GetGeneralAccessPermissions().GetString(),
			bd.GetSpecialAccessPermissions().GetString(),
			bd.GetLowerLimitNormalized())
		addr := bd.GetLowerLimitNormalized()
		for cx := 0; cx < len(bank.code); cx++ {
			fmt.Printf("    %08o: %012o\n", addr+uint64(cx), bank.code[cx])
		}
	}
}
