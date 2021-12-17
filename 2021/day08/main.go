package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/danicatalan/advent-of-code/2021/day08/sevensegmentsearch"
)

func main() {
	file, err := ioutil.ReadFile("data")
	if err != nil {
		log.Fatal(err)
	}
	data := string(file)
	r1, r2 := sevensegmentsearch.Part1(data), sevensegmentsearch.Part2(data)
	fmt.Println(r1, r2)
}