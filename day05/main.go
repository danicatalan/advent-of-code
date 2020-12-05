package main

import (
	"fmt"

	"github.com/danicatalan/advent-of-code-2020/day05/binaryboarding"
)

func main() {
	var seats []int
	r1, r2 := 0, 0
	for _, bp := range getBoardingPasses() {
		ID := binaryboarding.Part1(bp)
		seats = append(seats, ID)
		if ID > r1 {
			r1 = ID
		}
	}
	r2 = binaryboarding.Part2(seats)
	fmt.Println(r1, r2)
}
