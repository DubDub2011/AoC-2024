package solution

func DiskFragmenter(input []byte) int {
	pointer, filePointer := 0, len(input)-1
	fileCount := findFileCount(input)
	leftFileID, _ := 0, fileCount-1

	resultPositionCounter := 0

	if filePointer%2 != 0 { // files are on even positions
		filePointer--
	}
	res := make(map[int]int, 0)
	for {
		onFile := pointer%2 == 0
		if onFile {
			fileSize := charToInt(rune(input[pointer]))
			for idx := 0; idx < fileSize; idx++ {
				res[resultPositionCounter+idx] = leftFileID
			}
			leftFileID++
			resultPositionCounter += fileSize
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

func charToInt(val rune) int {
	if val < '0' || val > '9' {
		return 0
	}
	return int(val) - 48
}
