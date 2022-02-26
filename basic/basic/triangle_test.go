package basic

import "testing"

func TestCalcTriangle(t *testing.T) {
	testing := []struct {
		a, b, c int
	}{
		{3, 4, 5},
		{5, 12, 0},
		{8, 15, 17},
		{7, 24, 0},
		{30000, 40000, 50000},
	}

	for _, s := range testing {
		if target := calcTriangle(s.a, s.b); target != s.c {
			t.Errorf("func expected %d but result is %d", s.c, target)
		}
	}
}
