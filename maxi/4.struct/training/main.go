package main

import (
	"fmt"

	"example.com/note/input"
	"example.com/note/note"
)

func main() {
	title, content := getNoteData()

	var newNote, err = note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	newNote.Display()
	err = newNote.Save()
	if err != nil {
		fmt.Println("Saving the note is failed.")
		return
	}
	fmt.Println("Saving the note is successful~")

	// manager.WriteFile(*newNote)

	// newNote.GetNotes()

	// note, err := manager.ReadFile()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(note)

}

func getNoteData() (string, string) {
	title := input.GetInput("Note Title:")

	content := input.GetInput("Note content:")
	return title, content
}
