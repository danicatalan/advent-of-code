package dive

import (
	"math"
	"strconv"
	"strings"
)

//Part2 can take a series of commands like forward 1, down 2, or up 3.
func Part2(input string) int {
	x, y, aim := 0, 0, 0
	steps := strings.Split(input, "\n")
	for _, s := range steps {
		dir := strings.Split(s, " ")[0]
		units, _ := strconv.Atoi(strings.Split(s, " ")[1])
		switch dir {
		case "up":
			aim -= units
		case "down":
			aim += units
		case "forward":
			x += units
			y += units * aim
		}
	}
	return int(math.Abs(float64(x))) * int(math.Abs(float64(y)))
}
