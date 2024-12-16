package solution

import (
	gr "day16/grid"
	"fmt"
	"sort"
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

	fmt.Printf("Start, End: %v, %v\n", startPos, targetPos)
	fmt.Printf("Grid: %s\n", grid.String())

	reindeers := []*Reindeer{NewReindeer(startPos)}
	scores := []int{}
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
				reindeers = append(reindeers, newReindeer)
			}

			if right {
				newReindeer := reindeer.Clone()
				newReindeer.TurnRight()
				reindeers = append(reindeers, newReindeer)
			}

			reindeer.Move()
			val := grid.GetPos(reindeer.currPos.X, reindeer.currPos.Y)
			if val == '.' {
				continue
			} else if val == '#' {
				reindeers = append(reindeers[:idx], reindeers[idx+1:]...) // crashed and burned :(
				idx--
			} else if val == 'E' {
				scores = append(scores, reindeer.score)
				reindeers = append(reindeers[:idx], reindeers[idx+1:]...) // he won :)
				idx--
			}
		}
	}

	sort.Ints(scores)

	return scores[0]
}
