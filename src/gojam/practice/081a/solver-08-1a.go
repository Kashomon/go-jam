package main

import (
	"fmt"
	"gojam/base"
	"sort"
)

func main() {
	base.Solve("A-large-practice.in", 3, func(n int, lines []string) string {
		s := NewSolver(lines)
		min := s.MinInnerProduct()
		fmt.Printf("%v\n", min)
		return fmt.Sprintf("%v", min)
	})
}

type Solver struct {
	l    int
	arr1 []int
	arr2 []int
}

func NewSolver(lines []string) *Solver {
	l := base.GetInt(lines[0])
	a1 := base.GetIntList(lines[1])
	a2 := base.GetIntList(lines[2])
	return &Solver{
		l,
		a1,
		a2,
	}
}

// Algorithm: Multiply the largest number by the smallest number.
func (t *Solver) MinInnerProduct() int {
	a1 := sort.IntSlice(t.arr1)
	a1.Sort()
	a2 := sort.IntSlice(t.arr2)
	a2.Sort()
	base.ReverseInts(a2)

	out := 0
	for i := 0; i < len(a1); i++ {
		out += a1[i] * a2[i]
	}
	return out
}
