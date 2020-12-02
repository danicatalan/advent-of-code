package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

// Part1 Input gives the password policy and then the password.
// The password policy indicates the lowest and highest number of times
// a given letter must appear for the password to be valid.
func Part1(input string) bool {
	regex := regexp.MustCompile(`-| |: `)
	p := regex.Split(input, -1)

	c := strings.Count(p[3], p[2])

	min, err := strconv.Atoi(p[0])
	if err != nil {
		log.Fatal(err)
	}
	max, err := strconv.Atoi(p[1])
	if err != nil {
		log.Fatal(err)
	}

	if c >= min && c <= max {
		return true
	}
	return false
}

// Part2 Each policy actually describes two positions in the password,
// where 1 means the first character, 2 means the second character, and so on.
// Exactly one of these positions must contain the given letter.
// Other occurrences of the letter are irrelevant for the purposes of policy enforcement.
func Part2(input string) bool {
	regex := regexp.MustCompile(`-| |: `)
	p := regex.Split(input, -1)

	p1, err := strconv.Atoi(p[0])
	if err != nil {
		log.Fatal(err)
	}
	p2, err := strconv.Atoi(p[1])
	if err != nil {
		log.Fatal(err)
	}

	if p[2] == string(p[3][p1-1]) && p[2] != string(p[3][p2-1]) {
		return true
	}
	if p[2] != string(p[3][p1-1]) && p[2] == string(p[3][p2-1]) {
		return true
	}
	return false
}
