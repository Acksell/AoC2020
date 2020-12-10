package main

import (
	"fmt"

	"github.com/Acksell/aoc2020/util"
)

// AnswerCounter counts answers
type AnswerCounter struct {
	answers map[string]int
	people  int
	part1   int
	part2   int
}

func (c *AnswerCounter) Load(s string) error {
	if len(s) != 0 {
		for _, answer := range s {
			c.answers[string(answer)]++
		}
		c.people++
	} else {
		for _, v := range c.answers {
			c.part1++
			if v == c.people { // everyone answered this
				c.part2++
			}
		}
		// new group, reset closure variables.
		c.answers = make(map[string]int)
		c.people = 0
	}
	return nil
}

const inputFilePath = "../../inputs/answers.txt"

func main() {
	answers := make(map[string]int)
	c := AnswerCounter{answers, 0, 0, 0}
	util.ReadLines(inputFilePath, &c)
	fmt.Println(c.part1)
	fmt.Println(c.part2)
}
