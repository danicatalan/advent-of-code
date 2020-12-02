package passwordphilosophy

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
