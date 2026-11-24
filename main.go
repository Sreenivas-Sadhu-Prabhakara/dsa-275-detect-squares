package main

import "fmt"

// DetectSquares counts axis-aligned squares formed with a query point:
// Add stores points (duplicates counted); Count(p) sums, over every stored
// point diagonal to p, the product of the multiplicities of the two
// corners completing the square.
type DetectSquares struct {
	count map[[2]int]int
}

// NewDetectSquares returns an empty point store.
func NewDetectSquares() *DetectSquares {
	return &DetectSquares{count: map[[2]int]int{}}
}

// Add stores a point; duplicates accumulate.
func (d *DetectSquares) Add(p [2]int) {
	d.count[p]++
}

// Count returns how many axis-aligned squares of positive area have p as
// one corner and three stored points (with multiplicity) as the others.
func (d *DetectSquares) Count(p [2]int) int {
	total := 0
	for q, n := range d.count {
		dx, dy := q[0]-p[0], q[1]-p[1]
		if dx == 0 || dy == 0 || dx != dy && dx != -dy {
			continue // not a diagonal partner of p
		}
		// Corners completing the square share p's x or y.
		total += n * d.count[[2]int{p[0], q[1]}] * d.count[[2]int{q[0], p[1]}]
	}
	return total
}

func main() {
	ds := NewDetectSquares()
	ds.Add([2]int{3, 10})
	ds.Add([2]int{11, 2})
	ds.Add([2]int{3, 2})
	fmt.Println(ds.Count([2]int{11, 10}))
	ds.Add([2]int{11, 2})
	fmt.Println(ds.Count([2]int{11, 10}))
}
