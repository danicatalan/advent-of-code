package encodingerror

import (
	"log"
	"strconv"
	"strings"
)

//Part1 ...
func Part1(input string, preamble int) int {
	numbers := []int{}
	out := 0
	for _, l := range strings.Split(input, "\n") {
		n, err := strconv.Atoi(l)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, n)
	}
	for i, n := range numbers {
		if i < preamble {
			continue
		}
		prev := numbers[i-preamble : i]
		if !checkSums(n, prev) {
			out = n
			break
		}
	}
	return out
}

func checkSums(n int, prev []int) bool {
	for _, p := range prev {
		for _, q := range prev {
			if p != q && p+q == n {
				return true
			}
		}
	}
	return false
}
