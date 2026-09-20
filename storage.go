package main

import (
	"bufio"
	"fmt"
	"os"
)

// LoadNotes reads notes from the specified file line by line and returns them as a slice.
func LoadNotes(filename string) []string {
	// Open the file in read-only mode
	file, err := os.Open(filename)
	if err != nil {
		// If the file does not exist yet (first launch), return nil as an empty list
		return nil
	}
	// Guarantee that the operating system closes the file before the function returns
	defer file.Close()

	var notes []string

	// Create a scanner to read the file line by line without loading it all into memory
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Extract the current line text and append it to our notes slice
		notes = append(notes, scanner.Text())
	}

	// Check if scanning stopped because of a disk read error rather than reaching EOF
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return notes
}

// SaveNotes writes the provided notes slice into the file, overwriting any previous content.
func SaveNotes(filename string, notes []string) {
	// os.Create creates the file if missing, or truncates (clears) it if it already exists
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	// Ensure the file descriptor is released when saving is finished
	defer file.Close()

	// TODO: iterate over the notes slice and write each note on a new line
}
