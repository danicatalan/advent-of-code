package lanternfish

import (
	"strconv"
	"strings"
)

const MAX_DAYS_PART_2 = 256

//Part2 calculates the amount of fishes after 256 generations.
//Uses a map to parallelize the growth.
func Part2(input string) int {
	numbers := strings.Split(input, ",")
	fishes := make(map[int]int)
	for _, number := range numbers {
		n, _ := strconv.Atoi(number)
		fishes[n]++
	}
	for day := 0; day < MAX_DAYS_PART_2; day++ {
		newGeneration := make(map[int]int)
		for k, v := range fishes {
			if k == 0 {
				newGeneration[RESET_FISH] += v
				newGeneration[NEW_FISH] += v
			} else {
				newGeneration[k-1] += v
			}
		}
		fishes = newGeneration
	}
	output := 0
	for _, v := range fishes {
		output += v
	}
	return output
}
