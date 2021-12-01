package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func getExpenses(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var out []int
	s := bufio.NewScanner(file)
	for s.Scan() {
		line, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatal(err)
		}
		out = append(out, line)
	}

	return out
}
