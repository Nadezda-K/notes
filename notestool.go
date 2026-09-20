package main

import (
	"fmt"
	"bufio"
	"os"
//	"strings"
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
	infForLoop: for {
		// Display the main menu and get user input.
		op := SelectOperation(reader)

		switch op {
		case "1": // "show"
			fmt.Println()
			fmt.Println("Notes:")
			
			showNotes(fileName)

		case "2": // "add"
			fmt.Println()
			fmt.Println("Enter the note text:")
			
			addNotes(fileName, reader)
		case "3": // "delete"
			fmt.Println()
			fmt.Println("Enter the number of note to remove or 0 to cancel:")

			deleteNote(fileName, reader)
		case "4": // "exit"
			fmt.Println()
			fmt.Println("Exiting....")
			break infForLoop
		}
	}
	
}
