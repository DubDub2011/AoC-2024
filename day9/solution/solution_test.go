package solution

import "testing"

func TestCharToInt(t *testing.T) {
	res := runeToInt('4')
	if res != 4 {
		t.Errorf("expected 4, got %d", res)
	}

	res = runeToInt('0')
	if res != 0 {
		t.Errorf("expected 4, got %d", res)
	}

	res = runeToInt('9')
	if res != 9 {
		t.Errorf("expected 4, got %d", res)
	}

}

func TestFileCount(t *testing.T) {
	res := findFileCount([]byte("12345"))
	if res != 3 {
		t.Errorf("expected 3, got %d", res)
	}

	res = findFileCount([]byte("02045"))
	if res != 1 {
		t.Errorf("expected 1, got %d", res)
	}
}
