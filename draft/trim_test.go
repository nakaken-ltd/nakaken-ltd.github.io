package main

import "testing"

func TestPolygon(t *testing.T) {
	// p represents the following polygon:
	// (2,  3)-----(9,   0)
	//   |            |
	//   |            |
	//   |            |
	// (1, 11)-----(10, 10)
	p := polygon{
		S: []stroke{
			{2, 3, 1, 11},
			{1, 11, 10, 10},
			{10, 10, 9, 0},
			{9, 0, 2, 3},
		},
	}

	// Boundary test

	if p.Left() != 1 {
		t.Errorf("left: expected 1, got %d", p.Left())
	}

	if p.Right() != 10 {
		t.Errorf("right: expected 10, got %d", p.Right())
	}

	if p.Top() != 0 {
		t.Errorf("top: expected 0, got %d", p.Top())
	}

	if p.Bottom() != 11 {
		t.Errorf("bottom: expected 11, got %d", p.Bottom())
	}

	if p.Width() != 9 {
		t.Errorf("width: expected 9, got %d", p.Width())
	}

	if p.Height() != 11 {
		t.Errorf("height: expected 11, got %d", p.Height())
	}

	// Inside/outside calculation test

	ss := []struct {
		X, Y   int
		Inside bool
	}{
		{2, 2, false},
		{2, 4, true},
		{2, 10, true},
		{2, 12, false},
		{3, 2, false},
		{3, 3, true},
		{3, 10, true},
		{3, 11, false},
		{4, 2, false},
		{4, 3, true},
		{4, 10, true},
		{4, 11, false},
		{5, 1, false},
		{5, 2, true},
		{5, 10, true},
		{5, 11, false},
		{6, 1, false},
		{6, 2, true},
		{6, 10, true},
		{6, 11, false},
		{7, 0, false},
		{7, 1, true},
		{7, 10, true},
		{7, 11, false},
		{8, 0, false},
		{8, 1, true},
		{8, 10, true},
		{8, 11, false},
		{9, -1, false},
		{9, 1, true},
		{9, 10, true},
		{9, 11, false},
	}

	for _, s := range ss {
		actual := p.Inside(s.X, s.Y)
		if actual != s.Inside {
			t.Errorf("inside (%d, %d): expected %v, got %v", s.X, s.Y, s.Inside, actual)
		}
	}
}
