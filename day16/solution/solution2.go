package solution

import (
	gr "day16/grid"
)

func SeatFinder(mapData [][]rune) int {
	grid := gr.New(mapData)

	var startPos gr.Position
	for pos := range grid.Positions() {
		if grid.GetPos(pos.X, pos.Y) == 'S' {
			startPos = pos
		}
	}

	positionScores := make([][]int, grid.GetWidth())
	for idx := range positionScores {
		positionScores[idx] = make([]int, grid.GetLength())
		for idy := range positionScores[idx] {
			positionScores[idx][idy] = 9999999999
		}
	}

	reindeers := []*Reindeer{NewReindeer(startPos)}
	for {
		if len(reindeers) == 0 {
			break
		}

		for idx := 0; idx < len(reindeers); idx++ {
			reindeer := reindeers[idx]
			left, right := reindeer.CanTurn(grid) // can see us getting stuck in a loop...
			if left {
				newReindeer := reindeer.Clone()
				newReindeer.TurnLeft()
				newReindeer.Move()
				reindeers = append(reindeers, newReindeer)
			}

			if right {
				newReindeer := reindeer.Clone()
				newReindeer.TurnRight()
				newReindeer.Move()
				reindeers = append(reindeers, newReindeer)
			}

			reindeer.Move()
			val := grid.GetPos(reindeer.currPos.X, reindeer.currPos.Y)
			if val == '#' {
				reindeers = append(reindeers[:idx], reindeers[idx+1:]...) // crashed and burned :(
				idx--
				continue
			}

			if positionScores[reindeer.currPos.X][reindeer.currPos.Y] > reindeer.score {
				positionScores[reindeer.currPos.X][reindeer.currPos.Y] = reindeer.score
			} else {
				reindeers = append(reindeers[:idx], reindeers[idx+1:]...) // too slow :(
				idx--
				continue
			}

			if val == 'E' {
				reindeers = append(reindeers[:idx], reindeers[idx+1:]...) // he won :)
				idx--
				continue
			}

		}
	}

	// bit lazy but run it again with the updated grid position values, now we can see which reindeers finish
	// take their positions and count them up
	reindeers = []*Reindeer{NewReindeer(startPos)}
	winnerPositions := make([]gr.Position, 0)
	for {
		if len(reindeers) == 0 {
			break
		}

		for idx := 0; idx < len(reindeers); idx++ {
			reindeer := reindeers[idx]
			left, right := reindeer.CanTurn(grid) // can see us getting stuck in a loop...
			if left {
				newReindeer := reindeer.Clone()
				newReindeer.TurnLeft()
				newReindeer.Move()
				reindeers = append(reindeers, newReindeer)
			}

			if right {
				newReindeer := reindeer.Clone()
				newReindeer.TurnRight()
				newReindeer.Move()
				reindeers = append(reindeers, newReindeer)
			}

			reindeer.Move()
			val := grid.GetPos(reindeer.currPos.X, reindeer.currPos.Y)
			if val == '#' {
				reindeers = append(reindeers[:idx], reindeers[idx+1:]...) // crashed and burned :(
				idx--
				continue
			}

			if positionScores[reindeer.currPos.X][reindeer.currPos.Y] < reindeer.score {
				reindeers = append(reindeers[:idx], reindeers[idx+1:]...) // too slow :(
				idx--
				continue
			}

			if val == 'E' {
				winnerPositions = append(winnerPositions, reindeer.path...)
				reindeers = append(reindeers[:idx], reindeers[idx+1:]...) // he won :)
				idx--
				continue
			}
		}
	}

	allKeys := make(map[gr.Position]bool)
	list := []gr.Position{}
	for _, pos := range winnerPositions {
		if !allKeys[pos] {
			allKeys[pos] = true
			list = append(list, pos)
		}
	}

	return len(list)
}
