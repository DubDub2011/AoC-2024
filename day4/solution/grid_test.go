package solution

import "testing"

func TestGrid(t *testing.T) {
	g := grid{
		runes: [][]rune{
			{'0', '1', '2', '3'},
			{'4', '5', '6', '7'},
			{'8', '9', 'A', 'B'},
			{'C', 'D', 'E', 'F'},
		},
	}

	// General pos test
	if g.getPos(0, 0) != '0' {
		t.Errorf("Expected (0, 0) to be 0, got %s", string(g.getPos(0, 0)))
	}

	// Expect y axis to get correct pos
	if g.getPos(0, 2) != '8' {
		t.Errorf("Expected (0, 2) to be 8, got %s", string(g.getPos(0, 2)))
	}

	// Expect x axis to get correct pos
	if g.getPos(2, 0) != '2' {
		t.Errorf("Expected (2, 0) to be 2, got %s", string(g.getPos(2, 0)))
	}

	// Expect both axis to return correct pos
	if g.getPos(3, 3) != 'F' {
		t.Errorf("Expected (3, 3) to be F, got %s", string(g.getPos(3, 3)))
	}

	// Expect out of bounds to return false
	if g.getPos(4, 0) != ' ' {
		t.Errorf("Expected (4, 0) to be out of bounds")
	}

	if g.getPos(0, 4) != ' ' {
		t.Errorf("Expected (0, 4) to be out of bounds")
	}

	if g.getPos(-1, 0) != ' ' {
		t.Errorf("Expected (-1, 0) to be out of bounds")
	}
	if g.getPos(0, -1) != ' ' {
		t.Errorf("Expected (0, -1) to be out of bounds")
	}
}
