package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// deleteNote removes one note selected by its number.
func deleteNote(filename string, reader *bufio.Reader) {
	// Read the number entered by the user. io.EOF is allowed for the last line.
	idString, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Println(Red+"Error reading note number:"+Reset, err)
		fmt.Println()
		return
	}
	// Remove spaces and the new line before using the input.
	idString = strings.TrimSpace(idString)

	if idString == "0" {
		fmt.Println("Delete operation canceled.")
		fmt.Println()
		return
	}

	notes := LoadNotes(filename)

	// Convert the text number to an integer before using it as a slice index.
	index, err := strconv.Atoi(idString)

	if err != nil {
		fmt.Println(Red+"Error: note number is not a number."+Reset, err)
		fmt.Println()
		return
	}

	if index >= 1 && index <= len(notes) {
		// Users count notes from 1, but slice positions start from 0.
		sliceIndex := index - 1

		// Join the part before the note with the part after the note.
		notes = append(notes[:sliceIndex], notes[sliceIndex+1:]...)

		fmt.Println(Green + "Note deleted!" + Reset)
		fmt.Println()
		// Save the collection only after the selected note was removed.
		SaveNotes(filename, notes)
	} else {
		fmt.Println(Red + "Note not found." + Reset)
		fmt.Println()
	}
}
