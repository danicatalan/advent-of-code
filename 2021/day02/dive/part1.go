package dive

import (
	"math"
	"strconv"
	"strings"
)

//Part1 can take a series of commands like forward 1, down 2, or up 3.
func Part1(input string) int {
	x, y := 0, 0
	steps := strings.Split(input, "\n")
	for _, s := range steps {
		dir := strings.Split(s, " ")[0]
		units, _ := strconv.Atoi(strings.Split(s, " ")[1])
		switch dir {
		case "up":
			y += units
		case "down":
			y -= units
		case "forward":
			x += units
		}
	}
	return int(math.Abs(float64(x))) * int(math.Abs(float64(y)))
}
