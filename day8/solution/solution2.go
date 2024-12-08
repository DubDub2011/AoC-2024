package solution

import (
	gr "day8/grid"
)

func ResonanceHarmonics(input [][]rune) int {
	grid := gr.New(input)

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
			copy(temp, positions) // needed due to append having side-effects
			others := append(temp[:idx], temp[idx+1:]...)

			for _, other := range others {
				results.SetPos(current.X, current.Y, AntinodeRune)
				results.SetPos(other.X, other.Y, AntinodeRune)

				towardsOtherX := other.X - current.X
				towardsOtherY := other.Y - current.Y
				var nextPos gr.Position = other
				for {
					nextPos = gr.Position{nextPos.X + towardsOtherX, nextPos.Y + towardsOtherY}
					results.SetPos(nextPos.X, nextPos.Y, AntinodeRune)
					outOfBounds := nextPos.X < 0 ||
						nextPos.X >= results.GetWidth() ||
						nextPos.Y < 0 ||
						nextPos.Y >= results.GetLength()
					if outOfBounds {
						break
					}
				}

				towardsCurrentX := current.X - other.X
				towardsCurrentY := current.Y - other.Y
				nextPos = current
				for {
					nextPos = gr.Position{nextPos.X + towardsCurrentX, nextPos.Y + towardsCurrentY}
					results.SetPos(nextPos.X, nextPos.Y, AntinodeRune)
					outOfBounds := nextPos.X < 0 ||
						nextPos.X >= results.GetWidth() ||
						nextPos.Y < 0 ||
						nextPos.Y >= results.GetLength()
					if outOfBounds {
						break
					}
				}
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
