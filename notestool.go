package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
		switch op {
		case "1": // "show"
			fmt.Println()
			ShowNotes(notes) // Show all notes
		case "2": // "add"
			fmt.Println("Enter the note text:")
			text, _ := reader.ReadString('\n')
			notes = AddNote(notes, text)
			SaveNotes(fileName, notes)
		case "3": // "delete"
			fmt.Println("Enter the note number to delete:")
			numStr, _ := reader.ReadString('\n')
			cleanNum := strings.TrimSpace(numStr)
			num, err := strconv.Atoi(cleanNum)
			if err != nil {
				fmt.Println("Invalid note number")
			} else {
				notes = DeleteNote(notes, num)
				SaveNotes(fileName, notes)
			}
		case "4": // "exit"
			fmt.Println("Exiting....")
			break infForLoop
		}
	}

}
