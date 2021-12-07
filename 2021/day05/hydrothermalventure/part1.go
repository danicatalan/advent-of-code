package hydrothermalventure

import (
	"regexp"
	"strconv"
	"strings"
)

//Part1 determines the number of overlapping horizontal/vertical lines.
func Part1(input string) int {
	lines := strings.Split(input, "\n")
	re := regexp.MustCompile(`(?P<x1>\d{1,9}),(?P<y1>\d{1,9}) -> (?P<x2>\d{1,9}),(?P<y2>\d{1,9})`)
	diagram := make(map[string]int)
	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		x1, _ := strconv.Atoi(match[1])
		y1, _ := strconv.Atoi(match[2])
		x2, _ := strconv.Atoi(match[3])
		y2, _ := strconv.Atoi(match[4])

		if x1 == x2 || y1 == y2 {
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			if y1 > y2 {
				y1, y2 = y2, y1
			}
			for x := x1; x <= x2; x++ {
				for y := y1; y <= y2; y++ {
					coordinates := strconv.Itoa(x) + "," + strconv.Itoa(y)
					diagram[coordinates]++
				}
			}
		}
	}
	output := 0
	for _, v := range diagram {
		if v >= 2 {
			output++
		}
	}
	return output
}
