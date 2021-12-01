package toboggantrajectory

import (
	"strings"
)

// Part1 counts all the trees you would encounter for the slope right 3, down 1 in a given area.
func Part1(in string) int {
	var out int
	lines := strings.Split(in, "\n")
	w, h := len(lines[0]), len(lines)

	for c, x, y := 0, 0, 0; c < h; c++ {
		if x += 3; x >= w {
			x -= w
		}
		if y++; y >= h {
			break
		}
		if lines[y][x] == '#' {
			out++
		}
	}
	return out
}
