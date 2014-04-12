package main

import (
	"fmt"
	"gojam/base"
)

func main() {
	base.Solve("test.txt", 3, func(n int, lines []string) string {
		fmt.Printf("%v\n", lines)
		return "foo"
	})
}
