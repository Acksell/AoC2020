package main

import (
	"errors"
	"fmt"

	"github.com/Acksell/aoc2020/util"
)

func FindError(ints util.IntSlice) (int, error) {
	// initialize the sums with the premable of length 25.
	previous := make([]util.IntSet, 0)
	for i := range ints[:25] {
		partial := make(util.IntSet)
		for j := range ints[:25] {
			partial[ints[i]+ints[j]] = true
		}
		previous = append(previous, partial)
	}

	for i, k := range ints[25:] {
		// check that it is a sum of previous 25.
		var any bool
		for _, sumset := range previous {
			if sumset[k] {
				any = true
				break
			}
		}
		if !any { // not a sum of any of the previous 25. Return k.
			return k, nil
		}
		// update previous 25 to include k.
		previous = previous[1:]
		newsums := make(util.IntSet)
		for _, v := range ints[i : 25+i] {
			newsums[v+k] = true
		}
		previous = append(previous, newsums)
	}
	return 0, errors.New("No solution found")
}

// FindContiguous recursively finds a contiguous sequence of numbers of length smaller
// than or equal to n that sum to target. If there is one it returns true and the
// corresponding sequence. Otherwise return false and all the contiguous sums of size n.
func FindContiguous(ints util.IntSlice, n int, target int) (bool, util.IntSlice) {
	if n == 1 {
		return false, ints
	}
	found, sums := FindContiguous(ints, n-1, target)
	if found {
		return true, sums
	}
	next := make(util.IntSlice, 0)
	for i, v := range ints[n-1:] {
		sum := sums[i] + v
		if sum == target {
			return true, ints[i : n+i]
		}
		next = append(next, sum)
	}
	return false, next
}

const inputFilePath = "../../inputs/encodings.txt"

func main() {
	ints := make(util.IntSlice, 0)
	util.ReadLines(inputFilePath, &ints)
	// part 1
	i, _ := FindError(ints)
	fmt.Println(i)
	// part 2
	// 40 is just chosen as the max n we use, just to avoid infinite recursion.
	// Answer sequence turns out to be of length 17.
	found, contiguous := FindContiguous(ints, 40, i)
	if found {
		// print part 2 answer.
		max, _ := util.Max(contiguous)
		min, _ := util.Min(contiguous)
		fmt.Println(max + min)
	}
}
