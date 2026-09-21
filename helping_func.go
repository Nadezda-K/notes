package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// -----------------------------------------------
// Function to display the main menu and get user input.
// -----------------------------------------------
func SelectOperation(reader *bufio.Reader) string {
	for {
		fmt.Println(Magenta + "Select operation:" + Reset)
		fmt.Println("1. Show notes.")
		fmt.Println("2. Add a note.")
		fmt.Println("3. Delete a note.")
		fmt.Println("4. Exit.")

		op, _ := reader.ReadString('\n')
		op = strings.TrimSpace(op)

		if op == "1" || op == "2" || op == "3" || op == "4" {
			return op
		}
		fmt.Println(Red + "Incorrect input. Please choose 1, 2, 3, or 4." + Reset)
		fmt.Println()
	}

}

// -----------------------------------------------
// Function to check arguments, show help if needed,
// and get the file name.
// -----------------------------------------------
func CheckArguments(arguments []string) string {
	if len(arguments) != 2 ||
		arguments[1] == "help" ||
		arguments[1] == "--help" ||
		arguments[1] == "-h" {
		HelpMessage()
		os.Exit(0)
	}
	return arguments[1]
}

// -----------------------------------------------
// Function to check whether the file exists;
// create it if it does not.
// -----------------------------------------------
func CheckCreateFile(fileName string) {
	var _, err = os.Stat(fileName)

	if os.IsNotExist(err) {
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Println(Red+"Error creating file:"+Reset, err)
			os.Exit(1)
		}
		defer file.Close()

		fmt.Println(Green+"Created new notes collection:"+Reset, fileName)
		fmt.Println()
	} else {
		fmt.Println(Green+"Notes collection"+Reset, fileName, Green+"exists."+Reset)
		fmt.Println(Green+"Working with notes collection:"+Reset, fileName)
		fmt.Println()

		// file, err := os.Open(fileName)
		// if err != nil {
		// 	fmt.Println("Error opening file:", err)
		// 	return
		// }
		// // Ensure file is closed
		// defer file.Close()
	}

}

// -----------------------------------------------
//
//	Funcltion to display help message
//
// -----------------------------------------------
func HelpMessage() {
	fmt.Println()
	fmt.Println(Blue + "Usage: ./notestool <file>" + Reset)
	fmt.Println(Blue + "\tOpen an existing note file or create a new one if it does not exist." + Reset)
	fmt.Println()
	fmt.Println(Blue + "Options:" + Reset)
	fmt.Println(Blue + "\t-h, --help, help Display this help message." + Reset)
	fmt.Println()

	fmt.Println(Blue + "Examples:" + Reset)
	fmt.Println(Blue + "./notestool notes.txt" + Reset)
	fmt.Println(Blue + "\tOpen or create notes.txt." + Reset)
	fmt.Println(Blue + "./notestool help" + Reset)
	fmt.Println(Blue + "\tDisplay this help message." + Reset)
	fmt.Println()

	fmt.Println(Blue + "The program requires exactly one argument." + Reset)
	fmt.Println()
	fmt.Println(Blue + "If no argument is provided, more than one argument is provided, or the argument is help, -h, or --help, this help message is displayed." + Reset)
	fmt.Println()

}
