package solution

import "testing"

func TestAdv(t *testing.T) {
	t.Run("adv() - 0/4", func(t *testing.T) {
		A = 0
		operand := 2

		adv(operand)

		if A != 0 {
			t.Errorf("adv() = %v, want %v", A, 2)
		}
	})

	t.Run("adv() - 2/2", func(t *testing.T) {
		A = 2
		operand := 1

		adv(operand)

		if A != 1 {
			t.Errorf("adv() = %v, want %v", A, 2)
		}
	})

	t.Run("adv() - 500/4", func(t *testing.T) {
		A = 500
		operand := 2

		adv(operand)

		if A != 125 {
			t.Errorf("adv() = %v, want %v", A, 2)
		}
	})

	t.Run("adv() - 500/8", func(t *testing.T) {
		A = 500
		operand := 3

		adv(operand)

		if A != 62 {
			t.Errorf("adv() = %v, want %v", A, 2)
		}
	})
}

func TestBdv(t *testing.T) {
	t.Run("bdv() - 0/4", func(t *testing.T) {
		A = 0
		operand := 2

		bdv(operand)

		if B != 0 {
			t.Errorf("bdv() = %v, want %v", B, 2)
		}
	})

	t.Run("bdv() - 2/2", func(t *testing.T) {
		A = 2
		operand := 1

		bdv(operand)

		if B != 1 {
			t.Errorf("bdv() = %v, want %v", B, 2)
		}
	})

	t.Run("bdv() - 500/4", func(t *testing.T) {
		A = 500
		operand := 2

		bdv(operand)

		if B != 125 {
			t.Errorf("bdv() = %v, want %v", B, 2)
		}
	})

	t.Run("bdv() - 500/8", func(t *testing.T) {
		A = 500
		operand := 3

		bdv(operand)

		if B != 62 {
			t.Errorf("bdv() = %v, want %v", B, 2)
		}
	})
}

func TestCdv(t *testing.T) {
	t.Run("cdv() - 0/4", func(t *testing.T) {
		A = 0
		operand := 2

		cdv(operand)

		if C != 0 {
			t.Errorf("cdv() = %v, want %v", C, 2)
		}
	})

	t.Run("cdv() - 2/2", func(t *testing.T) {
		A = 2
		operand := 1

		cdv(operand)

		if C != 1 {
			t.Errorf("cdv() = %v, want %v", C, 2)
		}
	})

	t.Run("cdv() - 500/4", func(t *testing.T) {
		A = 500
		operand := 2

		cdv(operand)

		if C != 125 {
			t.Errorf("cdv() = %v, want %v", C, 2)
		}
	})

	t.Run("cdv() - 500/8", func(t *testing.T) {
		A = 500
		operand := 3

		cdv(operand)

		if C != 62 {
			t.Errorf("cdv() = %v, want %v", C, 2)
		}
	})
}

func TestBxl(t *testing.T) {
	t.Run("bxl() - 0/7", func(t *testing.T) {
		B = 0
		operand := 7

		bxl(operand)

		if B != 7 {
			t.Errorf("bxl() = %v, want %v", B, 7)
		}
	})

	t.Run("bxl() - 12345/54321", func(t *testing.T) {
		B = 12345
		operand := 54321

		bxl(operand)

		if B != 58376 {
			t.Errorf("bxl() = %v, want %v", B, 58376)
		}
	})
}
