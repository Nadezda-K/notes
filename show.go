package main 

import "fmt"

func showNotes(){
	//checking if notes are empty 
	if len(notes) == 0 {
		fmt.Println("No notes found.")
		return 
	}

	//displaying the notes just created 
	for _, noteCollection := range notes {
		//%d is used for inserting numbers, %s used for inserting Text
		fmt.Printf("%d: %s\n", noteCollection.ID, noteCollection.Text)
	}
}