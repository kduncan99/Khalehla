// Khalehla Project
// Copyright © 2023 by Kurt Duncan, BearSnake LLC
// All Rights Reserved

package ipEngine

import (
	"khalehla/tasm"
	"testing"
)

// Unconditional -------------------------------------------------------------------------------------------------------

var lmjBasicMode = []*tasm.SourceItem{
	segSourceItem(0),
	tasm.NewSourceItem("", "fjaxu", []string{fNOPBasic, jNOPBasic, aNOP, zero, zero}),
	tasm.NewSourceItem("", "fjaxu", []string{fNOPBasic, jNOPBasic, aNOP, zero, zero}),
	tasm.NewSourceItem("", "fjaxu", []string{fNOPBasic, jNOPBasic, aNOP, zero, zero}),
	tasm.NewSourceItem("", "fjaxu", []string{fLXM, jU, rX10, zero, "03"}),
	tasm.NewSourceItem("", "fjaxu", []string{fLX, jU, rX11, zero, "0"}),
	tasm.NewSourceItem("", "fjaxhiu", []string{fLMJ, jLMJ, rX11, rX10, zero, zero, "label"}),
	tasm.NewSourceItem("", "fjaxu", []string{fNOPBasic, jNOPBasic, aNOP, zero, zero}),
	tasm.NewSourceItem("", "fjaxu", []string{fNOPBasic, jNOPBasic, aNOP, zero, zero}),
	tasm.NewSourceItem("", "fjaxu", []string{fNOPBasic, jNOPBasic, aNOP, zero, zero}),
	tasm.NewSourceItem("label", "fjaxu", []string{fNOPBasic, jNOPBasic, aNOP, zero, zero}),
	iarSourceItem("", 1),
	iarSourceItem("", 2),
	iarSourceItem("target", 0),
}

func Test_LMJ_Basic(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", lmjBasicMode)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), false)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(true)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStoppedReason(t, engine, InitiateAutoRecoveryStop, 0)
	checkRegister(t, engine, X11, 0_000000_01006, "X11")
}

var sljBasicMode = []*tasm.SourceItem{
	segSourceItem(0),
	tasm.NewSourceItem("", "fjaxhiu", []string{fSLJ, jSLJ, zero, zero, zero, zero, "target"}),
	iarSourceItem("", 1),
	iarSourceItem("", 1),
	iarSourceItem("", 1),
	tasm.NewSourceItem("target", "w", []string{"0"}),
	iarSourceItem("", 0),
}

func Test_SLJ_Basic(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", sljBasicMode)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), false)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(true)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStoppedReason(t, engine, InitiateAutoRecoveryStop, 0)
	codeBankAddr := e.GetBanks()[0600004].GetBankDescriptor().GetBaseAddress()
	checkMemory(t, engine, codeBankAddr, 04, 01001)
}

var lmjExtendedMode = []*tasm.SourceItem{
	segSourceItem(0),
	tasm.NewSourceItem("", "fjaxu", []string{fNOPExtended, jNOPExtended, aNOP, zero, zero}),
	tasm.NewSourceItem("", "fjaxu", []string{fNOPExtended, jNOPExtended, aNOP, zero, zero}),
	tasm.NewSourceItem("", "fjaxu", []string{fNOPExtended, jNOPExtended, aNOP, zero, zero}),
	tasm.NewSourceItem("", "fjaxu", []string{fLXM, jU, rX10, zero, "03"}),
	tasm.NewSourceItem("", "fjaxu", []string{fLX, jU, rX11, zero, "0"}),
	tasm.NewSourceItem("", "fjaxhibd", []string{fLMJ, jLMJ, rX11, rX10, zero, zero, rB0, "label"}),
	tasm.NewSourceItem("", "fjaxu", []string{fNOPExtended, jNOPExtended, aNOP, zero, zero}),
	tasm.NewSourceItem("", "fjaxu", []string{fNOPExtended, jNOPExtended, aNOP, zero, zero}),
	tasm.NewSourceItem("", "fjaxu", []string{fNOPExtended, jNOPExtended, aNOP, zero, zero}),
	tasm.NewSourceItem("label", "fjaxu", []string{fNOPExtended, jNOPExtended, aNOP, zero, zero}),
	iarSourceItem("", 1),
	iarSourceItem("", 2),
	iarSourceItem("target", 0),
}

func Test_LMJ_Extended(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", lmjExtendedMode)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), true)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(false)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStoppedReason(t, engine, InitiateAutoRecoveryStop, 0)
	checkRegister(t, engine, X11, 0_000000_01006, "X11")
}

