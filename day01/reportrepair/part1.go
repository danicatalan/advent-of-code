package reportrepair

// Part1 returns the product of the two entries that sum to 2020.
func Part1(expenses []int) int {
	for i := range expenses {
		for j := i + 1; j < len(expenses); j++ {
			if expenses[i]+expenses[j] == 2020 {
				return expenses[i] * expenses[j]
			}
		}
	}
	return 0
}
