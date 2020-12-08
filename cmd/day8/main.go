package main

import (
	"strconv"
	"strings"

	"github.com/Acksell/aoc2020/util"
)

type Op int

const (
	NOP Op = iota // Zero value is NOP.
	ACC
	JMP
)

type Instruction struct {
	op  Op
	arg int
}

func ParseInstruction(s string) (Instruction, error) {
	instr := strings.Split(s, " ")
	var op Op
	arg, err := strconv.Atoi(instr[1])
	if err != nil {
		return Instruction{}, nil
	}
	switch instr[0] {
	case "nop":
		op = NOP
	case "acc":
		op = ACC
	case "jmp":
		op = JMP
	}
	return Instruction{op, arg}, nil
}

const inputFilePath = "../../inputs/boot_code.txt"

func main() {
	load := func(i *[]Instruction) func(string) error {
		return func(s string) error {
			instr, err := ParseInstruction(s)
			*i = append(*i, instr)
			return err
		}
	}
	instructions := make([]Instruction, 0)
	util.ReadLines(inputFilePath, load(&instructions))
}
