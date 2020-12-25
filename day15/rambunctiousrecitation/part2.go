package rambunctiousrecitation

import (
	"log"
	"strconv"
	"strings"
)

//Part2 ...
func Part2(input string) int {
	limit := 30000000
	numbers, mem := []int{}, make(map[int]int)
	for _, n := range strings.Split(input, ",") {
		num, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, num)
	}
	last := 0
	for i := 0; i < limit; i++ {
		if i < len(numbers) {
			mem[numbers[i]] = i
			continue
		}
		if pi, found := mem[last]; found {
			mem[last] = i
			if i != limit-1 {
				last = i - pi
			}
		} else {
			mem[last] = i
			if i != limit-1 {
				last = 0
			}
		}
	}
	return last
}
