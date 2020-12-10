package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Acksell/aoc2020/util"
)

// Op represents an instruction.
type Op int

const (
	// NOP Does nothing. This is the zero value.
	NOP Op = iota
	// ACC increments the global accumulator
	ACC
	// JMP jumps to another instruction.
	JMP
)

// Instruction is the combination of an operation and an argument
type Instruction struct {
	op  Op
	arg int
}

// Code is a list of instructions
type Code []Instruction

// State contains the global accumulator Acc, some Code and an interger Head
//  representing which instruciton in the Code we should execute next.
type State struct {
	Acc  int
	Head int // current instruction executing
	Code Code
}

// NewState returns a fresh state given some code.
func NewState(code Code) State {
	return State{0, 0, code}
}

// ParseInstruction returns an Instruction given the written boot code string.
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

// Next returns false if there is no next instruction, true otherwise.
func (s *State) Next() bool {
	s.Execute()
	if s.Head > len(s.Code) {
		fmt.Println("EOF")
		return false
	}
	return true
}

// Execute executes the current instruction and updates the State accordingly.
func (s *State) Execute() {
	i := s.Code[s.Head]
	switch i.op {
	case NOP:
		s.Head++
		return
	case ACC:
		s.Head++
		s.Acc += i.arg
	case JMP:
		s.Head += i.arg
		return
	}
}

// Load some instructions from a file.
func (c *Code) Load(s string) error {
	instr, err := ParseInstruction(s)
	*c = append(*c, instr)
	return err
}

func (s State) String() string {
	return fmt.Sprintf("{Acc %v Head %v}", s.Acc, s.Head)
}

const inputFilePath = "../../inputs/boot_code.txt"

func main() {
	code := make(Code, 0)
	util.ReadLines(inputFilePath, &code)

	state := NewState(code)
	state.Execute()
	linesRead := make(util.IntSet)
	for state.Next() {
		// fmt.Println(state)
		if linesRead[state.Head] { // Already read, means we are in a loop.
			break
		}
		linesRead[state.Head] = true
	}
	fmt.Println(state.Acc)
}
