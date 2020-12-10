package util

import (
	"bufio"
	"os"
)

// A Loadable defines a load function that takes a string and outputs a potential error.
// Used for loading inputs into memory.
type Loadable interface {
	Load(string) error
}

// ReadLines reads a file located at `path` and calls callback on each line.
func ReadLines(path string, l Loadable) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		err := l.Load(txt)
		if err != nil {
			return err
		}
	}
	return nil
}
