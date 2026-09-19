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
	//CheckCreateFile(fileName)

	// Open file, create it if it does not.
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
    	fmt.Println(err)
    	return
    }
    defer file.Close()

		// The actual appending is taken care of by the os.O_APPEND flag of the os.OpenFile()
		// function. This flag tells Go to write at the end of the file. Additionally, 
		// the os.O_CREATE flag will make os.OpenFile() create the file if it does not exist, 
		// which is pretty handy. Apart from that, the information is written to the file using 
		// fmt.Fprintf().
    

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
			//fileShow(file)
			file.Seek(0, 0) //Move the file position 0 bytes from the beginning of the file.
			fileReader := bufio.NewReader(file)
			i := 1
			for {
				line, err := fileReader.ReadString('\n')
				if err != nil {
					break
				}
				fmt.Printf("%03d - %v", i, line)
				i++
			}
			fmt.Println()

		case "2": // "add"
			fmt.Println()
			fmt.Println("Enter the note text:")
			//fileAdd()
			note, _ := reader.ReadString('\n')
			fmt.Fprintf(file, note)
		case "3": // "delete"
			fmt.Println()
			fmt.Println("Enter the number of note to remove or 0 to cancel:")
			//fileDelete()
		case "4": // "exit"
			fmt.Println()
			fmt.Println("Exiting....")
			break infForLoop
		}
	}
	
}
