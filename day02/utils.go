package main

import (
	"bufio"
	"log"
	"os"
)

func getPasswords(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var out []string
	s := bufio.NewScanner(file)
	for s.Scan() {
		line := s.Text()
		out = append(out, line)
	}

	return out
}
