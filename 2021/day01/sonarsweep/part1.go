package sonarsweep

import (
	"math"
	"strconv"
	"strings"
)

//Part1 counts the number of times a depth measurement increases from the previous measurement.
func Part1(input string) int {
	measurements := strings.Split(input, "\n")
	count, prev := 0, math.MaxInt32
	for _, m := range measurements {
		v, _ := strconv.Atoi(m)
		if v > prev {
			count++
		}
		prev = v
	}
	return count
}
