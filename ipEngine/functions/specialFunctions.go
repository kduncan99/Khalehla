// Khalehla Project
// Copyright © 2023 by Kurt Duncan, BearSnake LLC
// All Rights Reserved

package functions

import "khalehla/ipEngine"

// NoOperation evaluates the HIU field, but takes no other action (it does x-register incrementation)
func NoOperation(e *ipEngine.InstructionEngine) (completed bool, interrupt ipEngine.Interrupt) {
	completed, _, interrupt = e.GetJumpOperand(false)
	return
}
