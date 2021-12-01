package main

import (
	"fmt"

	"github.com/danicatalan/advent-of-code/day09/encodingerror"
)

func main() {
	n := getNumbers()
	r1 := encodingerror.Part1(n, 25)
	r2 := encodingerror.Part2(n, r1)
	fmt.Println(r1, r2)
}
