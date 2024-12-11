package solution

import (
	"strings"
)

func DiskFragmenter2(input []byte) int {
	disk, fileData := expand(input)

	for filePosition := len(fileData) - 1; filePosition >= 0; filePosition-- {
		data := fileData[filePosition]
		fileIdx, fileLength := data[0], data[1]

		freeSpaceIdx := findFreeSpace(disk, fileIdx, fileLength)
		if freeSpaceIdx == -1 {
			continue
		}

		for idx := freeSpaceIdx; idx < freeSpaceIdx+fileLength; idx++ {
			disk[idx] = '1'
		}

		for idx := fileIdx; idx < fileIdx+fileLength; idx++ {
			disk[idx] = '0'
		}
		fileData[filePosition] = []int{freeSpaceIdx, fileLength}
	}

	return calcFileSpace(fileData)
}

func findFreeSpace(disk []byte, limit int, size int) int {
	idx := 0
	for {
		freeSpace := disk[idx] == '0'
		if freeSpace {
			end := idx + 1
			for {
				if end == len(disk) {
					break
				} else if disk[end] != '0' {
					break
				}
				end++
			}
			freeSpaceLength := end - idx
			if freeSpaceLength >= size {
				return idx
			}
		}

		idx++
		if idx >= limit {
			break
		}
	}
	return -1
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
