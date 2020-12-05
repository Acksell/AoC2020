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
		} else {
			return hgt >= 59 && hgt <= 76
		}
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

// CountValidPassport returns a function that increments a provided counter if it finds a valid passport.
func CountValidPassport(nvalid *int, validate func(string, string) bool) func(string) error {
	validated := 0
	toCount := func(s string) error {
		if len(s) != 0 { // if not blank line, check data.
			linecontent := strings.Split(s, " ")
			for _, data := range linecontent {
				field := data[0:3]
				if required[field] { // only validate required fields.
					if validate(field, data[4:]) { // data[4:] is everything after the colon.
						validated++
					}
				}
			}
		} else { // blank line <=> new passport. reset `validated` and increment count if 7 fields valid.
			if validated == 7 {
				*nvalid++
			}
			validated = 0
		}
		return nil
	}
	return toCount
}

func main() {
	nvalid1 := 0
	nvalid2 := 0
	// dont care about restrictions on data in part1.
	isValidPart1 := func(_ string, _ string) bool { return true }
	util.ReadLines(inputFilePath, CountValidPassport(&nvalid1, isValidPart1))
	util.ReadLines(inputFilePath, CountValidPassport(&nvalid2, isValid))
	fmt.Println(nvalid1)
	fmt.Println(nvalid2)
}
