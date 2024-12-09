package solution

import (
	"sort"
)

func DiskFragmenter2(input []byte) int {
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
			if rune(input[pointer]) != 'X' {
				leftFileID++
			}
			resultPositionCounter += fileSize
		} else {
			whiteSpaceSize := runeToInt(rune(input[pointer]))

			tempFilePointer, tempFileID := filePointer, rightFileID
			for {
				if rune(input[tempFilePointer]) == 'X' {
					tempFilePointer -= 2
					continue
				}

				fileSize := runeToInt(rune(input[tempFilePointer]))
				if fileSize == 0 {
					tempFilePointer -= 2
					tempFileID--
					if tempFilePointer <= pointer {
						break
					}
					continue
				}

				if fileSize <= whiteSpaceSize { // fits
					for idx := 0; idx < fileSize; idx++ {
						res[resultPositionCounter+idx] = tempFileID
					}
					resultPositionCounter += fileSize
					whiteSpaceSize -= fileSize
					input[pointer] = byte(intToRune(whiteSpaceSize))

					final := make([]byte, len(input))
					copy(final, input)
					// update the input to insert the whitespace left behind
					final = append(final[:tempFilePointer], '0', byte(intToRune(fileSize)), 'X')
					input = append(final, input[tempFilePointer+1:]...)
					filePointer += 2
				}

				tempFilePointer -= 2
				tempFileID--
				if tempFilePointer <= pointer {
					break
				}
			}

			for idx := 0; idx < whiteSpaceSize; idx++ {
				res[resultPositionCounter+idx] = 0
			}
			resultPositionCounter += whiteSpaceSize
		}

		if pointer >= filePointer {
			break
		}
		pointer++
	}

	keys := []int{}
	for k := range res {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	test := []int{}
	for k := range keys {
		test = append(test, res[k])
	}

	checkSum := 0
	for idx, fileID := range res {
		checkSum += idx * fileID
	}
	return checkSum
}
