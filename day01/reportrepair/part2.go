package reportrepair

// Part2 returns the product of the three entries that sum to 2020.
func Part2(expenses []int) int {
	for i := range expenses {
		for j := i + 1; j < len(expenses); j++ {
			for k := j + 1; k < len(expenses); k++ {
				if expenses[i]+expenses[j]+expenses[k] == 2020 {
					return expenses[i] * expenses[j] * expenses[k]
				}
			}
		}
	}
	return 0
}
