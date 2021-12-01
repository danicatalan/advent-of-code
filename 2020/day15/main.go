package main

import (
	"fmt"

	"github.com/danicatalan/advent-of-code/day15/rambunctiousrecitation"
)

func main() {
	n := getStartingNumbers()
	r1, r2 := rambunctiousrecitation.Part1(n), rambunctiousrecitation.Part2(n)
	fmt.Println(r1, r2)
}
