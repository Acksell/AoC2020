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

func EvaluatePrecendence(s string) int { // ADD precedence over MUL
	accumulators := make(util.Stack, 0)
	accumulators.Push(0)
	operators := make(util.Stack, 0)
	operators.Push(ADD)
	for i, c := range s {
		switch c {
		case '(':
			// append new counter to stack when new expression encountered
			accumulators.Push(0)
			operators.Push(PAR)
			operators.Push(ADD)
		case ')':
			for { // Perform operations until we find the opening parenthesis.
				if operators.Top() == PAR {
					// remove subsequent closing parenthesis
					operators.Pop()
					if operators.Top() == ADD { // if operation is add, just add immediately.
						v := accumulators.Pop()
						ApplyOperator(&accumulators, operators.Pop(), v)
					}
					break
				}
				value := accumulators.Pop()
				op := operators.Pop()
				ApplyOperator(&accumulators, op, value)
			}
		case '+':
			operators.Push(ADD)
		case '*':
			operators.Push(MUL)
		case ' ':
			continue
		default: // integer value.
			v, _ := strconv.Atoi(string(c))
			// If highest precedence operator is on top of stack
			// apply the operator to the current value v and the next accumulator.
			if operators.Top() == ADD { // add has highest precedence
				op := operators.Pop()
				ApplyOperator(&accumulators, op, v)
			} else if i == len(s)-1 { // if at the end of string dont push it for later processing, just do it now.
				op := operators.Pop()
				ApplyOperator(&accumulators, op, v)
			} else {
				accumulators.Push(v)
			}
		}
	}
	for len(accumulators) > 1 { // Evaluate all remaining until only one left.
		v := accumulators.Pop()
		op := operators.Pop()
		ApplyOperator(&accumulators, op, v)
	}
	return accumulators.Pop() // return last accumulator.
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
	sum      int
	evaluate func(s string) int
}

func (c *Counter) Load(s string) error {
	v := c.evaluate(s)
	c.sum += v
	return nil
}

const inputFilePath = "../../inputs/homework.txt"

func main() {
	part1 := Counter{0, Evaluate}
	part2 := Counter{0, EvaluatePrecendence}
	util.ReadLines(inputFilePath, &part1)
	util.ReadLines(inputFilePath, &part2)
	fmt.Println(part1.sum) // part 1
	fmt.Println(part2.sum) // part 2
}
