package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Acksell/aoc2020/util"
)

// Password contains the restrictions on a single password and the password itself.
type Password struct {
	min  int
	max  int
	char byte
	pwd  string
}

// IsValidOld returns true if password is valid according to the old policy.
func (p Password) IsValidOld() bool {
	var count int
	for _, c := range p.pwd {
		if c == rune(p.char) {
			count++
		}
	}
	return p.min <= count && p.max >= count
}

// IsValid returns true if password is valid.
func (p Password) IsValid() bool {
	// xor
	return (p.pwd[p.min-1] == p.char) != (p.pwd[p.max-1] == p.char)
}

// NewPassword parses the input string and returns a Password struct.
func NewPassword(s string) (Password, error) {
	contents := strings.Split(s, " ")
	minmax := strings.Split(contents[0], "-")
	min, err := strconv.Atoi(minmax[0])
	if err != nil {
		return Password{}, err
	}
	max, err := strconv.Atoi(minmax[1])
	if err != nil {
		return Password{}, err
	}
	char := contents[1][0]
	pwd := contents[2]
	return Password{min, max, char, pwd}, nil
}

// Passwords is a list of passwords.
type Passwords []Password

// Load constructs a Password object and appends it to Passwords. Used in ReadLines.
func (slice *Passwords) Load(input string) error {
	pwd, err := NewPassword(input)
	if err != nil {
		return err
	}
	*slice = append(*slice, pwd)
	return nil

}

const inputFilePath = "../../inputs/passwords.txt"

var passwords = make(Passwords, 0, 100)

func init() {
	util.ReadLines(inputFilePath, &passwords)
}

func main() {
	var oldValid uint
	var valid uint
	for _, p := range passwords {
		if p.IsValidOld() {
			oldValid++
		}
		if p.IsValid() {
			valid++
		}
	}
	fmt.Println(oldValid)
	fmt.Println(valid)
}