var jumpBasicMode = []*tasm.SourceItem{
	segSourceItem(0),
	tasm.NewSourceItem("", "fjaxu", []string{fJ, jJBasic, aJBasic, zero, "target"}),
	iarSourceItem("", 1),
	iarSourceItem("target", 0),
}

func Test_J_Basic(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpBasicMode)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), false)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(true)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStopped(t, engine)
	checkProgramAddress(t, engine, 01003)
}

var jumpKeyBasicMode = []*tasm.SourceItem{
	segSourceItem(0),
	tasm.NewSourceItem("", "fjaxu", []string{fJK, jJK, "1", zero, "target"}),
	iarSourceItem("", 0),
	iarSourceItem("target", 1),
}

func Test_JK_Basic(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpKeyBasicMode)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), false)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(true)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStopped(t, engine)
	checkProgramAddress(t, engine, 01002)
}

var jumpExtendedMode = []*tasm.SourceItem{
	segSourceItem(0),
	tasm.NewSourceItem("", "fjaxu", []string{fJ, jJExtended, aJExtended, zero, "target"}),
	iarSourceItem("", 1),
	iarSourceItem("target", 0),
}

func Test_J_Extended(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpExtendedMode)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), true)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(false)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStopped(t, engine)
	checkProgramAddress(t, engine, 01003)
}

var haltKeysAndJumpBasicMode = []*tasm.SourceItem{
	segSourceItem(0),
	tasm.NewSourceItem("", "fjaxu", []string{fHKJ, jHKJ, zero, zero, "target"}),
	iarSourceItem("", 1),
	iarSourceItem("target", 0),
}

func Test_HKJ_Basic(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", haltKeysAndJumpBasicMode)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), false)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(true)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStopped(t, engine)
	checkProgramAddress(t, engine, 01003)
}

var haltJumpExtendedMode = []*tasm.SourceItem{
	segSourceItem(0),
	tasm.NewSourceItem("", "fjaxu", []string{fHLTJ, jHLTJ, aHLTJ, zero, "target"}),
	iarSourceItem("", 1),
	iarSourceItem("target", 0),
}

func Test_HLTJ_Extended(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", haltJumpExtendedMode)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), true)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(false)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStopped(t, engine)
	reason, detail := engine.GetStopReason()
	if reason != HaltJumpExecutedStop {
		t.Fatalf("Processor stopped for wrong reason: %d detail: %012o", reason, detail)
	}

	checkProgramAddress(t, engine, 01002)
}

// Conditional based on register ---------------------------------------------------------------------------------------

var jumpZeroExtendedPosZero = []*tasm.SourceItem{
	segSourceItem(0),
	tasm.NewSourceItem("", "fjaxu", []string{fLA, jU, rA5, zero, "0"}),
	tasm.NewSourceItem("", "fjaxhibd", []string{fJZ, jJZ, rA5, zero, zero, zero, rB0, "target"}),
	iarSourceItem("", 1),
	iarSourceItem("target", 0),
}

func Test_JZ_Extended_PosZero(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpZeroExtendedPosZero)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), true)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(false)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStoppedReason(t, engine, InitiateAutoRecoveryStop, 0)
}

var jumpZeroExtendedNegZero = []*tasm.SourceItem{
	segSourceItem(0),
	tasm.NewSourceItem("", "fjaxu", []string{fLA, jXU, rA5, zero, "0777777"}),
	tasm.NewSourceItem("", "fjaxhibd", []string{fJZ, jJZ, rA5, zero, zero, zero, rB0, "target"}),
	iarSourceItem("", 1),
	iarSourceItem("target", 0),
}

func Test_JZ_Extended_NegZero(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpZeroExtendedNegZero)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), true)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(false)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStoppedReason(t, engine, InitiateAutoRecoveryStop, 0)
}

var jumpZeroExtendedNotZero = []*tasm.SourceItem{
	segSourceItem(0),
	tasm.NewSourceItem("", "fjaxu", []string{fLA, jU, rA5, zero, "01"}),
	tasm.NewSourceItem("", "fjaxhibd", []string{fJZ, jJZ, rA5, zero, zero, zero, rB0, "target"}),
	iarSourceItem("", 1),
	iarSourceItem("target", 0),
}

func Test_JZ_Extended_NotZero(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpZeroExtendedNotZero)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), true)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(false)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStoppedReason(t, engine, InitiateAutoRecoveryStop, 1)
}

