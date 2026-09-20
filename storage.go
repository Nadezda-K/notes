package main

import (
	"bufio"
	"fmt"
	"os"
)

// LoadNotes reads notes from the specified file line by line.
func LoadNotes(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		// No file = no notes, so we can return an empty slice
		return nil
	}
	defer file.Close()

	var notes []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		notes = append(notes, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return notes
}

func SaveNotes(filename string, notes []string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	//TODO: add a check if the file broken, if yes - create a new one, if not - rewrite the existing one
	//TODO: at least add the finish func SaveNotes
}
