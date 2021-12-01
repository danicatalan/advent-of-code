package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func getBoardingPasses() []string {
	file, err := ioutil.ReadFile("data")
	if err != nil {
		log.Fatal(err)
	}

	out := strings.Split(string(file), "\n")
	return out
}
