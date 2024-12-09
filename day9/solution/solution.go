package solution

func DiskFragmenter(input []byte) int {
	pointer, filePointer := 0, len(input)-1
	fileCount := findFileCount(input)
	leftFileID, rightFileID := 0, fileCount-1

	resultPositionCounter := 0

	if filePointer%2 != 0 { // files are on even positions
		filePointer--
	}
	res := make(map[int]int, 0)
	for {
		onFile := pointer%2 == 0
		if onFile {
			fileSize := runeToInt(rune(input[pointer]))
			for idx := 0; idx < fileSize; idx++ {
				res[resultPositionCounter+idx] = leftFileID
			}
			leftFileID++
			resultPositionCounter += fileSize
		} else {
			for {
				fileSize := runeToInt(rune(input[filePointer]))
				whiteSpaceSize := runeToInt(rune(input[pointer]))
				if fileSize > whiteSpaceSize {
					for idx := 0; idx < whiteSpaceSize; idx++ {
						res[resultPositionCounter+idx] = rightFileID
					}
					resultPositionCounter += whiteSpaceSize
					input[filePointer] = byte(intToRune(fileSize - whiteSpaceSize)) // this is fine.. probably
					break
				}

				for idx := 0; idx < fileSize; idx++ {
					res[resultPositionCounter+idx] = rightFileID
				}
				rightFileID--
				resultPositionCounter += fileSize
				input[pointer] = byte(intToRune(whiteSpaceSize - fileSize))
				input[filePointer] = '0'
				filePointer -= 2
			}
		}

		if pointer >= filePointer {
			break
		}
		pointer++
	}

	checkSum := 0
	for idx, fileID := range res {
		checkSum += idx * fileID
	}
	return checkSum
}

func findFileCount(disk []byte) int {
	sum := 0
	for idx := 0; idx < len(disk); idx += 2 {
		if disk[idx] != '0' {
			sum++
		}
	}
	return sum
}

func runeToInt(val rune) int {
	if val < '0' || val > '9' {
		return 0
	}
	return int(val) - 48
}

func intToRune(val int) rune {
	if val > 9 {
		panic("can't be greater than 10")
	}
	if val < 0 {
		panic("can't be negative")
	}
	return rune(val + 48)
}
