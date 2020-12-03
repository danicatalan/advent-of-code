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

// Part2_1 counts all the trees you would encounter for the slope right `sx`, down `sy` in a given area.
func Part2_1(in string, sx, sy int) int {
	var out int
	lines := strings.Split(in, "\n")
	w, h := len(lines[0]), len(lines)

	for c, x, y := 0, 0, 0; c < h; c++ {
		if x += sx; x >= w {
			x -= w
		}
		if y += sy; y >= h {
			break
		}
		if lines[y][x] == '#' {
			out++
		}
	}
	return out
}
