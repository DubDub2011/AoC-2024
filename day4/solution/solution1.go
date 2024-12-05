package solution

func XMASOccurances(input [][]rune) int {
	gr := grid{
		runes: input,
	}

	width, length := gr.getWidth(), gr.getLength()

	totalOccurances := 0
	for y := 0; y < length; y++ {
		for x := 0; x < width; x++ {
			pos := gr.getPos(x, y)
			if pos == 'X' {
				totalOccurances += search(position{x, y}, gr)
			}
		}
	}

	return totalOccurances
}

func search(pos position, gr grid) int {
	found := 0

	found += checkDirection(pos, position.right, gr)
	found += checkDirection(pos, position.left, gr)
	found += checkDirection(pos, position.up, gr)
	found += checkDirection(pos, position.down, gr)

	found += checkDirection(pos, position.upRight, gr)
	found += checkDirection(pos, position.upLeft, gr)
	found += checkDirection(pos, position.downRight, gr)
	found += checkDirection(pos, position.downLeft, gr)

	return found
}

func checkDirection(startPos position, mover func(position) position, gr grid) int {
	pos := mover(startPos)
	if gr.getPos(pos.x, pos.y) != 'M' {
		return 0
	}

	pos = mover(pos)
	if gr.getPos(pos.x, pos.y) != 'A' {
		return 0
	}

	pos = mover(pos)
	if gr.getPos(pos.x, pos.y) != 'S' {
		return 0
	}

	return 1
}
