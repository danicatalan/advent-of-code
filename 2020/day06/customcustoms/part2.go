package customcustoms

import (
	"strings"
)

// Part2 accepts an answer from a group and returns the amount of questions everyone in the group answered in common.
func Part2(answers string) int {
	out := make(map[rune]int)
	p := len(strings.Split(answers, "\n"))
	for _, a := range answers {
		if a >= 97 && a <= 122 {
			out[a]++
		}
	}
	for k, v := range out {
		if v != p {
			delete(out, k)
		}
	}
	return len(out)
}
