package util

func Max(ints []int) int {
	max := ints[0]
	for _, v := range ints {
		if v > max {
			max = v
		}
	}
	return max
}

func Min(ints []int) int {
	min := ints[0]
	for _, v := range ints {
		if v < min {
			min = v
		}
	}
	return min
}
