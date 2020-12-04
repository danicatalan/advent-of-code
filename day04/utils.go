package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func getPassports() []string {
	file, err := ioutil.ReadFile("data")
	if err != nil {
		log.Fatal(err)
	}

	out := strings.Split(string(file), "\n\n")
	return out
}
