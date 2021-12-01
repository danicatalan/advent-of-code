package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/danicatalan/advent-of-code/2021/day01/sonarsweep"
)

func main() {
	instructions := getData()
	r1, r2 := sonarsweep.Part1(instructions), sonarsweep.Part2(instructions)
	fmt.Println(r1, r2)
}

func getData() string {
	file, err := ioutil.ReadFile("data")
	if err != nil {
		log.Fatal(err)
	}

	return string(file)
}
