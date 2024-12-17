package solution

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var (
	A, B, C            int
	instructionPointer int
	results            []string
)

func ThreeBitComputer(registers, program []int) string {
	A = registers[0]
	B = registers[1]
	C = registers[2]

	for {
		if instructionPointer >= len(program) {
			break
		}
		opcode, operand := program[instructionPointer], program[instructionPointer+1]

		switch opcode {
		case 0:
			adv(operand)
		case 1:
			bxl(operand)
		case 2:
			bst(operand)
		case 3:
			if jnz(operand) {
				continue
			}
		case 4:
			bxc(operand)
		case 5:
			out(operand)
		case 6:
			bdv(operand)
		case 7:
			cdv(operand)
		default:
			panic("unknown opcode")
		}

		instructionPointer += 2
	}

	return strings.Join(results, ",")
}

func adv(operand int) {
	fmt.Printf("adv by %d\n", operand)
	numerator := A

	if operand == 4 {
		operand = A
	} else if operand == 5 {
		operand = B
	} else if operand == 6 {
		operand = C
	}
	denominator := powInt(2, operand)

	A = numerator / denominator
}

func bxl(operand int) {
	fmt.Printf("bxl by %d\n", operand)

	B = bitwiseOr(B, operand)
}

func bst(operand int) {
	fmt.Printf("bst by %d\n", operand)

	if operand == 4 {
		operand = A
	} else if operand == 5 {
		operand = B
	} else if operand == 6 {
		operand = C
	}

	B = operand % 8
}

func jnz(operand int) bool {
	fmt.Printf("jnz by %d\n", operand)

	if A == 0 {
		return false
	}

	fmt.Printf("jumped to %d\n", operand)
	instructionPointer = operand
	return true
}

func bxc(operand int) {
	fmt.Printf("bxc by %d\n", operand)
	B = bitwiseOr(B, C)
}

func out(operand int) {
	fmt.Printf("out by %d\n", operand)
	if operand == 4 {
		operand = A
	} else if operand == 5 {
		operand = B
	} else if operand == 6 {
		operand = C
	}

	fmt.Printf("outValue %s\n", strconv.Itoa(operand%8))
	results = append(results, strconv.Itoa(operand%8))
}

func bdv(operand int) {
	fmt.Printf("bdv by %d\n", operand)
	numerator := A

	if operand == 4 {
		operand = A
	} else if operand == 5 {
		operand = B
	} else if operand == 6 {
		operand = C
	}
	denominator := powInt(2, operand)

	B = numerator / denominator
}

func cdv(operand int) {
	fmt.Printf("cdv by %d\n", operand)
	numerator := A

	if operand == 4 {
		operand = A
	} else if operand == 5 {
		operand = B
	} else if operand == 6 {
		operand = C
	}
	denominator := powInt(2, operand)

	C = numerator / denominator
}

// helpers
func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func bitwiseOr(a, b int) int {
	aBin := strconv.FormatInt(int64(a), 2)
	bBin := strconv.FormatInt(int64(b), 2)

	idx := 0
	res := ""
	for {
		left, right := '0', '0'

		if idx < len(aBin) {
			left = rune(aBin[len(aBin)-idx-1])
		}

		if idx < len(bBin) {
			right = rune(bBin[len(bBin)-idx-1])
		}

		if left == '1' || right == '1' {
			res = "1" + res
		} else {
			res = "0" + res
		}

		if idx >= len(aBin) && idx >= len(bBin) {
			break
		}

		idx++
	}

	final, err := strconv.ParseInt(res, 2, 0)
	if err != nil {
		panic(err)
	}
	return int(final)
}
