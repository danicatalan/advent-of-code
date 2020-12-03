package toboggantrajectory

import (
	"strings"
)

// Part2 multiplies together the number of trees encountered on each of the given slopes.
func Part2(in string, slopes [][]int) int {
	out := 1
	for _, s := range slopes {
		out *= Part2_1(in, s[0], s[1])
	}
	return out
}

// Part2_1 counts all the trees you would encounter for the slope right `sx`, down `sy` in a given map.
func Part2_1(in string, sx, sy int) int {
	var out int
	lines := strings.Split(in, "\n")
	w, h := len(lines[0]), len(lines)

	c, x, y := 0, 0, 0
	for c < h {
		x += sx
		if x >= w {
			x -= w
		}
		y += sy
		if y < h {
			if lines[y][x] == byte('#') {
				out++
			}
		}
		c++
	}
	return out
}
