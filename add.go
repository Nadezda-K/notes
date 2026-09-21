package main

import (
	"bufio"
	"fmt"

	//	"os"
	"strings"
)

func addNotes(filename string, reader *bufio.Reader) {
	// take text from user
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	if text == "" {
		fmt.Println(Red + "Text field is empty. Please, add text." + Reset)
		fmt.Println()
		return
	}
	notes := LoadNotes(filename)
	notes = append(notes, text)

	//calling this function to save notes to
	SaveNotes(filename, notes)

	fmt.Println(Green + "Note added!" + Reset)
	fmt.Println()
}
