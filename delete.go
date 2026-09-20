//The user should be able to enter an ID, for example 2

package main 

import "fmt"

func deleteNote(){
	var id int 

	fmt.Print("Enter note ID to delete: ")
	fmt.Scan(&id)

	//i -- the position of the slice, note - the actual note 
	// Range using notes from main.go 
	for i, note := range notes {
		if note.ID == id{
			notes = append(notes[:i], notes[i+1:]...)
			
			//calling this function to make changes in the json
			saveNotes()

			fmt.Println("Note deleted!")
			return 
		}
	}
	fmt. Println("Note not found")
}

