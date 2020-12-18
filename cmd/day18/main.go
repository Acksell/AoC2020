package main

import (
	"fmt"
	"strconv"

	"github.com/Acksell/aoc2020/util"
)

const (
	ADD = iota
	MUL
	PAR
)

func PerformOperation(left, right, op int) int {
	switch op {
	case ADD:
		return left + right
	case MUL:
		return left * right
	default:
		return 0
	}
}

func ApplyOperator(values *util.Stack, op int, value int) {
	(*values)[len(*values)-1] = PerformOperation((*values)[len(*values)-1], value, op)
}

func Evaluate(s string) int {
	accumulators := make(util.Stack, 0)
	accumulators.Push(0)
	operators := make(util.Stack, 0)
	operators.Push(ADD)
	for _, c := range s {
		switch c {
		case '(':
			// append new counter to stack when new expression encountered
			accumulators.Push(0)
			operators.Push(ADD)
		case ')':
			value := accumulators.Pop()
			op := operators.Pop()
			ApplyOperator(&accumulators, op, value)
		case '+':
			operators.Push(ADD)
		case '*':
			operators.Push(MUL)
		case ' ':
			continue
		default: // integer value.
			v, _ := strconv.Atoi(string(c))
			op := operators.Pop()
			ApplyOperator(&accumulators, op, v)
		}
	}
	return accumulators.Pop()
}

type Counter struct {
	sum int
}

func (c *Counter) Load(s string) error {
	v := Evaluate(s)
	c.sum += v
	return nil
}

const inputFilePath = "../../inputs/homework.txt"

func main() {
	c := Counter{}
	util.ReadLines(inputFilePath, &c)
	fmt.Println(c.sum) // part 1
}
