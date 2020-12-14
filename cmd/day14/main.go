package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Acksell/aoc2020/util"
)

const BitmaskSize = 36

type Bitmask struct {
	mask string
	one  int // bits to OR with
	zero int // bits to AND with
}

func NewBitmask(s string) Bitmask {
	one := 0
	zero := 0
	for i := 0; i < BitmaskSize; i++ {
		switch s[i] {
		case 'X':
			// one += 0 << (BitmaskSize - 1 - i)  // stay the same, d OR 0 gives d.
			zero += 1 << (BitmaskSize - 1 - i) // stay the same, d AND 1 gives d.
		case '1':
			one += 1 << (BitmaskSize - 1 - i)  // set to one
			zero += 1 << (BitmaskSize - 1 - i) // stay the same, d AND 1 gives d.
		case '0':
			continue
			// zero += 0 << (BitmaskSize - 1 - i) // set to zero
			// one += 0 << (BitmaskSize - 1 - i)  // stay the same, d OR 0 gives d.
		}
	}
	return Bitmask{s, one, zero}
}

func (b Bitmask) SetOnes(i int) int {
	return i | b.one // set one bits to one.
}

func (b Bitmask) SetZeros(i int) int {
	return i & b.zero // set zero bits to zero.
}

func (b Bitmask) Apply(i int) int {
	res := i
	res = b.SetZeros(res) // set zero bits to zero.
	res = b.SetOnes(res)  // set one bits to one.
	return res
}

type Memory map[int]int

type WritableMemory interface {
	Write(int, int, Bitmask)
	Read(int) int
	Get() Memory
}

// StandardMemory is a writable memory.
type StandardMemory struct {
	mem Memory
}

func (sm *StandardMemory) Write(address, value int, bitmask Bitmask) {
	sm.mem[address] = bitmask.Apply(value)
}

func (sm StandardMemory) Read(address int) int {
	return sm.mem[address] // no error checks for now.
}

func (sm StandardMemory) Get() Memory {
	return sm.mem
}

// Program contains the instructions, memory and the current bitmask.
type Program struct {
	bitmask      Bitmask
	memory       WritableMemory
	instructions []string
}

func (p *Program) SetBitmask(b Bitmask) {
	p.bitmask = b
}

func (p *Program) SetMemory(address, value int) {
	p.memory.Write(address, value, p.bitmask)
}

func (p *Program) Execute() error {
	for _, i := range p.instructions {
		contents := strings.Split(i, " = ")
		if len(contents) == 0 {
			continue
		}
		if contents[0] == "mask" {
			p.SetBitmask(NewBitmask(contents[1]))
		} else {
			addr := strings.Trim(contents[0][3:], "[]") // get the address inside brackets.
			address, err := strconv.Atoi(addr)
			value, err := strconv.Atoi(contents[1])
			if err != nil {
				return err
			}
			p.SetMemory(address, value)
		}
	}
	return nil
}

func NewProgram(m WritableMemory) Program {
	instr := make([]string, 0)
	return Program{Bitmask{}, m, instr}
}

// Part 2

// FloatingMemory is a StandardMemory memory with a special write function.
type FloatingMemory struct {
	mem Memory
}

func (fm *FloatingMemory) Write(address, value int, bitmask Bitmask) {
	floating := make([]int, 0) // "floating" bits
	for i, b := range bitmask.mask {
		if b == 'X' {
			floating = append(floating, 1<<(BitmaskSize-1-i))
		}
	}
	addresses := make([]int, 0)
	addresses = append(addresses, bitmask.SetOnes(address)) // set the ones to one (OR operation with appropriate mask)
	fm.mem[bitmask.SetOnes(address)] = value
	for _, f := range floating {
		for _, addr := range addresses {
			faddress := addr ^ f // Append address XOR with the float.
			addresses = append(addresses, faddress)
			fm.mem[faddress] = value
		}
	}
}

func (fm FloatingMemory) Read(address int) int {
	return fm.mem[address] // no error checks for now.
}

func (fm FloatingMemory) Get() Memory {
	return fm.mem
}

// Load loads the instructions to the program.
func (p *Program) Load(s string) error {
	p.instructions = append(p.instructions, s)
	return nil
}

const inputFilePath = "../../inputs/memory.txt"

func SumMemory(p Program) int {
	s := 0
	for _, value := range p.memory.Get() {
		s += value
	}
	return s
}

func main() {
	// part 1
	mem := StandardMemory{make(Memory)} // should really provide a "new method" here.
	program := NewProgram(&mem)
	util.ReadLines(inputFilePath, &program) // load program.

	program.Execute()
	fmt.Println(SumMemory(program))

	// part 2
	fmem := FloatingMemory{make(Memory)} // should really provide a "new method" here.
	program = NewProgram(&fmem)
	util.ReadLines(inputFilePath, &program) // load program.

	program.Execute()
	fmt.Println(SumMemory(program))
}
