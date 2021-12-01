package passportprocessing

import (
	"strings"
)

// Part1 checks if the provided passport contains all the required fields.
func Part1(pass string) bool {
	required := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}

	for _, f := range required {
		if !strings.Contains(pass, f) {
			return false
		}
	}
	return true
}
