package solution

import gr "day10/grid"

func HeightmapRouter(heightmap [][]rune) int {
	grid := gr.New(heightmap)
	trailheads := getTrailheads(grid)

	sum := 0
	for _, pos := range trailheads {
		sum += route(pos)
	}

	return 0
}

func route(startingPos gr.Position) int {
	return 0
}

func getTrailheads(grid gr.Grid) []gr.Position {
	trailheads := make([]gr.Position, 0)
	for pos := range grid.Positions() {
		if grid.GetPos(pos.X, pos.Y) == '0' {
			trailheads = append(trailheads, pos)
		}
	}
	return trailheads
}
