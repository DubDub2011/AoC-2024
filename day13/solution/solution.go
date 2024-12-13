package solution

func ClawMachine(input [][]float64) int {
	// wonder if this wont work because computers don't math
	tokens := 0
	for _, puzzle := range input {
		x := []float64{puzzle[0], puzzle[2], puzzle[4]} // #a + #b = #prize
		y := []float64{puzzle[1], puzzle[3], puzzle[5]} // #a + #b = #prize

		// see if the way it's done on paper works
		x = []float64{x[0], -x[1], x[2]}                     // minus #b to get #a = #prize - #b
		x = []float64{x[0] / x[0], x[1] / x[0], x[2] / x[0]} // divide by #a to get a = (#prize - #b) / #a

		y = []float64{y[0], y[1] + (x[1] * y[0]), y[2] - (x[2] * y[0])}

		b := y[2] / y[1]
		a := x[2] + (x[1] * b)

		if b != float64(int(b)) {
			continue
		} else if a != float64(int(a)) {
			continue
		}

		tokens += int(a)*3 + int(b)
	}
	return tokens
}
