package main

import (
	"fmt"

	"github.com/danicatalan/advent-of-code-2020/day01/reportrepair"
)

func main() {
	var expenses = getExpenses("data")
	r1, r2 := reportrepair.Part1(expenses), reportrepair.Part2(expenses)
	fmt.Println(r1, r2)
}
