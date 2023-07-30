// Khalehla Project
// tiny assembler
// Copyright © 2023 by Kurt Duncan, BearSnake LLC
// All Rights Reserved

package tasm

import "fmt"

type CodeBlock struct {
	sourceSet     *SourceSet
	lineNumber    int
	sourceItem    *SourceItem
	segmentNumber int
	segmentOffset int
	code          []uint64
	references    []*Reference
	diagnostics   *DiagnosticSet
}

func NewCodeBlock(sourceSet *SourceSet, lineNumber int, segmentNumber int, segmentOffset int) *CodeBlock {
	return &CodeBlock{
		sourceSet:     sourceSet,
		lineNumber:    lineNumber,
		sourceItem:    sourceSet.sourceItems[lineNumber-1],
		segmentNumber: segmentNumber,
		segmentOffset: segmentOffset,
		code:          make([]uint64, 0),
		references:    make([]*Reference, 0),
		diagnostics:   NewDiagnosticSet(),
	}
}

func (cb *CodeBlock) Emit() {
	genStr := ""
	if len(cb.code) > 0 {
		genStr = fmt.Sprintf("%03o:%06o  %012o", cb.segmentNumber, cb.segmentOffset, cb.code[0])
	}

	fmt.Printf("  %24s  %-20s:%6d  %s\n", genStr, cb.sourceSet.name, cb.lineNumber, cb.sourceSet.sourceItems[cb.lineNumber-1].GetString())
	for cx := 1; cx < len(cb.code); cx++ {
		genStr = fmt.Sprintf("%03o:%06o  %012o", cb.segmentNumber, cb.segmentOffset, cb.code[cx])
		fmt.Printf("  %s\n" + genStr)
	}

	for _, dArray := range cb.diagnostics.diagnostics {
		for _, diag := range dArray {
			fmt.Printf("  %s\n", diag.GetString())
		}
	}
}
