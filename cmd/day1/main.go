package main

import (
	"fmt"
	"strconv"

	"github.com/Acksell/aoc2020/util"
)

const inputFilePath = "../../inputs/expense_report.txt"

type Int64Slice []uint64

// Load the input to an Int64Slice
func (i *Int64Slice) Load(input string) error {
	u64, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		return err
	}
	*i = append(*i, u64)
	return nil
}

var input = make(Int64Slice, 0)

func init() {
	util.ReadLines(inputFilePath, &input)
}

func main() {
	// result is the result of multiplication of the two expenses
	var resultPart1 uint64
	var resultPart2 uint64

	// add it to the set of expenses for constant lookup.
	expenses := make(util.Int64Set)
	differences := make([]uint64, 0)
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
