package main

import (
	"fmt"

	"github.com/danicatalan/advent-of-code/day16/tickettranslation"
)

func main() {
	n := getNotes()
	r1, r2 := tickettranslation.Part1(n), tickettranslation.Part2(n)
	fmt.Println(r1, r2)
}
