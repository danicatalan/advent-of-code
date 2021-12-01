package main

import (
	"fmt"
	"strings"

	"github.com/danicatalan/advent-of-code/day13/shuttlesearch"
)

func main() {
	s1 := getSchedules()
	s2 := strings.Split(s1, "\n")[1]
	r1, r2 := shuttlesearch.Part1(s1), shuttlesearch.Part2(s2)
	fmt.Println(r1, r2)
}
