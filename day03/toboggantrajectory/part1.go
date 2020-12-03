package toboggantrajectory

import (
	"strings"
)

// Part1 counts all the trees you would encounter for the slope right 3, down 1 in a given map.
func Part1(in string) int {
	var out int
	lines := strings.Split(in, "\n")
	w, h := len(lines[0]), len(lines)

	x := 0
	for y := range lines {
		x += 3
		if x >= w {
			x -= w
		}
		y++
		if y < h {
			if lines[y][x] == byte('#') {
				out++
			}
		}
	}
	return out
}
