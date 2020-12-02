package util

import (
	"bufio"
	"os"
	"strconv"
)

// ReadLines reads a file located at `path` and calls callback on each line.
func ReadLines(path string, callback func(string) error) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		err := callback(txt)
		if err != nil {
			return err
		}
	}
	return nil
}

// ToIntSlice returns a function that takes a string and adds it to slice. Used in ReadLines.
func ToIntSlice(slice *[]uint64) func(string) error {
	storeInt := func(input string) error {
		u64, err := strconv.ParseUint(input, 10, 64)
		if err != nil {
			return err
		}
		*slice = append(*slice, u64)
		return nil
	}
	return storeInt
}
