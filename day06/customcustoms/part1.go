package customcustoms

// Part1 accepts an answer from a group and returns the total amount of questions answered by the people in that group.
func Part1(answer string) int {
	out := make(map[rune]bool)
	for _, a := range answer {
		if a >= 97 && a <= 122 {
			out[a] = true
		}
	}
	return len(out)
}
