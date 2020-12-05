package binaryboarding

import (
	"math"
)

// Part1 accepts a boarding pass and returns the seat ID.
func Part1(bp string) int {
	rows := []float64{0, 127}
	for _, c := range bp[:7] {
		diff := rows[1] - rows[0]
		if c == 'F' {
			rows[1] -= math.Ceil(diff / 2)
		}
		if c == 'B' {
			rows[0] += math.Ceil(diff / 2)
		}
	}

	cols := []float64{0, 7}
	for _, c := range bp[7:] {
		diff := cols[1] - cols[0]
		if c == 'L' {
			cols[1] -= math.Ceil(diff / 2)
		}
		if c == 'R' {
			cols[0] += math.Ceil(diff / 2)
		}
	}

	return int(rows[0]*8 + cols[0])
}
