package main

import (
	"fmt"
	"bufio"
	"gojam/base"
	"strings"
)

func main() {
	base.SolveVariable("test.txt", func(n int, bs *bufio.Scanner) string {
		s := NewSolver(n, bs)
		fmt.Printf("%v\n", s)
		v := s.Batch()
		fmt.Printf("%v\n", v)
		return fmt.Sprintf("%v", v)
	})
}

// return 
func (t *Solver) Batch() string {
	out := []string{"1", "2", "3"}
	return strings.Join(out, " ")
}


/////////////////////////
// Boiler platey stuff //
/////////////////////////

type Solver struct {
	flavors int 
	numcust int 
	cust []*Customer
}

func (t *Solver) String() string {
	return fmt.Sprintf("f[%v] n[%v] cust[%v]", t.flavors, t.numcust, t.cust)
}

func NewSolver(n int, bs *bufio.Scanner) *Solver {
	flavors := base.GetInt(base.GetNextLine(n, bs))
	numcust := base.GetInt(base.GetNextLine(n, bs))
	cust := make([]*Customer, numcust, numcust)
	for i := 0; i < numcust; i++ {
		l := base.GetIntList(base.GetNextLine(n, bs))
		n := l[0]
		flav := make([]int, n, n)
		malt := make([]bool, n, n)
		for j := 1; j <= n; j++ {
			flav[j - 1] = l[j * 2 - 1]
			malt[j - 1] = l[j * 2] == 1
		}
		cust[i] = &Customer{flav, malt}
	}
	return &Solver{
		flavors,
		numcust,
		cust,
	}
}

type Customer struct {
	flav []int
	malt []bool
}

func (t *Customer) String() string {
	return fmt.Sprintf("flav[%v] malt[%v]", t.flav, t.malt)
}
