package main

import "testing"

func TestDetectSquares(t *testing.T) {
	ds := NewDetectSquares()
	ds.Add([2]int{3, 10})
	ds.Add([2]int{11, 2})
	ds.Add([2]int{3, 2})

	if got := ds.Count([2]int{11, 10}); got != 1 {
		t.Errorf("Count = %d, want 1", got)
	}
	if got := ds.Count([2]int{14, 8}); got != 0 {
		t.Errorf("Count = %d, want 0", got)
	}

	// A duplicate corner doubles the square count through multiplicity.
	ds.Add([2]int{11, 2})
	if got := ds.Count([2]int{11, 10}); got != 2 {
		t.Errorf("after duplicate: %d, want 2", got)
	}

	// Zero-area "squares" (the query point itself stored) must not count.
	zs := NewDetectSquares()
	zs.Add([2]int{1, 1})
	zs.Add([2]int{1, 1})
	if got := zs.Count([2]int{1, 1}); got != 0 {
		t.Errorf("zero area = %d, want 0", got)
	}

	// All four corners stored, query at each corner sees the same square.
	sq := NewDetectSquares()
	for _, p := range [][2]int{{0, 0}, {0, 2}, {2, 0}, {2, 2}} {
		sq.Add(p)
	}
	for _, p := range [][2]int{{0, 0}, {0, 2}, {2, 0}, {2, 2}} {
		if got := sq.Count(p); got != 1 {
			t.Errorf("corner %v sees %d squares, want 1", p, got)
		}
	}
}
