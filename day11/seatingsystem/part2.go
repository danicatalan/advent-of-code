package seatingsystem

import (
	"strings"
)

//Part2 accepts a layout of seats and returns the amount of occupied seats when movement has been stabilized between rounds.
func Part2(input string) int {
	s1 := run2(input)
	for {
		s2 := run2(s1)
		if s2 == s1 {
			return strings.Count(s2, "#")
		}
		s1 = s2
	}
}

func run2(grid string) string {
	var m1, m2 [][]string
	rows := strings.Split(grid, "\n")
	for _, r := range rows {
		m1 = append(m1, strings.Split(r, ""))
		m2 = append(m2, strings.Split(r, ""))
	}

	for i, r := range m1 {
		for j, c := range r {
			if c == "." {
				continue
			}
			n := getVisibleNeighbours(i, j, m1)
			if strings.Count(n, "#") == 0 {
				m2[i][j] = "#"
			}
			if strings.Count(n, "#") >= 5 {
				m2[i][j] = "L"
			}
		}
	}

	out := ""
	for _, r := range m2 {
		for _, c := range r {
			out += c
		}
	}
	w := len(m2[0])
	for i := w; i < len(out); i += w + 1 {
		out = out[:i] + "\n" + out[i:]
	}
	return out
}

func getVisibleNeighbours(r, c int, seats [][]string) string {
	n := ""

	dir := [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	for _, d := range dir {
		i, j, stop := r+d[0], c+d[1], false
		for !stop {
			if i >= 0 && j >= 0 && i < len(seats) && j < len(seats[0]) {
				seat := seats[i][j]
				if seat == "#" || seat == "L" {
					n += seat
					stop = true
				}
				i += d[0]
				j += d[1]
			} else {
				stop = true
			}
		}
	}
	return n
}
