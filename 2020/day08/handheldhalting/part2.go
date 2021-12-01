package handheldhalting

import (
	"log"
	"strconv"
	"strings"
)

// Part2 accepts a sequence of instructions and finds the sequence that fixes the infinite loop. Returns the accumulator for that sequence.
func Part2(input string) int {
	acc := 0
	instructions := strings.Split(input, "\n")
	i := 0
	for {
		tempInstructions := make([]string, len(instructions))
		copy(tempInstructions, instructions)
		switch tempInstructions[i][:3] {
		case "jmp":
			tempInstructions[i] = strings.Replace(tempInstructions[i], "jmp", "nop", 1)
		case "nop":
			tempInstructions[i] = strings.Replace(tempInstructions[i], "nop", "jmp", 1)
		}
		finished, tempAcc := run(strings.Join(tempInstructions, "\n"))
		if finished {
			acc = tempAcc
			break
		} else {
			i++
		}
	}
	return acc
}

func run(input string) (bool, int) {
	finished := false
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
		if p >= len(instructions) {
			finished = true
			return finished, acc
		}
	}
	return finished, acc
}
