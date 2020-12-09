package main

import (
	"fmt"

	"github.com/danicatalan/advent-of-code-2020/day08/handheldhalting"
)

func main() {
	instructions := getInstructions()
	r1, r2 := handheldhalting.Part1(instructions), handheldhalting.Part2(instructions)
	fmt.Println(r1, r2)
}
