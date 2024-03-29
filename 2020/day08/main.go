package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/danicatalan/advent-of-code/day08/handheldhalting"
)

func main() {
	instructions := getData()
	r1, r2 := handheldhalting.Part1(instructions), handheldhalting.Part2(instructions)
	fmt.Println(r1, r2)
}

func getData() string {
	file, err := ioutil.ReadFile("data")
	if err != nil {
		log.Fatal(err)
	}

	return string(file)
}
