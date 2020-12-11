package main

import (
	"errors"
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
	if s.Head >= len(s.Code) {
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

// Reset to initial state.
func (s *State) Reset() {
	s.Acc = 0
	s.Head = 0
}

// Load instructions from a file.
func (c *Code) Load(s string) error {
	instr, err := ParseInstruction(s)
	*c = append(*c, instr)
	return err
}

func (s State) String() string {
	return fmt.Sprintf("{Acc %v Head %v}", s.Acc, s.Head)
}

const inputFilePath = "../../inputs/boot_code.txt"

// Run the state code until EOF or error encountered.
func (s State) Run() (int, error) {
	linesRead := make(util.IntSet)
	for s.Next() {
		if linesRead[s.Head] { // Already read, means we are in a loop.
			return s.Acc, errors.New("Infinite loop encountered")
		}
		linesRead[s.Head] = true
	}
	return s.Acc, nil
}

// RepairState changes single NOPs or JMP instructions until the program executes.
// Note: Can use backtrack algorithm but I chose to just simply run the entire code again to check.
func RepairState(state *State) (int, error) {
	linesRead := make(util.IntSet)
	instrToSwitch := make([]int, 0)
	for state.Next() {
		if linesRead[state.Head] { // Already read, means we are in a loop.
			break
		}
		// record where the nops and jmps are.
		if op := state.Code[state.Head].op; op == NOP || op == JMP {
			instrToSwitch = append(instrToSwitch, state.Head)
		}
		linesRead[state.Head] = true
	}

	// loop over the previously touched NOPs and JMPS and switch one by one. If it runs, then return acc.
	for _, line := range instrToSwitch {
		state.Reset()
		switch op := state.Code[line].op; op { // switch the NOP for JMP or JMP for NOP.
		case NOP:
			state.Code[line].op = JMP
		case JMP:
			state.Code[line].op = NOP
		}
		acc, err := state.Run()
		if err != nil { // switch back to what it was before.
			switch op := state.Code[line].op; op {
			case NOP:
				state.Code[line].op = JMP
			case JMP:
				state.Code[line].op = NOP
			}
		} else {
			return acc, nil
		}
	}
	return 0, errors.New("No solution found")
}

func main() {
	code := make(Code, 0)
	util.ReadLines(inputFilePath, &code)
	state := NewState(code)
	acc, _ := state.Run()
	fmt.Println(acc)
	acc, err := RepairState(&state)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(acc)
}
