package solution

import (
	gr "day15/grid"
	"testing"
)

func TestRobot(t *testing.T) {
	robot := NewRobot(gr.Position{0, 0}, []rune(">>vv"))

	for {
		mover, _ := robot.GetMoveDirection()
		if mover == nil {
			break
		}

		robot.Move(mover)
	}

	expected := gr.Position{2, 2}
	if robot.Position != expected {
		t.Errorf("Expected %v, got %v", expected, robot.Position)
	}

	robot.commands = []rune(">v<^>v<^>v<^") // move in circle 3 times
	for {
		mover, _ := robot.GetMoveDirection()
		if mover == nil {
			break
		}

		robot.Move(mover)
	}

	expected = gr.Position{2, 2}
	if robot.Position != expected {
		t.Errorf("Expected %v, got %v", expected, robot.Position)
	}

	robot.commands = []rune("<^^<") // move in circle 3 times
	for {
		mover, _ := robot.GetMoveDirection()
		if mover == nil {
			break
		}

		robot.Move(mover)
	}

	expected = gr.Position{0, 0}
	if robot.Position != expected {
		t.Errorf("Expected %v, got %v", expected, robot.Position)
	}

}
