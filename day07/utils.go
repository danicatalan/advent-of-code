package main

import (
	"io/ioutil"
	"log"
)

func getRules() string {
	file, err := ioutil.ReadFile("data")
	if err != nil {
		log.Fatal(err)
	}

	return string(file)
}
