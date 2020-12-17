package main

import (
	"fmt"

	"github.com/danicatalan/advent-of-code-2020/day11/seatingsystem"
)

func main() {
	l := getLayout()
	r1, r2 := seatingsystem.Part1(l), seatingsystem.Part2(l)
	fmt.Println(r1, r2)
}
