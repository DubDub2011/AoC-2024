package solution

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func Multiplier(input string) int {
	reader := bufio.NewReader(strings.NewReader(input))

	total := 0
	for {
		char := readRune(reader)

		if char == rune(156) { //terminator
			break
		}

		if char == 'm' {
			reader.UnreadRune()
			success, first, second := readMultiplier(reader)
			if success {
				total += first * second
			}
		}

	}

	return total
}

func Multiplier2(input string) int {
	reader := bufio.NewReader(strings.NewReader(input))

	total := 0

	var skip bool
	for {
		char := readRune(reader)

		if char == rune(156) { //terminator
			break
		}

		if !skip {
			if char == 'm' {
				reader.UnreadRune()
				success, first, second := readMultiplier(reader)
				if success {
					total += first * second
				}
			}
		}

		if char == 'd' {
			reader.UnreadRune()
			success, shouldSkip := readDivider(reader)
			if success {
				skip = shouldSkip
			}
		}

	}

	return total
}

func readMultiplier(reader *bufio.Reader) (bool, int, int) {
	currChar := readRune(reader)
	if currChar != 'm' {
		return false, 0, 0
	}

	currChar = readRune(reader)
	if currChar != 'u' {
		return false, 0, 0
	}

	currChar = readRune(reader)
	if currChar != 'l' {
		return false, 0, 0
	}

	currChar = readRune(reader)
	if currChar != '(' {
		return false, 0, 0
	}

	var num1, num2 string
	var next, fail bool
	for {
		currChar = readRune(reader)
		if currChar >= '0' && currChar <= '9' {
			if !next {
				num1 += string(currChar)
			} else {
				num2 += string(currChar)
			}
		} else if currChar == ',' {
			next = true
		} else if currChar == ')' {
			break
		} else {
			fail = true
			break
		}
	}

	if fail {
		return false, 0, 0
	}

	return true, toInt(num1), toInt(num2)
}

func readDivider(reader *bufio.Reader) (bool, bool) {
	currChar := readRune(reader)
	if currChar != 'd' {
		return false, false
	}

	currChar = readRune(reader)
	if currChar != 'o' {
		return false, false
	}

	currChar = readRune(reader)
	var shouldSkip bool
	if currChar == 'n' {
		// handle "dont"
		currChar = readRune(reader)
		if currChar != '\'' {
			return false, false
		}
		currChar = readRune(reader)
		if currChar != 't' {
			return false, false
		}
		shouldSkip = true
		currChar = readRune(reader)
	}

	if currChar != '(' {
		return false, false
	}

	currChar = readRune(reader)
	if currChar != ')' {
		return false, false
	}

	return true, shouldSkip
}

func readRune(reader *bufio.Reader) rune {
	char, _, err := reader.ReadRune()
	if err == io.EOF {
		return rune(156) // terminator
	} else if err != nil {
		panic(err)
	}
	return char
}

func toInt(inp string) int {
	val, err := strconv.Atoi(inp)
	if err != nil {
		panic(err)
	}
	return val
}
