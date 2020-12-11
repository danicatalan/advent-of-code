package encodingerror

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

//Part2 ...
func Part2(input string, target int) int {
	numbers := []int{}
	for _, l := range strings.Split(input, "\n") {
		n, err := strconv.Atoi(l)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, n)
	}
	set := searchContiguousSet(numbers, target)
	sort.Ints(set)
	return set[0] + set[len(set)-1]
}

func searchContiguousSet(numbers []int, target int) []int {
	for i, n := range numbers {
		acc, set := n, []int{n}
		for _, m := range numbers[i:] {
			if acc == target {
				return set
			}
			if n != m && acc < target {
				acc += m
				set = append(set, m)
			}
		}
	}
	return []int{}
}
