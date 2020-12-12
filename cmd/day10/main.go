package main

import (
	"fmt"
	"math"

	"github.com/Acksell/aoc2020/util"
)

const inputFilePath = "../../inputs/adapters.txt"

func differences(adapters util.IntSlice) util.IntSlice {
	adapterSet := make(util.IntSet) // convert slice to set.
	for _, v := range adapters {
		adapterSet[v] = true
	}
	max := util.Max(adapters)
	differences := make(util.IntSlice, 0)
	differences = append(differences, 1)
	for i := 1; i < max; i++ {
		if adapterSet[i] {
			for j := 1; j <= 3; j++ {
				if adapterSet[i+j] {
					differences = append(differences, j)
					break // break inner for loop. continue to next i.
				}
			}
		}
	}
	// append 3 to end since pc can handle 3 higher than max adapter.
	differences = append(differences, 3)
	return differences
}

func combinations(diffs util.IntSlice) int64 {
	tribonacci := []float64{1, 2, 4, 7, 13, 24}
	var counts [7]float64
	streakOf1s := 0
	for _, d := range diffs { // count the lengths of non-overlapping sequences of consecutive 1s.
		if d == 1 {
			streakOf1s++
		} else { // d=3
			counts[streakOf1s]++
			streakOf1s = 0
		}
	}
	c := 1.0
	for i := 1; i < len(tribonacci); i++ {
		c *= math.Pow(tribonacci[i], counts[i+1])
	}
	return int64(c)
}

func main() {
	adapters := make(util.IntSlice, 0)
	util.ReadLines(inputFilePath, &adapters)
	d := differences(adapters)
	fmt.Println(util.Count(d, 1) * util.Count(d, 3)) // part1
	fmt.Println(combinations(d))                     // part2
}
