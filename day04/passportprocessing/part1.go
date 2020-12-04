package passportprocessing

import (
	"strings"
)

var required = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
}

// Part1 ...
func Part1(pass string) bool {
	for _, f := range required {
		if !strings.Contains(pass, f) {
			return false
		}
	}
	return true
}
