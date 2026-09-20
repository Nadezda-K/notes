package main 

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func addNotes(){
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Print("Enter your note: ")

	text,_ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	//checking all notes plus adding the new one to existing 
	id := len(notes) + 1
	
	//declaring a variable for managing the state 
	var nextID = 1

	//creating a note using the type note from main, storing the values in this 
	note := Note{
	ID: id,
	Text: text,
	}
	
	//state management 
	nextID++
	
	notes = append(notes, note)
	
	//calling this function to save notes to 
	saveNotes()

	fmt.Println("Note added!")
}