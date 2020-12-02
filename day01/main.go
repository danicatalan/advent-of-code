package main

import (
	"fmt"
)

func main() {
	var expenses = getExpenses("data")
	r1, r2 := Part1(expenses), Part2(expenses)
	fmt.Println(r1, r2)
}
