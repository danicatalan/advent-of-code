package adapterarray

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

//Part2 ...
func Part2(input string) int {
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
	arr := map[int]int{0: 1}
	for i := range numbers {
		if i < len(numbers)-1 {
			arr[numbers[i+1]] = arr[numbers[i+1]-1] + arr[numbers[i+1]-2] + arr[numbers[i+1]-3]
		}
	}
	return arr[numbers[len(numbers)-1]]
}