var doubleJumpZeroExtendedMode = []*tasm.SourceItem{
	segSourceItem(3),
	tasm.NewSourceItem("posZero", "w", []string{"0"}),
	tasm.NewSourceItem("", "w", []string{"0"}),
	tasm.NewSourceItem("negZero", "w", []string{"0777777777777"}),
	tasm.NewSourceItem("", "w", []string{"0777777777777"}),
	tasm.NewSourceItem("notZero1", "w", []string{"0"}),
	tasm.NewSourceItem("", "w", []string{"011"}),
	tasm.NewSourceItem("notZero2", "w", []string{"0777777777777"}),
	tasm.NewSourceItem("", "w", []string{"0"}),

	segSourceItem(0),
	tasm.NewSourceItem("", "fjaxhibd", []string{fDL, jDL, rA0, zero, zero, zero, rB3, "posZero"}),
	tasm.NewSourceItem("", "fjaxhibd", []string{fDJZ, jDJZ, rA0, zero, zero, zero, rB0, "target1"}),
	iarSourceItem("", 1),

	tasm.NewSourceItem("target1", "fjaxhibd", []string{fDL, jDL, rA2, zero, zero, zero, rB3, "negZero"}),
	tasm.NewSourceItem("", "fjaxhibd", []string{fDJZ, jDJZ, rA2, zero, zero, zero, rB0, "target2"}),
	iarSourceItem("", 2),

	tasm.NewSourceItem("target2", "fjaxhibd", []string{fDL, jDL, rA4, zero, zero, zero, rB3, "notZero1"}),
	tasm.NewSourceItem("", "fjaxhibd", []string{fDJZ, jDJZ, rA4, zero, zero, zero, rB0, "bad3"}),

	tasm.NewSourceItem("target3", "fjaxhibd", []string{fDL, jDL, rA6, zero, zero, zero, rB3, "notZero2"}),
	tasm.NewSourceItem("", "fjaxhibd", []string{fDJZ, jDJZ, rA6, zero, zero, zero, rB0, "bad4"}),
	tasm.NewSourceItem("", "fjaxu", []string{fJ, jJExtended, aJExtended, zero, "end"}),

	iarSourceItem("bad3", 3),
	iarSourceItem("bad4", 4),
	iarSourceItem("end", 0),
}

func Test_DJZ_Extended_PosZero(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", doubleJumpZeroExtendedMode)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkBankPerSegment(a.GetSegments(), true)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(false)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStoppedReason(t, engine, InitiateAutoRecoveryStop, 0)
}

var jumpNonZeroExtendedPosZero = []*tasm.SourceItem{
	segSourceItem(0),
	tasm.NewSourceItem("", "fjaxu", []string{fLA, jU, rA5, zero, "0"}),
	tasm.NewSourceItem("", "fjaxhibd", []string{fJNZ, jJNZ, rA5, zero, zero, zero, rB0, "target"}),
	iarSourceItem("", 1),
	iarSourceItem("target", 0),
}

func Test_JNZ_Extended_PosZero(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpNonZeroExtendedPosZero)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), true)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(false)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStoppedReason(t, engine, InitiateAutoRecoveryStop, 1)
}

var jumpNonZeroExtendedNegZero = []*tasm.SourceItem{
	segSourceItem(0),
	tasm.NewSourceItem("", "fjaxu", []string{fLA, jXU, rA5, zero, "0777777"}),
	tasm.NewSourceItem("", "fjaxhibd", []string{fJNZ, jJNZ, rA5, zero, zero, zero, rB0, "target"}),
	iarSourceItem("", 1),
	iarSourceItem("target", 0),
}

func Test_JNZ_Extended_NegZero(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpNonZeroExtendedNegZero)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), true)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(false)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStoppedReason(t, engine, InitiateAutoRecoveryStop, 1)
}

var jumpNonZeroExtendedNotZero = []*tasm.SourceItem{
	segSourceItem(0),
	tasm.NewSourceItem("", "fjaxu", []string{fLA, jU, rA5, zero, "01"}),
	tasm.NewSourceItem("", "fjaxhibd", []string{fJNZ, jJNZ, rA5, zero, zero, zero, rB0, "target"}),
	iarSourceItem("", 1),
	iarSourceItem("target", 0),
}

func Test_JNZ_Extended_NotZero(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpNonZeroExtendedNotZero)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), true)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(false)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStoppedReason(t, engine, InitiateAutoRecoveryStop, 0)
}

