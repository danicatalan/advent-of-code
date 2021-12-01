package main

import (
	"io/ioutil"
	"log"
)

func getMap(path string) string {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	out := string(file)

	return out
}
