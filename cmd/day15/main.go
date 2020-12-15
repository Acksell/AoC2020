package main

import (
	"fmt"
)

func play(startseq []int, target int) int {
	hist := make(map[int]int)

	for i, n := range startseq { // initialize found.
		hist[n] = i
	}

	previous := 0 // start with 0, assume that the last number in startseq is new.
	for i := len(startseq) + 1; i < target; i++ {
		if lastUsed, found := hist[previous]; found {
			hist[previous] = i - 1
			previous = i - 1 - lastUsed
		} else {
			hist[previous] = i - 1
			previous = 0
		}
	}
	return previous
}

func main() {
	startseq := []int{1, 20, 8, 12, 0, 14}
	fmt.Println(play(startseq, 2020))     // part 1
	fmt.Println(play(startseq, 30000000)) // part 2
}
