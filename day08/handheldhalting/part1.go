package handheldhalting

import (
	"log"
	"strconv"
	"strings"
)

// Part1 accepts a sequence of instructions and returns the accumulator value when the sequence stops due to an infinite loop.
func Part1(input string) int {
	acc := 0
	instructions := strings.Split(input, "\n")
	visited := make(map[int]bool, len(instructions))
	p := 0
	for {
		if visited[p] {
			break
		}
		visited[p] = true
		i := instructions[p][0:3]
		d := instructions[p][4:5]
		v, err := strconv.Atoi(instructions[p][5:])
		if err != nil {
			log.Fatal(err)
		}

		switch i {
		case "nop":
			p++
		case "acc":
			if d == "+" {
				acc += v
			}
			if d == "-" {
				acc -= v
			}
			p++
		case "jmp":
			if d == "+" {
				p += v
			}
			if d == "-" {
				p -= v
			}
		}
	}
	return acc
}
