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

infForLoop:
	for {
		op := SelectOperation(reader)

		switch op {
		case "1": // "show"
			fmt.Println()
			fmt.Println("show")
		case "2": // "add"
			fmt.Println()
			fmt.Println("add")
		case "3": // "delete"
			fmt.Println()
			fmt.Println("delete")
		case "4": // "exit"
			fmt.Println()
			fmt.Println("Exiting....")
			break infForLoop
		}
	}

}
