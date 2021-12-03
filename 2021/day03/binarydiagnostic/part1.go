package binarydiagnostic

import (
	"strconv"
	"strings"
)

//Part1 calculates the gamma rate and epsilon rate.
func Part1(input string) int {
	g, e := "", ""
	report := strings.Split(input, "\n")
	for i := 0; i < len(report[0]); i++ {
		zeros, ones := 0, 0
		for j := 0; j < len(report); j++ {
			n, _ := strconv.Atoi(report[j][i : i+1])
			if n == 0 {
				zeros++
			}
			if n == 1 {
				ones++
			}
		}

		if ones > zeros {
			g += "1"
			e += "0"
		} else {
			g += "0"
			e += "1"
		}
	}
	gd, _ := strconv.ParseInt(g, 2, 64)
	ed, _ := strconv.ParseInt(e, 2, 64)
	return int(gd * ed)
}
