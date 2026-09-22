package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// addNotes reads one note and saves it in the collection.
func addNotes(filename string, reader *bufio.Reader) {
	// Read until Enter. io.EOF is also allowed because the last line
	// can exist without a new line character.
	text, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Println(Red+"Error reading note:"+Reset, err)
		fmt.Println()
		return
	}

	// Remove spaces and the new line before checking the note.
	text = strings.TrimSpace(text)

	if text == "" {
		fmt.Println(Red + "Text field is empty. Please, add text." + Reset)
		fmt.Println()
		return
	}

	// Load old notes, add the new note to the end, and save all notes again.
	notes := LoadNotes(filename)
	notes = append(notes, text)

	SaveNotes(filename, notes)

	fmt.Println(Green + "Note added!" + Reset)
	fmt.Println()
}
