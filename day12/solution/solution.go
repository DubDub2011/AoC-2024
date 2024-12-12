package solution

import (
	gr "day12/grid"
	"fmt"
	"strings"
)

func FencePricer(input [][]rune) int {
	grid := gr.New(input)

	alreadyVisited := make(map[gr.Position]bool)
	total := 0
	for pos := range grid.Positions() {
		if alreadyVisited[pos] {
			continue
		}

		regionPositions := []gr.Position{pos}
		val := grid.GetPos(pos.X, pos.Y)
		perimiter, area := 0, 0
		for idx := 0; idx < len(regionPositions); idx++ {
			regionPos := regionPositions[idx]
			if alreadyVisited[regionPos] {
				continue
			}

			up, right, down, left := regionPos.Up(), regionPos.Right(), regionPos.Down(), regionPos.Left()
			if grid.GetPos(up.X, up.Y) == val {
				regionPositions = append(regionPositions, up)
			} else {
				perimiter++
			}

			if grid.GetPos(right.X, right.Y) == val {
				regionPositions = append(regionPositions, right)
			} else {
				perimiter++
			}

			if grid.GetPos(down.X, down.Y) == val {
				regionPositions = append(regionPositions, down)
			} else {
				perimiter++
			}

			if grid.GetPos(left.X, left.Y) == val {
				regionPositions = append(regionPositions, left)
			} else {
				perimiter++
			}
			area++
			alreadyVisited[regionPos] = true
		}

		total += area * perimiter
	}
	return total
}

func FencePricerDiscount(input [][]rune) int {
	grid := gr.New(input)

	alreadyVisited := make(map[gr.Position]bool)
	total := 0
	for pos := range grid.Positions() {
		if alreadyVisited[pos] {
			continue
		}

		regionPositions := []gr.Position{pos}
		val := grid.GetPos(pos.X, pos.Y)
		sides, area := 0, 0
		for idx := 0; idx < len(regionPositions); idx++ {
			regionPos := regionPositions[idx]
			if alreadyVisited[regionPos] {
				regionPositions = append(regionPositions[:idx], regionPositions[idx+1:]...)
				idx--
				continue
			}

			up, right, down, left := regionPos.Up(), regionPos.Right(), regionPos.Down(), regionPos.Left()
			if grid.GetPos(up.X, up.Y) == val {
				regionPositions = append(regionPositions, up)
			}

			if grid.GetPos(right.X, right.Y) == val {
				regionPositions = append(regionPositions, right)
			}

			if grid.GetPos(down.X, down.Y) == val {
				regionPositions = append(regionPositions, down)
			}

			if grid.GetPos(left.X, left.Y) == val {
				regionPositions = append(regionPositions, left)
			}

			area++
			alreadyVisited[regionPos] = true
		}

		sides = calculateSides(regionPositions, grid)

		total += area * sides
	}

	fmt.Println(grid.String())
	return total
}

func calculateSides(shape []gr.Position, grid gr.Grid) int {
	sidePieces := make([]gr.Position, 0)
	for _, shapePos := range shape {
		val := grid.GetPos(shapePos.X, shapePos.Y)
		up, right, down, left := shapePos.Up(), shapePos.Right(), shapePos.Down(), shapePos.Left()

		notSide := true
		if grid.GetPos(up.X, up.Y) != val {
			notSide = false
		}

		if grid.GetPos(right.X, right.Y) != val {
			notSide = false
		}

		if grid.GetPos(down.X, down.Y) != val {
			notSide = false
		}

		if grid.GetPos(left.X, left.Y) != val {
			notSide = false
		}

		if !notSide {
			sidePieces = append(sidePieces, shapePos)
		}
	}

	// Got all side pieces, use a map to store already calculated sides
	sidesUsed := make(map[gr.Position]bool)
	sides := 0
	for _, sidePiece := range sidePieces {
		val := grid.GetPos(sidePiece.X, sidePiece.Y)
		up, right, down, left := sidePiece.Up(), sidePiece.Right(), sidePiece.Down(), sidePiece.Left()
		if grid.GetPos(up.X, up.Y) != val {
			isNewSide := true
			isNewSide = isNewSide && !checkIfSideCounted(sidePiece, grid, gr.Position.Left, gr.Position.Up, sidesUsed)
			isNewSide = isNewSide && !checkIfSideCounted(sidePiece, grid, gr.Position.Right, gr.Position.Up, sidesUsed)

			if isNewSide {
				sides++
			}
		}

		if grid.GetPos(right.X, right.Y) != val {
			isNewSide := true

			isNewSide = isNewSide && !checkIfSideCounted(sidePiece, grid, gr.Position.Up, gr.Position.Right, sidesUsed)
			isNewSide = isNewSide && !checkIfSideCounted(sidePiece, grid, gr.Position.Down, gr.Position.Right, sidesUsed)

			if isNewSide {
				sides++
			}
		}

		if grid.GetPos(down.X, down.Y) != val {
			isNewSide := true
			isNewSide = isNewSide && !checkIfSideCounted(sidePiece, grid, gr.Position.Left, gr.Position.Down, sidesUsed)
			isNewSide = isNewSide && !checkIfSideCounted(sidePiece, grid, gr.Position.Right, gr.Position.Down, sidesUsed)
			if isNewSide {
				sides++
			}
		}

		if grid.GetPos(left.X, left.Y) != val {
			isNewSide := true
			isNewSide = isNewSide && !checkIfSideCounted(sidePiece, grid, gr.Position.Up, gr.Position.Left, sidesUsed)
			isNewSide = isNewSide && !checkIfSideCounted(sidePiece, grid, gr.Position.Down, gr.Position.Left, sidesUsed)
			if isNewSide {
				sides++
			}
		}

		sidesUsed[sidePiece] = true
	}

	for _, side := range sidePieces {
		val := grid.GetPos(side.X, side.Y)
		val = rune(strings.ToLower(string(val))[0])
		grid.SetPos(side.X, side.Y, val)
	}
	return sides
}

func checkIfSideCounted(point gr.Position, grid gr.Grid, parallel, perpendicular func(gr.Position) gr.Position, usedSides map[gr.Position]bool) bool {
	sideCounted := false

	val := grid.GetPos(point.X, point.Y)
	directionalPoint := point
	for {
		directionalPoint = parallel(directionalPoint)
		if grid.GetPos(directionalPoint.X, directionalPoint.Y) != val {
			break
		}

		inOne := perpendicular(directionalPoint)
		if grid.GetPos(inOne.X, inOne.Y) == val {
			break
		}

		if usedSides[directionalPoint] {
			cornerPoint := perpendicular(directionalPoint)
			if grid.GetPos(cornerPoint.X, cornerPoint.Y) != val {
				sideCounted = true
			}
		}
	}
	return sideCounted
}
