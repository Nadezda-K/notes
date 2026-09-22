package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
)

func main() {
	// Get the collection filename before starting the program.
	fileName := CheckArguments(os.Args)

	fmt.Println()
	fmt.Println("Welcome to the Notes Tool!")
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)

	// Make sure the collection file is ready before using it.
	CheckCreateFile(fileName)

infForLoop:
	for {
		// The loop lets the user perform several actions in one run.
		op := SelectOperation(reader)

		switch op {
		case "1":
			// Show all notes stored in the selected collection.
			fmt.Println()
			fmt.Println(Magenta + "Notes:" + Reset)

			showNotes(fileName)

		case "2":
			// Read and save one new note.
			fmt.Println()
			fmt.Println(Magenta + "Enter the note text:" + Reset)

			addNotes(fileName, reader)
		case "3":
			// Ask for a note number and remove that note.
			fmt.Println()
			fmt.Println(Magenta + "Enter the number of note to remove or 0 to cancel:" + Reset)

			deleteNote(fileName, reader)
		case "4":
			// Leave the loop and finish the program.
			fmt.Println()
			fmt.Println("Exiting....")
			break infForLoop
		}
	}

}