var jumpPosNegExtended = []*tasm.SourceItem{
	segSourceItem(0),
	tasm.NewSourceItem("", "fjaxu", []string{fLA, jU, rA10, zero, "0"}),
	tasm.NewSourceItem("", "fjaxu", []string{fJP, jJP, rA10, zero, "target1"}),
	iarSourceItem("bad1", 1),
	tasm.NewSourceItem("target1", "fjaxu", []string{fJN, jJN, rA10, zero, "bad2"}),
	tasm.NewSourceItem("", "fjaxu", []string{fNOPExtended, jNOPExtended, zero, zero, zero}),

	tasm.NewSourceItem("", "fjaxu", []string{fLA, jXU, rA10, zero, "0444444"}),
	tasm.NewSourceItem("", "fjaxu", []string{fJP, jJP, rA10, zero, "bad3"}),
	tasm.NewSourceItem("", "fjaxu", []string{fJN, jJN, rA10, zero, "end"}),

	iarSourceItem("bad4", 4),
	iarSourceItem("bad2", 2),
	iarSourceItem("bad3", 3),
	iarSourceItem("end", 0),
}

func Test_JP_JN_Extended(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpPosNegExtended)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), true)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(false)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStoppedReason(t, engine, InitiateAutoRecoveryStop, 0)
}

//	TODO JPS, JNS
//	TODO JB, JNB
//	TODO JGD, JMGI

// Conditional based on designator register bits -----------------------------------------------------------------------

var jumpCarryBasic = []*tasm.SourceItem{
	segSourceItem(0),
	tasm.NewSourceItem("", "fjaxu", []string{fJC, jJCBasic, aJCBasic, zero, "target"}),
	iarSourceItem("", 1),
	iarSourceItem("target", 0),
}

func Test_JC_Basic_Pos(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpCarryBasic)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), false)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(true)
		ute.GetEngine().GetDesignatorRegister().SetCarry(true)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStopped(t, engine)
	checkProgramAddress(t, engine, 01003)
}

func Test_JC_Basic_Neg(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpCarryBasic)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), false)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(true)
		ute.GetEngine().GetDesignatorRegister().SetOverflow(false)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStopped(t, engine)
	checkProgramAddress(t, engine, 01002)
}

var jumpCarryExtended = []*tasm.SourceItem{
	segSourceItem(0),
	tasm.NewSourceItem("", "fjaxu", []string{fJC, jJCExtended, aJCExtended, zero, "target"}),
	iarSourceItem("", 1),
	iarSourceItem("target", 0),
}

func Test_JC_Extended_Pos(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpCarryExtended)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), true)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(false)
		ute.GetEngine().GetDesignatorRegister().SetCarry(true)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStopped(t, engine)
	checkProgramAddress(t, engine, 01003)
}

var jumpDivideFault = []*tasm.SourceItem{
	segSourceItem(0),
	tasm.NewSourceItem("", "fjaxu", []string{fJDF, jJDF, aJDF, zero, "target"}),

	iarSourceItem("", 1),
	iarSourceItem("target", 0),
}

func Test_JDF_Basic_Pos(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpDivideFault)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), false)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(true)
		ute.GetEngine().GetDesignatorRegister().SetDivideCheck(true)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStopped(t, engine)
	checkProgramAddress(t, engine, 01003)
}

func Test_JDF_Basic_Neg(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpDivideFault)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), false)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(true)
		ute.GetEngine().GetDesignatorRegister().SetDivideCheck(false)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStopped(t, engine)
	checkProgramAddress(t, engine, 01002)
}

func Test_JDF_Extended_Pos(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpDivideFault)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), true)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(false)
		ute.GetEngine().GetDesignatorRegister().SetDivideCheck(true)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStopped(t, engine)
	checkProgramAddress(t, engine, 01003)
}

func Test_JDF_Extended_Neg(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpDivideFault)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), true)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(false)
		ute.GetEngine().GetDesignatorRegister().SetDivideCheck(false)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStopped(t, engine)
	checkProgramAddress(t, engine, 01002)
}

var jumpFloatingOverflow = []*tasm.SourceItem{
	segSourceItem(0),
	tasm.NewSourceItem("", "fjaxu", []string{fJFO, jJFO, aJFO, zero, "target"}),
	iarSourceItem("", 1),
	iarSourceItem("target", 0),
}

func Test_JFO_Basic_Pos(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpFloatingOverflow)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), false)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(true)
		ute.GetEngine().GetDesignatorRegister().SetCharacteristicOverflow(true)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStopped(t, engine)
	checkProgramAddress(t, engine, 01003)
}

