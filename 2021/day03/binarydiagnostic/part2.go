package binarydiagnostic

import (
	"strconv"
	"strings"
)

//Part2 calculates the oxygen generator rating by the CO2 scrubber rating.
func Part2(input string) int {
	report := strings.Split(input, "\n")
	g, e := report, report

	i := 0
	for len(g) > 1 {
		zeros, ones := 0, 0
		startsWithZeros, startsWithOnes := []string{}, []string{}
		for j := 0; j < len(g); j++ {
			n, _ := strconv.Atoi(g[j][i : i+1])
			if n == 0 {
				zeros++
				startsWithZeros = append(startsWithZeros, g[j])
			}
			if n == 1 {
				ones++
				startsWithOnes = append(startsWithOnes, g[j])
			}
		}

		if ones >= zeros {
			g = startsWithOnes
		} else {
			g = startsWithZeros
		}
		i++
	}

	i = 0
	for len(e) > 1 {
		zeros, ones := 0, 0
		startsWithZeros, startsWithOnes := []string{}, []string{}
		for j := 0; j < len(e); j++ {
			n, _ := strconv.Atoi(e[j][i : i+1])
			if n == 0 {
				zeros++
				startsWithZeros = append(startsWithZeros, e[j])
			}
			if n == 1 {
				ones++
				startsWithOnes = append(startsWithOnes, e[j])
			}
		}

		if ones < zeros {
			e = startsWithOnes
		} else {
			e = startsWithZeros
		}
		i++
	}

	gd, _ := strconv.ParseInt(g[0], 2, 64)
	ed, _ := strconv.ParseInt(e[0], 2, 64)
	return int(gd * ed)
}
