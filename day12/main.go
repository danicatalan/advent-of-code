package main

import (
	"fmt"

	"github.com/danicatalan/advent-of-code-2020/day12/rainrisk"
)

func main() {
	instructions := getInstructions()
	r1, r2 := rainrisk.Part1(instructions), rainrisk.Part2(instructions)
	fmt.Println(r1, r2)
}
