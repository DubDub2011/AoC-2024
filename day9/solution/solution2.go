package solution

import (
	"strings"
)

func DiskFragmenter2(input []byte) uint64 {
	disk, fileData := expand(input)

	for fileNumber := len(fileData) - 1; fileNumber >= 0; fileNumber-- {
		data := fileData[fileNumber]
		if len(data) == 0 {
			continue
		}
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
		fileData[fileNumber] = []int{freeSpaceIdx, fileLength}
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

func calcFileSpace(fileData map[int][]int) uint64 {
	sum := uint64(0)
	for fileNumber, file := range fileData {
		startIdx, fileLength := file[0], file[1]
		for fileIdx := startIdx; fileIdx < startIdx+fileLength; fileIdx++ {
			sum += uint64(fileIdx) * uint64(fileNumber)
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
				fileCounter++
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
