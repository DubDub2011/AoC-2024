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
			nextPos := idx + 1
			if nextPos == len(positions) {
				nextPos = 0
			}
			a, b := positions[idx], positions[nextPos]
			if a == b {
				continue
			}
			a2b := gr.Position{b.X + (b.X - a.X), b.Y + (b.Y - a.Y)}
			b2a := gr.Position{a.X + (a.X - b.X), a.Y + (a.Y - b.Y)}
			results.SetPos(a2b.X, a2b.Y, AntinodeRune)
			results.SetPos(b2a.X, b2a.Y, AntinodeRune)
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
