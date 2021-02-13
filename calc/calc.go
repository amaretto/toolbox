package calc

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxInts(a []int) int {
	// Set max value
	var max int
	for _, n := range a {
		if n > max {
			max = n
		}
	}
	return max
}

func minInts(a []int) int {
	// Set min value
	var min int
	for _, n := range a {
		if n < min {
			min = n
		}
	}
	return min
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
