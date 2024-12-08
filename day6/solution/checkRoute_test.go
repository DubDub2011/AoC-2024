package solution

import (
	gr "day6/grid"
	"testing"
)

func TestCheckRoute(t *testing.T) {

	points := [][]rune{
		{'.', '.', '.', '.', '#', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '#'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '#', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '#', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '#', '.', '#', '^', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '#', '.'},
		{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
	}

	pos := gr.Position{4, 6}
	grid := gr.Grid{points}

	if grid.GetPos(pos.X, pos.Y) != GuardRune {
		t.Error("Wrong starting position for guard")
	}

	res := checkRoute(pos, grid)
	t.Logf("Grid: %s", grid.String())
	if !res {
		t.Error("Expected route to loop")
	}
}
