package main

import "fmt"

func showNotes(filename string) {
	notes := LoadNotes(filename)

	//checking if notes are empty
	if len(notes) == 0 {
		fmt.Println(Red + "No notes found." + Reset)
		fmt.Println()
		return
	}

	//displaying the notes just created
	for i, noteCollection := range notes {
		//%d is used for inserting numbers, %s used for inserting Text
		fmt.Printf("%03d - %s\n", i+1, noteCollection)
	}
	fmt.Println()
}
