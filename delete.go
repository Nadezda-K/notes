//The user should be able to enter an ID, for example 2

package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func deleteNote(filename string, reader *bufio.Reader) {
	id_str, _ := reader.ReadString('\n')
	id_str = strings.TrimSpace(id_str)

	if id_str == "0" {
		fmt.Println("Cancel delete operation")
		fmt.Println()
		return
	}

	notes := LoadNotes(filename)

	//i -- the position of the slice, note - the actual note
	// Range using notes from main.go
	inx, err := strconv.Atoi(id_str)

	if err != nil {
		fmt.Println("Error: index is not a number", err)
		fmt.Println()
		return
	}

	if inx <= len(notes) {
		for i, _ := range notes {
			if i+1 == inx {
				notes = append(notes[:i], notes[i+1:]...)

				fmt.Println("Note deleted!")
				fmt.Println()
			}
			//calling this function to make changes in the json
			SaveNotes(filename, notes)
		}
	} else {
		fmt.Println("Note not found")
		fmt.Println()
	}
	return
}
