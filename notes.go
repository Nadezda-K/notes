package main

import "fmt"

// ShowNotes prints all notes to the console with numbered entries.
func ShowNotes(notes []string) {
	fmt.Println("Notes:") // Print a header for the notes
	for i, note := range notes {
		fmt.Printf("%03d - %s\n", i+1, note)
	}
	// "%03d - %s\n" formats the index as a three-digit number with leading zeros, followed by the note text.
}

// AddNote appends a new note to the existing slice and returns the updated list.
func AddNote(notes []string, text string) []string {
	return append(notes, text) // Append the new note to the slice and return the updated slice
}

// DeleteNote removes a note by its 1-based position in the list.
// If the number is zero or out of range, the function returns the original slice unchanged.
func DeleteNote(notes []string, num int) []string {
	if num == 0 { // Check if input is zero
		fmt.Println("Invalid note number (can't be zero) and must be an integer type!")
		return notes
	} else if num < 1 || num > len(notes) { // Check if input is less than 1 or greater than the length of notes
		fmt.Println("Invalid note number (can't be less than 1) and must be an integer type!")
		return notes
	} else { // Valid input, proceed to delete the note
		index := num - 1
		// Remove the note at the specified index
		notes = append(notes[:index], notes[index+1:]...)
		return notes
	}

}
