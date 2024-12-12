package solution

import gr "day12/grid"

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
			if sidesUsed[left] {
				aboveLeft := left.Up()
				if grid.GetPos(aboveLeft.X, aboveLeft.Y) != val {
					isNewSide = false
				}
			}

			if sidesUsed[right] {
				aboveRight := right.Up()
				if grid.GetPos(aboveRight.X, aboveRight.Y) != val {
					isNewSide = false
				}
			}
			if isNewSide {
				sides++
			}
		}

		if grid.GetPos(right.X, right.Y) != val {
			isNewSide := true
			if sidesUsed[up] {
				upRight := up.Right()
				if grid.GetPos(upRight.X, upRight.Y) != val {
					isNewSide = false
				}
			}

			if sidesUsed[down] {
				downRight := down.Right()
				if grid.GetPos(downRight.X, downRight.Y) != val {
					isNewSide = false
				}
			}
			if isNewSide {
				sides++
			}
		}

		if grid.GetPos(down.X, down.Y) != val {
			isNewSide := true
			if sidesUsed[right] {
				rightDown := right.Down()
				if grid.GetPos(rightDown.X, rightDown.Y) != val {
					isNewSide = false
				}
			}

			if sidesUsed[left] {
				leftDown := left.Down()
				if grid.GetPos(leftDown.X, leftDown.Y) != val {
					isNewSide = false
				}
			}
			if isNewSide {
				sides++
			}
		}

		if grid.GetPos(left.X, left.Y) != val {
			isNewSide := true
			if sidesUsed[up] {
				upLeft := up.Left()
				if grid.GetPos(upLeft.X, upLeft.Y) != val {
					isNewSide = false
				}
			}

			if sidesUsed[down] {
				downLeft := down.Left()
				if grid.GetPos(downLeft.X, downLeft.Y) != val {
					isNewSide = false
				}
			}
			if isNewSide {
				sides++
			}
		}

		sidesUsed[sidePiece] = true
	}
	return sides
}
