package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/Acksell/aoc2020/util"
)

const inputFilePath = "../../inputs/airplane_seats.txt"

// AirplaneSeat represents the seat's position in a plane.
type AirplaneSeat struct {
	row int
	col int
}

func (seat AirplaneSeat) getSeatID() int {
	return seat.row*8 + seat.col
}

// NewAirplaneSeat takes a string of form /[FB]{7}[LR]{3}/ and returns a parsed AirplaneSeat.
func NewAirplaneSeat(s string) (AirplaneSeat, error) {
	s = strings.ReplaceAll(s, "F", "0")
	s = strings.ReplaceAll(s, "B", "1")
	s = strings.ReplaceAll(s, "L", "0")
	s = strings.ReplaceAll(s, "R", "1")
	row, err := strconv.ParseUint(s[:7], 2, 0)
	col, err := strconv.ParseUint(s[7:], 2, 0)
	return AirplaneSeat{int(row), int(col)}, err
}

func toSeatIDs(seatIDs *[]int) func(string) error {
	toIDs := func(s string) error {
		seat, err := NewAirplaneSeat(s)
		if err != nil {
			return err
		}
		*seatIDs = append(*seatIDs, seat.getSeatID())
		return nil
	}
	return toIDs
}

func main() {
	seatIDs := make([]int, 100)
	util.ReadLines(inputFilePath, toSeatIDs(&seatIDs))
	sort.Ints(seatIDs)
	fmt.Println(seatIDs[len(seatIDs)-1]) // part 1.
	for i, id := range seatIDs {
		if i+1 == len(seatIDs) {
			break
		}
		if seatIDs[i+1]-seatIDs[i] == 2 {
			fmt.Println(id + 1) // part 2. Should only print once.
		}
	}
}
