package main

import (
	"fmt"

	"github.com/Acksell/aoc2020/util"
)

const inputFilePath = "../../inputs/expense_report.txt"

var input = make(util.IntSlice, 0)

func init() {
	util.ReadLines(inputFilePath, &input)
}

func main() {
	// result is the result of multiplication of the two expenses
	var resultPart1 int
	var resultPart2 int

	// add it to the set of expenses for constant lookup.
	expenses := make(util.IntSet)
	differences := make([]int, 0)
	for _, expense := range input {
		expenses[expense] = true
		expensePair := (2020 - expense)
		differences = append(differences, expensePair)
		if expenses[expensePair] {
			resultPart1 = expensePair * expense
		}
	}

	for _, diff := range differences {
		for _, expense := range input {
			if expenses[diff-expense] {
				// recover the three numbers and multiply them.
				resultPart2 = (2020 - diff) * expense * (diff - expense)
			}
		}
	}

	fmt.Printf("Part 1 answer: %v\n", resultPart1)
	fmt.Printf("Part 2 answer: %v\n", resultPart2)
}
