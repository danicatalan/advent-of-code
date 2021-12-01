package passwordphilosophy

import (
	"fmt"
	"strings"
)

// Part1 Input gives the password policy and then the password.
// The password policy indicates the lowest and highest number of times
// a given letter must appear for the password to be valid.
func Part1(input string) bool {
	var min, max int
	var letter byte
	var password string
	fmt.Sscanf(input, "%d-%d %c: %v", &min, &max, &letter, &password)

	c := strings.Count(password, string(letter))
	if c >= min && c <= max {
		return true
	}
	return false
}
