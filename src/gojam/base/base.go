// Some Base utilites for codejam for the Go programming language.
package base

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Processor func(pnum int, lines []string) string

type VariableProcessor func(pnum int, b *bufio.Scanner) string

type ProblemReader struct {
	fi *os.File
	outfi *os.File
	s *bufio.Scanner
	num  int
}

func NewProblemReader (fname string) *ProblemReader {
	fi, err := os.Open(fname)
	if err != nil {
		panic(err)
	}

	ext := filepath.Ext(fname)
	outname := fname[0:len(fname)-len(ext)] + ".out"
	fmt.Printf("Writing output to %v\n", outname)
	outfi, oerr := os.Create(outname)
	if oerr != nil {
		panic(oerr)
	}

	s := bufio.NewScanner(fi)
	ok := s.Scan()
	if !ok {
		panic("EOF. Expected number of problems")
	}
	num := GetInt(s.Text())
	return &ProblemReader{ fi, outfi, s, num }
}

// Most problems have a uniform number of lines per problem.
func Solve(fname string, linesPer int, proc Processor) {
	p := NewProblemReader(fname)
	defer p.fi.Close()
	defer p.outfi.Close()
	for i := 1; i <= p.num; i++ {
		buf := make([]string, linesPer, linesPer)
		for j := 0; j < linesPer; j++ {
			buf[j] = GetNextLine(i, p.s)
		}
		outstr := proc(i, buf)
		p.outfi.WriteString(fmt.Sprintf("Case #%v: %v\n", i, outstr))
	}
}

func GetNextLine(i int, s *bufio.Scanner) string {
	ok := s.Scan()
	if !ok {
		panic(fmt.Sprintf("EOF. expected more lines during problem %v", i))
	}
	return s.Text()
}

// Some problems have a variable number of lines per problem.
func SolveVariable(fname string, proc VariableProcessor) {
	p := NewProblemReader(fname)
	defer p.fi.Close()
	defer p.outfi.Close()
	for i := 1; i <= p.num; i++ {
		outstr := proc(i, p.s)
		p.outfi.WriteString(fmt.Sprintf("Case #%v: %v\n", i, outstr))
	}
}

/////////////
// Helpers //
/////////////

// INTS

func GetInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}

func GetIntList(s string) []int {
	splat := strings.Split(s, " ")
	l := len(splat)
	buf := make([]int, l, l)
	for i := 0; i < l; i++ {
		buf[i] = GetInt(splat[i])
	}
	return buf
}

func ReverseInts(f []int) []int {
	var end = len(f) - 1
	for i := 0; i < len(f)/2; i++ {
		var l = f[i]
		var r = f[end-i]
		f[i] = r
		f[end-i] = l
	}
	return f
}

// Floats

func GetFloat(s string) float64 {
	num, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return num
}

func GetFloatList(s string) []float64 {
	splat := strings.Split(s, " ")
	l := len(splat)
	buf := make([]float64, l, l)
	for i := 0; i < l; i++ {
		buf[i] = GetFloat(splat[i])
	}
	return buf
}

func ReverseFloats(f []float64) []float64 {
	var end = len(f) - 1
	for i := 0; i < len(f)/2; i++ {
		var l = f[i]
		var r = f[end-i]
		f[i] = r
		f[end-i] = l
	}
	return f
}

// General
func ReverseAny(f []interface{}) []interface{} {
	var end = len(f) - 1
	for i := 0; i < len(f)/2; i++ {
		var l = f[i]
		var r = f[end-i]
		f[i] = r
		f[end-i] = l
	}
	return f
}
