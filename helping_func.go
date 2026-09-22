package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// SelectOperation shows the menu and returns a valid option.
func SelectOperation(reader *bufio.Reader) string {
	for {
		fmt.Println(Magenta + "Select operation:" + Reset)
		fmt.Println("1. Show notes.")
		fmt.Println("2. Add a note.")
		fmt.Println("3. Delete a note.")
		fmt.Println("4. Exit.")

		// Keep asking until the user enters one of the four menu options.
		op, _ := reader.ReadString('\n')
		op = strings.TrimSpace(op)

		if op == "1" || op == "2" || op == "3" || op == "4" {
			return op
		}
		fmt.Println(Red + "Incorrect input. Please choose 1, 2, 3, or 4." + Reset)
		fmt.Println()
	}

}

// CheckArguments checks the command-line arguments and returns the filename.
func CheckArguments(arguments []string) string {
	// The program needs exactly one filename. Help words show the help text.
	if len(arguments) != 2 ||
		arguments[1] == "help" ||
		arguments[1] == "--help" ||
		arguments[1] == "-h" {
		HelpMessage()
		os.Exit(0)
	}
	return arguments[1]
}

// CheckCreateFile checks the collection file and creates it when needed.
func CheckCreateFile(fileName string) {
	// os.Stat tells us whether the file already exists.
	_, err := os.Stat(fileName)

	if os.IsNotExist(err) {
		// Create an empty file on the first run.
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Println(Red+"Error creating file:"+Reset, err)
			os.Exit(1)
		}
		defer file.Close()

		fmt.Println(Green+"Created new notes collection:"+Reset, fileName)
		fmt.Println()
	} else if err == nil {
		// The file exists, so the program can use it.
		fmt.Println(Green+"Notes collection"+Reset, fileName, Green+"exists."+Reset)
		fmt.Println(Green+"Working with notes collection:"+Reset, fileName)
		fmt.Println()
	} else {
		fmt.Println(Red+"Error checking notes collection:"+Reset, err)
		os.Exit(1)
	}
}

// HelpMessage prints simple instructions for starting the program.
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
