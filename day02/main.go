package main

import (
	"fmt"
)

func main() {
	var passwords = getPasswords("data")
	r1, r2 := 0, 0
	for _, p := range passwords {
		if Part1(p) == true {
			r1++
		}
		if Part2(p) == true {
			r2++
		}
	}
	fmt.Println(r1, r2)
}
