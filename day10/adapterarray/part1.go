package adapterarray

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

//Part1 ...
func Part1(input string) int {
	numbers := []int{0}
	for _, l := range strings.Split(input, "\n") {
		n, err := strconv.Atoi(l)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, n)
	}
	sort.Ints(numbers)
	numbers = append(numbers, numbers[len(numbers)-1]+3)
	diff := map[int]int{}
	for i := range numbers {
		if i < len(numbers)-1 {
			diff[numbers[i+1]-numbers[i]]++
		}
	}
	return diff[1] * diff[3]
}
