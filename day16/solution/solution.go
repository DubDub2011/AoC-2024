package solution

import (
	gr "day16/grid"
)

func ReindeerMaze(mapData [][]rune) int {
	grid := gr.New(mapData)

	var startPos, targetPos gr.Position
	for pos := range grid.Positions() {
		if grid.GetPos(pos.X, pos.Y) == 'S' {
			startPos = pos
		}

		if grid.GetPos(pos.X, pos.Y) == 'E' {
			targetPos = pos
		}
	}

	reindeers := []*Reindeer{NewReindeer(startPos)}
	positionScores := make([][]int, grid.GetWidth())
	for idx := range positionScores {
		positionScores[idx] = make([]int, grid.GetLength())
		for idy := range positionScores[idx] {
			positionScores[idx][idy] = 9999999999
		}
	}

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

	return positionScores[targetPos.X][targetPos.Y]
}
