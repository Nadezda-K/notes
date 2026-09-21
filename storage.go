package main

import (
	"bufio"
	"fmt"
	"os"
)

// LoadNotes reads all notes from a file.
func LoadNotes(filename string) []string {
	// Open the file in read-only mode.
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(Red+"Error opening file:"+Reset, err)
		return nil
	}

	defer file.Close()

	var notes []string

	// Scanner reads the file one line at a time.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		notes = append(notes, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(Red+"Error reading file:"+Reset, err)
	}

	return notes
}

// SaveNotes writes all notes to a file.
func SaveNotes(filename string, notes []string) {
	// os.Create makes a file or clears the old file before writing.
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(Red+"Error creating file:"+Reset, err)
		return
	}
	defer file.Close()

	// Each note is stored on its own line.
	for _, note := range notes {
		fmt.Fprintln(file, note)
	}
}
