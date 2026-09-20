package main

import (
	"bufio"
	"fmt"
	"os"
	// "strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	//	Taking arguments form command line
	arguments := os.Args

	//	Checking number of arguments, showing help menue if necessary,
	//	taking name of file
	fileName := CheckArguments(arguments)
	fmt.Printf("File name: %v, %T", fileName, fileName)

	// Check if file exist - open, if not - create

	fmt.Println()
	fmt.Println("Welcome to the Notes Tool!")
	fmt.Println()

	notes := LoadNotes(fileName)
infForLoop:
	for {
		op := SelectOperation(reader)
		//TODO: add a new note to the notes slice
		switch op {
		case "0":
		case "1": // "show"
			fmt.Println()
			ShowNotes(notes) // Show all notes
		case "2": // "add"
			fmt.Println()
			AddNote(notes, text) // Add a new note)
		case "3": // "delete"
			fmt.Println()
			DeleteNote(notes)
		case "4": // "exit"
			fmt.Println()
			fmt.Println("Exiting....")
			break infForLoop
		}
	}

}
