package solution

import (
	"strings"
)

func DiskFragmenter2(input []byte) int {
	disk, fileData := expand(input)

	idx := 0
	for {
		freeSpace := disk[idx] == '0'
		if freeSpace {
			end := idx + 1
			for {
				if end == len(disk) && disk[end] == '0' {
					break
				}
				end++
			}
			freeSpaceLength := end - idx
			file := searchForFileThatFits(fileData, freeSpaceLength, idx)
			if file != nil {
				filePosition, fileLength, fileIdx := file[0], file[1], file[2]
				// update disk to insert file here and remove file from above
				for fileReplacerIdx := idx; fileReplacerIdx <= idx+fileLength; fileReplacerIdx++ {
					disk[fileReplacerIdx] = '1'
				}

				for fileReplacerIdx := filePosition; fileReplacerIdx <= filePosition+fileLength; fileReplacerIdx++ {
					disk[fileReplacerIdx] = '0'
				}
				fileData[fileIdx] = []int{idx, fileLength}
			}
		}
		idx++
		if idx == len(disk) {
			break
		}
	}
	return calcFileSpace(fileData)
}

func searchForFileThatFits(fileData map[int][]int, size int, stopIndex int) []int {
	for idx := len(fileData); idx >= 0; idx-- {
		data := fileData[idx]
		fileIdx := data[0]
		if fileIdx > stopIndex {
			return nil
		}

		fileLength := data[1]
		if fileLength < size {
			return append(data, idx)
		}
	}
	return nil
}

func calcFileSpace(fileData map[int][]int) int {
	sum := 0
	for fileIdx, file := range fileData {
		filePosition, fileLength := file[0], file[1]
		for x := 0; x < fileLength; x++ {
			sum += (filePosition + x) * fileIdx
		}
	}
	return sum
}

func expand(input []byte) ([]byte, map[int][]int) { // int slice contains the file position and the length
	fileCounter := 0
	disk := make([]byte, 0)
	fileData := make(map[int][]int)
	for idx, char := range input {
		partLen := runeToInt(rune(char))
		if idx%2 == 0 { // on file pointer
			if partLen == 0 {
				continue
			}
			fileData[fileCounter] = []int{len(disk), partLen}
			disk = append(disk, strings.Repeat("1", partLen)...)
			fileCounter++
		} else {
			disk = append(disk, strings.Repeat("0", partLen)...)
		}
	}
	return disk, fileData
}
