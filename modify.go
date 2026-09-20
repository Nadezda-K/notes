package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func modifyFile() {

	// Get all files in the current directory
	files, err := os.ReadDir(".")

	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	fmt.Println()
	fmt.Println("Files:")
	fmt.Println("------")

	fileNumber := 1

	// Display files
	for _, file := range files {

		if !file.IsDir() {
			fmt.Printf("%d. %s\n", fileNumber, file.Name())
			fileNumber++
		}
	}

	// Ask the user which file they want
	var choice int

	fmt.Print("\nChoose a file: ")
	fmt.Scan(&choice)

	// Find the selected file
	fileNumber = 1

	for _, file := range files {

		if !file.IsDir() {

			if fileNumber == choice {

				fmt.Println()
				fmt.Println("You selected:", file.Name())

				// Read the existing file
				content, err := os.ReadFile(file.Name())

				if err != nil {
					fmt.Println("Error reading file:", err)
					return
				}

				fmt.Println()
				fmt.Println("Current contents:")
				fmt.Println("-----------------")
				fmt.Println(string(content))

				// Get new content from user
				reader := bufio.NewReader(os.Stdin)

				fmt.Println()
				fmt.Println("Enter the new content:")
				fmt.Println("(Type your content and press Enter)")
				
				// declares a variable to get the input from user 
				newContent, err := reader.ReadString('\n')

				if err != nil {
					fmt.Println("Error reading new content:", err)
					return
				}

				newContent = strings.TrimSpace(newContent)

				// Main logic of saving the new content to the file
				// 1st arg - name of the file selected by user 
				// os.WriteFile -- expects the data as bytes 
				// 0644 - Create/write a file with normal readable permissions 
				err = os.WriteFile(
					file.Name(), 
					[]byte(newContent+"\n"),
					0644,
				)

				if err != nil {
					fmt.Println("Error saving file:", err)
					return
				}

				fmt.Println("File saved successfully!")

				return
			}

			fileNumber++
		}
	}

	fmt.Println("Invalid file selection.")
}