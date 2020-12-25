package main

import (
	"fmt"

	"github.com/danicatalan/advent-of-code-2020/day14/dockingdata"
)

func main() {
	i := getInstructions()
	r1, r2 := dockingdata.Part1(i), dockingdata.Part2(i)
	fmt.Println(r1, r2)
}
