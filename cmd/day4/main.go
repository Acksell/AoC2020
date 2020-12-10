package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Acksell/aoc2020/util"
)

const inputFilePath = "../../inputs/passports.txt"

var required = util.NewStringSet("byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid")

func isValid(field string, value string) bool {
	switch field {
	case "cid":
		return true
	case "byr":
		byr, _ := strconv.Atoi(value) // ignore error because data allows it this time.
		return len(value) == 4 && byr >= 1920 && byr <= 2002
	case "iyr":
		iyr, _ := strconv.Atoi(value) // ignore error because data allows it this time.
		return len(value) == 4 && iyr >= 2010 && iyr <= 2020
	case "eyr":
		eyr, _ := strconv.Atoi(value) // ignore error because data allows it this time.
		return len(value) == 4 && eyr >= 2020 && eyr <= 2030
	case "hgt":
		hgt, _ := strconv.Atoi(value[:len(value)-2]) // ignore error because data allows it this time.
		if value[len(value)-2:] == "cm" {
			return hgt >= 150 && hgt <= 193
		}
		return hgt >= 59 && hgt <= 76
	case "hcl":
		if value[0] != '#' {
			return false
		}
		accepted := util.NewStringSet("0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f")
		for _, c := range value[1:] {
			if !accepted[string(c)] {
				return false
			}
		}
		return true
	case "ecl":
		accepted := util.NewStringSet("amb", "blu", "brn", "gry", "grn", "hzl", "oth")
		return accepted[value]
	case "pid":
		_, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		return len(value) == 9
	}
	return true
}

// ValidCounter represent a counter. Type implements Loadable.
type ValidCounter struct {
	validated1 int
	validated2 int
	part1      int
	part2      int
}

// Load returns a function that increments a provided counter if it finds a valid passport.
func (c *ValidCounter) Load(s string) error {
	if len(s) != 0 { // if not blank line, check data.
		linecontent := strings.Split(s, " ")
		for _, data := range linecontent {
			field := data[0:3]
			if required[field] { // only validate required fields.
				c.validated1++
				if isValid(field, data[4:]) { // data[4:] is everything after the colon.
					c.validated2++
				}
			}
		}
	} else { // blank line <=> new passport. reset `validated` and increment count if 7 fields valid.
		if c.validated1 == 7 {
			c.part1++
		}
		if c.validated2 == 7 {
			c.part2++
		}
		c.validated1 = 0
		c.validated2 = 0
	}
	return nil
}

func main() {
	c := ValidCounter{}
	util.ReadLines(inputFilePath, &c)
	fmt.Println(c.part1)
	fmt.Println(c.part2)
}
