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

func Solve(fname string, linesPer int, proc Processor) {
	fi, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	ext := filepath.Ext(fname)
	outname := fname[0:len(fname)-len(ext)] + "-out" + ext
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

	for i := 1; i <= num; i++ {
		buf := make([]string, linesPer, linesPer)
		for j := 0; j < linesPer; j++ {
			ok := s.Scan()
			if !ok {
				panic(fmt.Sprintf("EOF. expected more lines during problem %v", i))
			}
			buf[j] = s.Text()
		}
		outstr := proc(i, buf)
		outfi.WriteString(fmt.Sprintf("Case #%v: %v\n", i, outstr))
	}
}

/////////////
// Helpers //
/////////////

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
