package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

//-----------------------------------------------
// Function to display the main menu and get user input.
//-----------------------------------------------
func SelectOperation(reader *bufio.Reader) string {
	for {
		fmt.Println("Select operation:")
		fmt.Println("1. Show notes.")
		fmt.Println("2. Add a note.")
		fmt.Println("3. Delete a note.")
		fmt.Println("4. Exit.")

		op, _ := reader.ReadString('\n')
		op = strings.TrimSpace(op)

		if op == "1" || op == "2" || op =="3" || op == "4" { 
			return op
		}
		fmt.Println("Incorrect input. Please choose 1, 2, 3, or 4.")
	}
	
}



//-----------------------------------------------
// Function to check arguments, show help if needed, 
// and get the file name.
//-----------------------------------------------
func CheckArguments(arguments []string) string {
	if ( len(arguments) != 2 || arguments[1] == "help" || 
		 arguments[1] == "--help" || arguments[1] == "-h" ) {
			HelpMessage()
			os.Exit(0)
	}	
	return arguments[1]
}



//-----------------------------------------------
// Function to check whether the file exists;
// create it if it does not.
//-----------------------------------------------
func CheckCreateFile(fileName string) {
	var _, err = os.Stat(fileName)

    if os.IsNotExist(err) {
        file, err := os.Create(fileName)
        if err != nil {
            fmt.Println("Error creating file:",err)
        }
        defer file.Close() // Close file
		fmt.Println("Created new notes collection:", fileName)
		fmt.Println()
    } else {
        fmt.Println("Notes collection", fileName, "exists.")
		fmt.Println("Working with notes collection:", fileName)
		fmt.Println()

		file, err := os.Open(fileName)
    	if err != nil {
        	fmt.Println("Error opening file:", err)
        	return
    	}
    	// Ensure file is closed
    	defer file.Close()
    }
}



//-----------------------------------------------
//	Funcltion to display help message
//-----------------------------------------------
func HelpMessage() {
	fmt.Println()
	fmt.Println("Usage: ./notestool <file>")
	fmt.Println("\tOpen an existing note file or create a new one if it does not exist.")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("\t-h, --help, help Display this help message.")

	fmt.Println("Examples:")
	fmt.Println("./notestool notes.txt")
	fmt.Println("\tOpen or create notes.txt.")

	fmt.Println("./notestool help")
	fmt.Println("\tDisplay this help message.")

	fmt.Println("The program requires exactly one argument.")
	fmt.Println("If no argument is provided, more than one argument is provided,")
	fmt.Println("or the argument is help, -h, or --help, this help message is displayed.")
}




