package solution_test

import (
	"day15/solution"
	"fmt"
	"testing"
)

// Each problem is getting harder to test by purely the inputs
// Going to start breaking down testing into specific areas
func TestRobotMover(t *testing.T) {

	testCases := []struct {
		mapData  [][]rune
		commands []rune
		output   int
	}{
		{
			[][]rune{
				[]rune("#####"),
				[]rune("#...#"),
				[]rune("#.@.#"),
				[]rune("#...#"),
				[]rune("#####"),
			}, []rune("<v>>^^<<v>"), 0,
		},
		{
			[][]rune{
				[]rune("#########"),
				[]rune("#.......#"),
				[]rune("#.....O.#"),
				[]rune("#.......#"),
				[]rune("#...@...#"),
				[]rune("#.......#"),
				[]rune("#...O...#"),
				[]rune("#.....O.#"),
				[]rune("#########"),
			}, []rune(">vv<<<v<^^^^^v>>>>>^>^<<<<<<v>>>>>>>>>vvvvvvvvvv<<<<<"), 904, // should end with (1,1),(1,2),(7,1)
		},
		{
			[][]rune{
				[]rune("########"),
				[]rune("#..O.O.#"),
				[]rune("##@.O..#"),
				[]rune("#...O..#"),
				[]rune("#.#.O..#"),
				[]rune("#...O..#"),
				[]rune("#......#"),
				[]rune("########"),
			}, []rune("<^^>>>vv<v>>v<<"), 2028, // should end with (1,1),(1,2),(7,1)
		},
		{
			[][]rune{
				[]rune("##########"),
				[]rune("#..O..O.O#"),
				[]rune("#......O.#"),
				[]rune("#.OO..O.O#"),
				[]rune("#..O@..O.#"),
				[]rune("#O#..O...#"),
				[]rune("#O..O..O.#"),
				[]rune("#.OO.O.OO#"),
				[]rune("#....O...#"),
				[]rune("##########"),
			}, []rune(`<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^`), 10092, // should end with (1,1),(1,2),(7,1)
		},
	}

	for idx, testCase := range testCases {
		t.Run(fmt.Sprintf("Test case %d", idx), func(t *testing.T) {
			res := solution.BoxChecker(testCase.mapData, testCase.commands)
			if res != testCase.output {
				t.Errorf("Expected %d, got %d", testCase.output, res)
			}

		})
	}
}
