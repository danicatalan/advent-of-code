package lanternfish

import (
	"strconv"
	"strings"
)

const MAX_DAYS_PART_1 = 80
const NEW_FISH = 8
const RESET_FISH = 6

//Part1 calculates the amount of fishes after 80 generations.
func Part1(input string) int {
	numbers := strings.Split(input, ",")
	fishes := []int{}
	for _, number := range numbers {
		n, _ := strconv.Atoi(number)
		fishes = append(fishes, n)
	}
	for day := 0; day < MAX_DAYS_PART_1; day++ {
		newGeneration := fishes
		for i := range fishes {
			if fishes[i] == 0 {
				newGeneration[i] = RESET_FISH
				newGeneration = append(newGeneration, NEW_FISH)
			} else {
				newGeneration[i]--
			}
		}
		fishes = newGeneration
	}
	return len(fishes)
}
