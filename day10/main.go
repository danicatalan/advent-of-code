package main

import (
	"fmt"

	"github.com/danicatalan/advent-of-code-2020/day10/adapterarray"
)

func main() {
	j := getJolgates()
	r1, r2 := adapterarray.Part1(j), adapterarray.Part2(j)
	fmt.Println(r1, r2)
}
