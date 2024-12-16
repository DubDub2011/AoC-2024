package solution

import (
	gr "day16/grid"
	"fmt"
)

func ReindeerMaze(mapData [][]rune) int {
	grid := gr.New(mapData)

	fmt.Printf("Grid: %s\n", grid.String())
	return 10
}
