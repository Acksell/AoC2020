package util

func Max(ints []int) (int, int) {
	max := ints[0]
	idx := 0
	for i, v := range ints {
		if v > max {
			max = v
			idx = i
		}
	}
	return max, idx
}

func Min(ints []int) (int, int) {
	min := ints[0]
	idx := 0
	for i, v := range ints {
		if v < min {
			min = v
			idx = i
		}
	}
	return min, idx
}

func Count(ints []int, target int) int {
	c := 0
	for _, v := range ints {
		if v == target {
			c++
		}
	}
	return c
}
