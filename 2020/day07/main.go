package main

import (
	"fmt"

	"github.com/danicatalan/advent-of-code/day07/handyhaversacks"
)

func main() {
	rules := getRules()
	r1, r2 := handyhaversacks.Part1(rules), handyhaversacks.Part2(rules)
	fmt.Println(r1, r2)
}
