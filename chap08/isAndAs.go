package main

import (
	"errors"
	"fmt"
	"os"
)

func fileChecker(name string) error {
	f, err := os.Open(name)
	defer f.Close()
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err)
	}
	return nil
}

func main() {
	err := fileChecker("not_here.txt")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("That file does'nt exist....")
		}
	}
}