func Test_JFO_Basic_Neg(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpFloatingOverflow)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), false)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(true)
		ute.GetEngine().GetDesignatorRegister().SetCharacteristicOverflow(false)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStopped(t, engine)
	checkProgramAddress(t, engine, 01002)
}

func Test_JFO_Extended_Pos(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpFloatingOverflow)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), true)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(false)
		ute.GetEngine().GetDesignatorRegister().SetCharacteristicOverflow(true)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStopped(t, engine)
	checkProgramAddress(t, engine, 01003)
}

func Test_JFO_Extended_Neg(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpFloatingOverflow)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), true)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(false)
		ute.GetEngine().GetDesignatorRegister().SetCharacteristicOverflow(false)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStopped(t, engine)
	checkProgramAddress(t, engine, 01002)
}

var jumpFloatingUnderflow = []*tasm.SourceItem{
	segSourceItem(0),
	tasm.NewSourceItem("", "fjaxu", []string{fJFU, jJFU, aJFU, zero, "target"}),
	iarSourceItem("", 1),
	iarSourceItem("target", 0),
}

func Test_JFU_Basic_Pos(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpFloatingUnderflow)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), false)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(true)
		ute.GetEngine().GetDesignatorRegister().SetCharacteristicUnderflow(true)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStopped(t, engine)
	checkProgramAddress(t, engine, 01003)
}

func Test_JFU_Basic_Neg(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpFloatingUnderflow)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), false)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(true)
		ute.GetEngine().GetDesignatorRegister().SetCharacteristicUnderflow(false)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStopped(t, engine)
	checkProgramAddress(t, engine, 01002)
}

func Test_JFU_Extended_Pos(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpFloatingUnderflow)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), true)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(false)
		ute.GetEngine().GetDesignatorRegister().SetCharacteristicUnderflow(true)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStopped(t, engine)
	checkProgramAddress(t, engine, 01003)
}

func Test_JFU_Extended_Neg(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpFloatingUnderflow)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), true)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(false)
		ute.GetEngine().GetDesignatorRegister().SetCharacteristicUnderflow(false)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStopped(t, engine)
	checkProgramAddress(t, engine, 01002)
}

//	--------------------------------------------------------------------------------------------------------------------

var jumpNoOverflow = []*tasm.SourceItem{
	segSourceItem(0),
	tasm.NewSourceItem("", "fjaxu", []string{fJNO, jJNO, aJNO, zero, "target"}),
	iarSourceItem("", 1),
	iarSourceItem("target", 0),
}

func Test_JNO_Basic_Pos(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpNoOverflow)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), false)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(true)
		ute.GetEngine().GetDesignatorRegister().SetOverflow(false)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStopped(t, engine)
	checkProgramAddress(t, engine, 01003)
}

func Test_JNO_Extended_Neg(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpNoOverflow)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), true)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(false)
		ute.GetEngine().GetDesignatorRegister().SetOverflow(true)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStopped(t, engine)
	checkProgramAddress(t, engine, 01002)
}

var jumpNoCarryBasic = []*tasm.SourceItem{
	segSourceItem(0),
	tasm.NewSourceItem("", "fjaxu", []string{fJNC, jJNCBasic, aJNCBasic, zero, "target"}),
	iarSourceItem("", 1),
	iarSourceItem("target", 0),
}

func Test_JNC_Basic_Pos(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpNoCarryBasic)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), false)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(true)
		ute.GetEngine().GetDesignatorRegister().SetCarry(false)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStopped(t, engine)
	checkProgramAddress(t, engine, 01003)
}

func Test_JNC_Basic_Neg(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpNoOverflow)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), false)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(true)
		ute.GetEngine().GetDesignatorRegister().SetOverflow(true)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStopped(t, engine)
	checkProgramAddress(t, engine, 01002)
}

var jumpNoCarryExtended = []*tasm.SourceItem{
	segSourceItem(0),
	tasm.NewSourceItem("", "fjaxu", []string{fJNC, jJNCExtended, aJNCExtended, zero, "target"}),
	iarSourceItem("", 1),
	iarSourceItem("target", 0),
}

func Test_JNC_Extended_Pos(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpNoCarryExtended)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), true)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(false)
		ute.GetEngine().GetDesignatorRegister().SetCarry(false)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStopped(t, engine)
	checkProgramAddress(t, engine, 01003)
}

