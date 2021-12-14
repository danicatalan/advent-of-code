package thetreacheryofwhales

import (
	"math"
	"strconv"
	"strings"
)

//Part2 calculates the horizontal position that the crabs can align to using the least fuel possible (incremental fuel consumption).
func Part2(input string) int {
	positions := strings.Split(input, ",")
	output := math.MaxInt
	for i := 0; i < len(positions); i++ {
		temp := 0
		for _, position := range positions {
			n, _ := strconv.Atoi(position)
			steps := int(math.Abs(float64(i - n)))
			temp += getFuel(steps)
		}
		if temp < output {
			output = temp
		}
	}
	return output
}

func getFuel(steps int) int {
	output := 0
	for i := 0; i <= steps; i++ {
		output += i
	}
	return output
}
