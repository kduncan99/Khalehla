// Khalehla Project
// Copyright © 2023 by Kurt Duncan, BearSnake LLC
// All Rights Reserved

package pkg

// AddSimple takes two numbers which are 36-bit signed values packed into uint64's,
// and adds them according to ones-complement rules.
func AddSimple(operand1 uint64, operand2 uint64) uint64 {
	if (operand1 == NegativeZero) && (operand2 == NegativeZero) {
		return NegativeZero
	} else {
		native1 := GetTwosComplement(operand1)
		native2 := GetTwosComplement(operand2)
		return GetOnesComplement(native1 + native2)
	}
}

// And calculates the logical AND of two 36-bit values
func And(operand1 uint64, operand2 uint64) uint64 {
	return (operand1 & operand2) & NegativeZero
}

// Compare indicates whether operand1 is less than, equal to, or greater than operand2.
// the result negative if less than, zero if equal, or positive if greater than.
// For our purposes, negative zero is less than (and thus, not equal to) positive zero.
func Compare(operand1 uint64, operand2 uint64) int {
	if operand1 == operand2 {
		return 0
	}

	pos1 := IsPositive(operand1)
	pos2 := IsPositive(operand2)
	if pos1 && pos2 {
		if operand1 < operand2 {
			return -1
		} else {
			return 1
		}
	} else if !pos1 && !pos2 {
		if operand1 > operand2 {
			return -1
		} else {
			return 1
		}
	} else if pos1 {
		return 1
	} else {
		return -1
	}
}

// CompareDouble indicates whether operand1 is less than, equal to, or greater than operand2.
// the result negative if less than, zero if equal, or positive if greater than.
// For our purposes, negative zero is less than (and thus, not equal to) positive zero.
// both operands consist of a 72-bit value, stored as two consecutive 36-bit values wrapped
// in uint64's.
func CompareDouble(operand1 []uint64, operand2 []uint64) int {
	if (operand1[0] == operand2[0]) && (operand1[1] == operand2[1]) {
		return 0
	}

	pos1 := IsPositive(operand1[0])
	pos2 := IsPositive(operand2[0])
	if pos1 != pos2 {
		if pos1 {
			return 1
		} else {
			return -1
		}
	} else if pos1 {
		if operand1[0] > operand2[0] {
			return 1
		} else if operand1[0] < operand2[0] {
			return -1
		} else {
			if operand1[1] > operand2[1] {
				return 1
			} else if operand1[1] < operand2[1] {
				return -1
			} else {
				return 0
			}
		}
	} else {
		if operand1[0] > operand2[0] {
			return -1
		} else if operand1[0] < operand2[0] {
			return 1
		} else {
			if operand1[1] > operand2[1] {
				return -1
			} else if operand1[1] < operand2[1] {
				return 1
			} else {
				return 0
			}
		}
	}
}

// GetOnesComplement takes a standard twos-complement value and converts it to a
// 36-bit ones-complement value packed in a uint64.
func GetOnesComplement(operand uint64) uint64 {
	if int64(operand) < 0 {
		return Negate(-operand)
	} else {
		return operand
	}
}

// GetSignExtended12 sign-extends an 12-bit value to 36 bits
func GetSignExtended12(value uint64) (result uint64) {
	result = value & 0_7777
	if (result & 0_04000) != 0 {
		result |= 0_777777_770000
	}
	return
}

// GetSignExtended18 sign-extends an 18-bit value to 36 bits
func GetSignExtended18(value uint64) (result uint64) {
	result = value & 0_777777
	if (result & 0_400000) != 0 {
		result |= 0_777777_000000
	}
	return
}

// GetSignExtended24 sign-extends a 24-bit value to 36 bits
func GetSignExtended24(value uint64) (result uint64) {
	result = value & 0_7777_7777
	if (result & 0_4000_0000) != 0 {
		result |= 0_7777_0000_0000
	}
	return
}

// GetTwosComplement takes a number which is a 36-bit signed value packed into a uint64,
// and converts it to twos-complement.
func GetTwosComplement(operand uint64) uint64 {
	if IsNegative(operand) {
		return -Negate(operand)
	} else {
		return operand
	}
}

func IsNegative(value uint64) bool {
	return (value & 0_400000_000000) != 0
}

func IsPositive(value uint64) bool {
	return (value & 0_400000_000000) == 0
}

func IsZero(operand uint64) bool {
	return operand == PositiveZero || operand == NegativeZero
}

func IsDoubleZero(operand []uint64) bool {
	return IsZero(operand[0]) && operand[0] == operand[1]
}

// LeftShiftCircular shifts the 36-bit word to the left by the given count value,
// where every bit shifted out of bit 0 is end-around shifted into bit 35.
func LeftShiftCircular(operand uint64, count uint64) uint64 {
	result := operand
	if count > 0 {
		count %= 36
		if count > 18 {
			result <<= 18
			result |= result >> 36
			result &= NegativeZero
			count -= 18
		}

		if count > 0 {
			result <<= count
			result |= result >> 36
		}
	}

	return result
}

// LeftShiftLogical shifts the 36-bit word to the left by the given count value.
// Bits shifted out of bit 0 are lost, and zeroes are shift into bit 35.
func LeftShiftLogical(operand uint64, count uint64) uint64 {
	if count >= 36 {
		return 0
	} else {
		return (operand << count) & NegativeZero
	}
}

// Magnitude returns the absolute value of the given operand
func Magnitude(operand uint64) uint64 {
	if IsPositive(operand) {
		return operand
	} else {
		return Not(operand)
	}
}

// MagnitudeDouble returns the absolute value of the given 72-bit operand
func MagnitudeDouble(operand uint64) uint64 {
	if IsPositive(operand) {
		return operand
	} else {
		return Not(operand)
	}
}

// Negate returns the additive inverse of a given 36-bit signed value packed into a uint64
func Negate(operand uint64) uint64 {
	return operand ^ NegativeZero
}

func Not(op uint64) uint64 {
	return (op ^ NegativeZero) & NegativeZero
}

func Or(lhs uint64, rhs uint64) uint64 {
	return (lhs | rhs) & NegativeZero
}

func Xor(lhs uint64, rhs uint64) uint64 {
	return (lhs ^ rhs) & NegativeZero
}

// RightShiftAlgebraic shifts the 36-bit word to the left by the given count value,
// Bits shifted out of bit 35 are lost while bit 0 is propagated to the right.
func RightShiftAlgebraic(operand uint64, count uint64) uint64 {
	if count >= 36 {
		return 0
	} else if count > 0 {
		//	TODO
	} else {
		return operand
	}
}

// RightShiftCircular shifts the 36-bit word to the right by the given count value,
// where every bit shifted out of bit 35 is end-around shifted into bit 0.
func RightShiftCircular(operand uint64, count uint64) uint64 {
	result := operand
	if count > 0 {
		count %= 36
		if count > 18 {
			result <<= 18
			result |= result >> 36
			result &= NegativeZero
			count -= 18
		}

		if count > 0 {
			result <<= count
			result |= result >> 36
		}
	}

	return result
}

// RightShiftLogical shifts the 36-bit word to the left by the given count value.
// Bits shifted out of bit 35 are lost, and zeroes are shift into bit 0.
func RightShiftLogical(operand uint64, count uint64) uint64 {
	if count >= 36 {
		return 0
	} else if count > 0 {
		return (operand & NegativeZero) >> count
	} else {
		return operand
	}
}
