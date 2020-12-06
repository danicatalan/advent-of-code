package main

import (
	"fmt"

	"github.com/danicatalan/advent-of-code-2020/day06/customcustoms"
)

func main() {
	fmt.Println('a', 'z')
	answers := getAnswers()
	r1, r2 := 0, 0
	for _, a := range answers {
		r1 += customcustoms.Part1(a)
		r2 += customcustoms.Part2(a)
	}
	fmt.Println(r1, r2)
}
