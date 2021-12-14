package thetreacheryofwhales

import (
	"math"
	"strconv"
	"strings"
)

//Part1 calculates the horizontal position that the crabs can align to using the least fuel possible.
func Part1(input string) int {
	positions := strings.Split(input, ",")
	output := math.MaxInt
	for i := 0; i < len(positions); i++ {
		temp := 0
		for _, position := range positions {
			n, _ := strconv.Atoi(position)
			temp += int(math.Abs(float64(i - n)))
		}
		if temp < output {
			output = temp
		}
	}
	return output
}
