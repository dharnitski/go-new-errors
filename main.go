package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("notes.txt")
	if err != nil {
		err = fmt.Errorf("Error opening file: %w", err)
		fmt.Printf(" %+v", err)
	}
	defer f.Close()
}
