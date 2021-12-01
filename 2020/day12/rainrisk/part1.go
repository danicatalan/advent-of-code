package rainrisk

import (
	"log"
	"math"
	"strconv"
	"strings"
)

//Part1 ...
func Part1(input string) int {
	instructions := strings.Split(input, "\n")
	p, m, d := []int{0, 0}, []int{0, 0}, []int{1, 0}
	for _, i := range instructions {
		a := i[:1]
		v, err := strconv.Atoi(i[1:])
		if err != nil {
			log.Fatal(err)
		}
		switch a {
		case "N":
			m = []int{0, v * -1}
		case "S":
			m = []int{0, v * 1}
		case "E":
			m = []int{v * 1, 0}
		case "W":
			m = []int{v * -1, 0}
		case "L":
			for i := 0; i < v/90; i++ {
				d = []int{d[1], -d[0]}
			}
		case "R":
			for i := 0; i < v/90; i++ {
				d = []int{-d[1], d[0]}
			}
		case "F":
			m = []int{v * d[0], v * d[1]}
		}
		if a != "L" && a != "R" {
			p[0] += m[0]
			p[1] += m[1]
		}
	}
	return int(math.Abs(float64(p[0])) + math.Abs(float64(p[1])))
}