var jumpNoDivideFault = []*tasm.SourceItem{
	segSourceItem(0),
	tasm.NewSourceItem("", "fjaxu", []string{fJNDF, jJNDF, aJNDF, zero, "target"}),

	iarSourceItem("", 1),
	iarSourceItem("target", 0),
}

func Test_JNDF_Basic_Pos(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpNoDivideFault)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), false)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(true)
		ute.GetEngine().GetDesignatorRegister().SetDivideCheck(false)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStopped(t, engine)
	checkProgramAddress(t, engine, 01003)
}

func Test_JNDF_Basic_Neg(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpNoDivideFault)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), false)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(true)
		ute.GetEngine().GetDesignatorRegister().SetDivideCheck(true)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStopped(t, engine)
	checkProgramAddress(t, engine, 01002)
}

func Test_JNDF_Extended_Pos(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpNoDivideFault)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), true)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(false)
		ute.GetEngine().GetDesignatorRegister().SetDivideCheck(false)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStopped(t, engine)
	checkProgramAddress(t, engine, 01003)
}

func Test_JNDF_Extended_Neg(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpNoDivideFault)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), true)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(false)
		ute.GetEngine().GetDesignatorRegister().SetDivideCheck(true)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStopped(t, engine)
	checkProgramAddress(t, engine, 01002)
}

var jumpNoFloatingOverflow = []*tasm.SourceItem{
	segSourceItem(0),
	tasm.NewSourceItem("", "fjaxu", []string{fJNFO, jJNFO, aJNFO, zero, "target"}),
	iarSourceItem("", 1),
	iarSourceItem("target", 0),
}

func Test_JNFO_Basic_Pos(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpNoFloatingOverflow)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), false)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(true)
		ute.GetEngine().GetDesignatorRegister().SetCharacteristicOverflow(false)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStopped(t, engine)
	checkProgramAddress(t, engine, 01003)
}

func Test_JNFO_Basic_Neg(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpNoFloatingOverflow)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), false)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(true)
		ute.GetEngine().GetDesignatorRegister().SetCharacteristicOverflow(true)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStopped(t, engine)
	checkProgramAddress(t, engine, 01002)
}

func Test_JNFO_Extended_Pos(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpNoFloatingOverflow)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), true)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(false)
		ute.GetEngine().GetDesignatorRegister().SetCharacteristicOverflow(false)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStopped(t, engine)
	checkProgramAddress(t, engine, 01003)
}

func Test_JNFO_Extended_Neg(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpNoFloatingOverflow)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), true)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(false)
		ute.GetEngine().GetDesignatorRegister().SetCharacteristicOverflow(true)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStopped(t, engine)
	checkProgramAddress(t, engine, 01002)
}

var jumpNoFloatingUnderflow = []*tasm.SourceItem{
	segSourceItem(0),
	tasm.NewSourceItem("", "fjaxu", []string{fJNFU, jJNFU, aJNFU, zero, "target"}),
	iarSourceItem("", 1),
	iarSourceItem("target", 0),
}

func Test_JNFU_Basic_Pos(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpNoFloatingUnderflow)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), false)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(true)
		ute.GetEngine().GetDesignatorRegister().SetCharacteristicUnderflow(false)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStopped(t, engine)
	checkProgramAddress(t, engine, 01003)
}

func Test_JNFU_Basic_Neg(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpNoFloatingUnderflow)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), false)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(true)
		ute.GetEngine().GetDesignatorRegister().SetCharacteristicUnderflow(true)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStopped(t, engine)
	checkProgramAddress(t, engine, 01002)
}

func Test_JNFU_Extended_Pos(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpNoFloatingUnderflow)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), true)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(false)
		ute.GetEngine().GetDesignatorRegister().SetCharacteristicUnderflow(false)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStopped(t, engine)
	checkProgramAddress(t, engine, 01003)
}

func Test_JNFU_Extended_Neg(t *testing.T) {
	sourceSet := tasm.NewSourceSet("Test", jumpNoFloatingUnderflow)
	a := tasm.NewTinyAssembler()
	a.Assemble(sourceSet)

	e := tasm.Executable{}
	e.LinkSimple(a.GetSegments(), true)

	ute := NewUnitTestExecutor()
	err := ute.Load(&e)
	if err == nil {
		ute.GetEngine().GetDesignatorRegister().SetBasicModeEnabled(false)
		ute.GetEngine().GetDesignatorRegister().SetCharacteristicUnderflow(true)
		err = ute.Run()
	}

	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	engine := ute.GetEngine()
	checkStopped(t, engine)
	checkProgramAddress(t, engine, 01002)
}
