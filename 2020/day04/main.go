package main

import (
	"fmt"

	"github.com/danicatalan/advent-of-code/day04/passportprocessing"
)

func main() {
	passports := getPassports()
	r1, r2 := 0, 0
	for _, p := range passports {
		if passportprocessing.Part1(p) {
			r1++
		}
		if passportprocessing.Part1(p) && passportprocessing.Part2(p) {
			r2++
		}
	}
	fmt.Println(r1, r2)
}
