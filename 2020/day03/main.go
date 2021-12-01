package main

import (
	"fmt"

	"github.com/danicatalan/advent-of-code/day03/toboggantrajectory"
)

func main() {
	var m = getMap("data")
	slopes := [][]int{
		{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2},
	}
	r1, r2 := toboggantrajectory.Part1(m), toboggantrajectory.Part2(m, slopes)
	fmt.Println(r1, r2)
}
