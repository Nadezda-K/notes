package main

import "fmt"

// showNotes loads and prints all notes in the collection.
func showNotes(filename string) {
	notes := LoadNotes(filename)

	// Stop here when the collection does not contain any notes.
	if len(notes) == 0 {
		fmt.Println(Red + "No notes found." + Reset)
		fmt.Println()
		return
	}

	// Add 1 because users see note numbers starting from 1, not 0.
	for index, note := range notes {
		fmt.Printf("%03d - %s\n", index+1, note)
	}
	fmt.Println()
}
