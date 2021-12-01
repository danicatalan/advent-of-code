package main

import (
	"io/ioutil"
	"log"
)

func getJolgates() string {
	file, err := ioutil.ReadFile("data")
	if err != nil {
		log.Fatal(err)
	}

	out := string(file)
	return out
}
