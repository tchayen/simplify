package simplify

import "testing"

func TestArea(t *testing.T) {
	if area(Point{0, 0, 0}, Point{4, 0, 0}, Point{4, 4, 0}) != 8 {
		t.Error("incorrect area")
	}
}

func TestSimplify(t *testing.T) {
	t.Run("sample line", func(t *testing.T) {
		Simplify([]Point{{0, 3, 0}, {3, 0, 0}, {6, 2, 0}, {9, 0, 0}, {12, 1, 0}, {15, 3, 0}})
	})

	t.Run("line with collinear points", func(t *testing.T) {
		Simplify([]Point{{0, 0, 0}, {1, 0, 0}, {2, 0, 0}, {3, 2, 0}})
	})

	t.Run("line shorter than 3 points", func(t *testing.T) {
		Simplify([]Point{{0, 0, 0}, {0, 1, 0}})
	})
}
