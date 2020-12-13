package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Acksell/aoc2020/util"
)

const inputFilePath = "../../inputs/buses.txt"

func IsLCM(ints []int, k int) bool {
	isLCM := true
	for _, i := range ints {
		if i%k != 0 {
			isLCM = false
			break
		}
	}
	return isLCM
}

func main() {
	input := make(util.StringSlice, 0)
	util.ReadLines(inputFilePath, &input)
	// part 1
	arrival, _ := strconv.Atoi(input[0])
	busIDs := strings.Split(input[1], ",")
	buses := make([]int, 0)
	busMap := make(map[int]int)
	for i, b := range busIDs {
		if b == "x" {
			continue
		} else {
			period, _ := strconv.Atoi(b)
			buses = append(buses, period)
			busMap[period] = i
		}
	}

	delays := make([]int, 0)
	for _, b := range buses {
		delays = append(delays, b-(arrival%b))
	}
	mindelay, i := util.Min(delays)
	fmt.Println(mindelay * buses[i])
	// part 2, use prime
	time := 1
	n := 1
	for _, b := range buses {
		pos := busMap[b]
		for (time+pos)%b != 0 {
			time += n
		}
		n *= b
	}
	fmt.Println(time)
}
