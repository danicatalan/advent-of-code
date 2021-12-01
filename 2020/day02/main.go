package main

import (
	"fmt"

	"github.com/danicatalan/advent-of-code/day02/passwordphilosophy"
)

func main() {
	var passwords = getPasswords("data")
	r1, r2 := 0, 0
	for _, p := range passwords {
		if passwordphilosophy.Part1(p) == true {
			r1++
		}
		if passwordphilosophy.Part2(p) == true {
			r2++
		}
	}
	fmt.Println(r1, r2)
}
