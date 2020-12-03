package main

import (
	"fmt"

	"github.com/Acksell/aoc2020/util"
)

// Forest stores a slice of array of trees
type Forest [][]bool

var forest Forest

// ToForest returns a function that initializes trees to the forest given a string.
func ToForest(forest *Forest) func(string) error {
	toForest := func(s string) error {
		trees := make([]bool, 31)
		for i, c := range s {
			if c == '#' {
				trees[i] = true
			}
		}
		*forest = append(*forest, trees)
		return nil
	}
	return toForest
}

const inputFilePath = "../../inputs/forest.txt"

func init() {
	util.ReadLines(inputFilePath, ToForest(&forest))
}

func traverse(right int, down int) int {
	count := 0
	for i, j := 0, 0; i < len(forest); i, j = i+down, j+right {
		if forest[i][j%31] {
			count++
		}
	}
	return count
}

func main() {
	result1 := traverse(3, 1)
	result2 := traverse(1, 1) * traverse(3, 1) * traverse(5, 1) * traverse(7, 1) * traverse(1, 2)
	fmt.Println(result1, result2)
}
