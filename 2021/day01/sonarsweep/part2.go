package sonarsweep

import (
	"math"
	"strconv"
	"strings"
)

//Part2 counts the number of times the sum of measurements in this sliding window increases from the previous sum.
func Part2(input string) int {
	measurements := strings.Split(input, "\n")
	count, prev := 0, math.MaxInt32
	for i := 0; i <= len(measurements)-3; i++ {
		m1, _ := strconv.Atoi(measurements[i])
		m2, _ := strconv.Atoi(measurements[i+1])
		m3, _ := strconv.Atoi(measurements[i+2])
		v := m1 + m2 + m3
		if v > prev {
			count++
		}
		prev = v
	}
	return count
}
