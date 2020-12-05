package binaryboarding

// Part2 accepts a list of occupied seats and returns the available one.
func Part2(seats []int) int {
	min, max := seats[0], seats[0]
	list := make(map[int]bool)
	for _, v := range seats {
		list[v] = true
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	for i := min; i < max; i++ {
		if !list[i] {
			return i
		}
	}
	return 0
}
