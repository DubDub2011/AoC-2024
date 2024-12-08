package solution

import (
	gr "day8/grid"
)

const (
	EmptyRune    = '.'
	AntinodeRune = '#'
)

func Resonance(input [][]rune) int {
	grid := gr.New(input)

	// first get the positions for each frequency
	antennaPositions := make(map[rune][]gr.Position)
	for pos := range grid.Positions() {
		val := grid.GetPos(pos.X, pos.Y)
		if val != EmptyRune {
			antennaPositions[val] = append(antennaPositions[val], pos)
		}
	}

	results := gr.New(emptyPoints(len(input), len(input[0])))
	for _, positions := range antennaPositions {
		for idx := 0; idx < len(positions); idx++ {
			// get two positions being processed
			current := positions[idx]

			temp := make([]gr.Position, len(positions))
			copy(temp, positions)
			others := append(temp[:idx], temp[idx+1:]...)

			for _, other := range others {
				current2Other := gr.Position{other.X + (other.X - current.X), other.Y + (other.Y - current.Y)}
				other2Current := gr.Position{current.X + (current.X - other.X), current.Y + (current.Y - other.Y)}
				results.SetPos(current2Other.X, current2Other.Y, AntinodeRune)
				results.SetPos(other2Current.X, other2Current.Y, AntinodeRune)
			}
		}
	}

	total := 0
	for pos := range results.Positions() {
		val := results.GetPos(pos.X, pos.Y)
		if val == AntinodeRune {
			total++
		}
	}
	return total
}

func emptyPoints(width, length int) [][]rune {
	res := make([][]rune, width)
	for idx := range res {
		res[idx] = make([]rune, length)
	}
	return res
}
