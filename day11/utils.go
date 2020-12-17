package main

import (
	"io/ioutil"
	"log"
)

func getLayout() string {
	file, err := ioutil.ReadFile("data")
	if err != nil {
		log.Fatal(err)
	}

	out := string(file)
	return out
}
