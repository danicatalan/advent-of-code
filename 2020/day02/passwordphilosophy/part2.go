package passwordphilosophy

import (
	"fmt"
)

// Part2 Each policy actually describes two positions in the password,
// where 1 means the first character, 2 means the second character, and so on.
// Exactly one of these positions must contain the given letter.
// Other occurrences of the letter are irrelevant for the purposes of policy enforcement.
func Part2(input string) bool {
	var pos1, pos2 int
	var letter byte
	var password string
	fmt.Sscanf(input, "%d-%d %c: %v", &pos1, &pos2, &letter, &password)

	if (letter == password[pos1-1] || letter == password[pos2-1]) && (password[pos1-1] != password[pos2-1]) {
		return true
	}
	return false
}
