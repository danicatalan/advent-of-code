package main

	import (
		"fmt"
		"io/ioutil"
		"log"

		"github.com/danicatalan/advent-of-code/2021/day03/binarydiagnostic"
	)

	func main() {
		file, err := ioutil.ReadFile("data")
		if err != nil {
			log.Fatal(err)
		}
		data := string(file)
		r1, r2 := binarydiagnostic.Part1(data), binarydiagnostic.Part2(data)
		fmt.Println(r1, r2)
	}