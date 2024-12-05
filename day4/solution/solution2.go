package solution

func XMASOccurances2(input [][]rune) int {
	gr := grid{
		runes: input,
	}

	width, length := gr.getWidth(), gr.getLength()

	totalOccurances := 0
	for y := 0; y < length; y++ {
		for x := 0; x < width; x++ {
			pos := gr.getPos(x, y)
			if pos == 'A' {
				totalOccurances += search2(position{x, y}, gr)
			}
		}
	}

	return totalOccurances
}

func search2(pos position, gr grid) int {
	found := checkDirection2(pos.upLeft(), position.downRight, gr)
	found = found && checkDirection2(pos.upRight(), position.downLeft, gr)

	if found {
		return 1
	}

	return 0
}

func checkDirection2(startPos position, mover func(position) position, gr grid) bool {
	startVal := gr.getPos(startPos.x, startPos.y)
	if startVal != 'M' && startVal != 'S' {
		return false
	}

	// move twice, we already know the 2nd pos is A
	pos := mover(startPos)
	pos = mover(pos)

	if startVal == 'M' && gr.getPos(pos.x, pos.y) != 'S' {
		return false
	} else if startVal == 'S' && gr.getPos(pos.x, pos.y) != 'M' {
		return false
	}

	return true
}
