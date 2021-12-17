package sevensegmentsearch

import (
	"strings"
)

const DELIMITER = " | "

var UNIQUE = map[int]int{
	2: 1,
	3: 7,
	4: 4,
	7: 8,
}

//Part1 counts instances of digits that use a unique number of segments (1, 4, 7, 8) in the output values.
func Part1(input string) int {
	entries := strings.Split(input, "\n")
	count := 0
	for _, entry := range entries {
		outputValue := strings.Split(entry, DELIMITER)[1]
		for _, digit := range strings.Split(outputValue, " ") {
			if _, found := UNIQUE[len(digit)]; found {
				count++
			}
		}
	}
	return count
}
