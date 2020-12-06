package main

import (
	"fmt"

	"github.com/Acksell/aoc2020/util"
)

// ToAnswersCount returns a function which counts total questions *anyone*
// answered yes to and *everyone* answered yes to and increments
// the provided pointers part1 and part2 by their respective amount.
func ToAnswersCount(part1 *int, part2 *int) func(string) error {
	answers := make(map[string]int)
	people := 0
	toCount := func(s string) error {
		if len(s) != 0 {
			for _, c := range s {
				answers[string(c)]++
			}
			people++
		} else {
			for _, v := range answers {
				*part1++
				if v == people { // everyone answered this
					*part2++
				}
			}
			// new group, reset closure variables.
			answers = make(map[string]int)
			people = 0
		}
		return nil
	}
	return toCount
}

const inputFilePath = "../../inputs/answers.txt"

func main() {
	part1 := 0
	part2 := 0
	util.ReadLines(inputFilePath, ToAnswersCount(&part1, &part2))
	fmt.Println(part1)
	fmt.Println(part2)
}
