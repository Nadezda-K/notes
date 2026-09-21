package main

import (
	"bufio"
	"fmt"
	"os"
	// "strings"
)

const (
	Reset = "\033[0m"
	Red = "\033[31m"
	Green = "\033[32m"
	Blue = "\033[34m"
	Magenta = "\033[35m"
)

func main() {
	fmt.Println()
	fmt.Println("Welcome to the Notes Tool!")
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)

	//	Taking arguments form command line
	arguments := os.Args

	// Check arguments, show help if needed, and get the file name.
	fileName := CheckArguments(arguments)

	// Check whether the file exists; create it if it does not.
	CheckCreateFile(fileName)

	//-----------------------------------------------
	//	Main menu of the tool
	//-----------------------------------------------
infForLoop:
	for {
		// Display the main menu and get user input.
		op := SelectOperation(reader)

		switch op {
		case "1": // "show"
			fmt.Println()
			fmt.Println(Magenta + "Notes:" + Reset)

			showNotes(fileName)

		case "2": // "add"
			fmt.Println()
			fmt.Println(Magenta + "Enter the note text:" + Reset)

			addNotes(fileName, reader)
		case "3": // "delete"
			fmt.Println()
			fmt.Println(Magenta + "Enter the number of note to remove or 0 to cancel:" + Reset)

			deleteNote(fileName, reader)
		case "4": // "exit"
			fmt.Println()
			fmt.Println("Exiting....")
			break infForLoop
		}
	}

}
