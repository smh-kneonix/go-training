package main

import (
	"fmt"

	"example.com/note/input"
	"example.com/note/note"
	"example.com/note/todo"
)

// its a convention that if we have one method in our intreface we should choose name for interface as the name of method +er
type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

type outputable interface {
	saver
	Display()
}

// type outputable interface {
// 	Save() error
// 	Display()
// }

func main() {

	testValue := add("2", "ff")
	fmt.Println(testValue)

	title, content, text := getNoteData()

	newNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	newTodo, err := todo.New(text)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(newNote)
	if err != nil {

		return
	}

	// newTodo.Display()
	// err = saveData(newTodo)
	// if err != nil {

	// 	return
	// }
	err = outputData(newTodo)
	if err != nil {
		return
	}
}

func add[T int | float64 | string](a, b T) T {
	fmt.Printf("a type is: %T, b type is: %T \n", a, b)
	return a + b
}

func outputData(data outputable) error {
	data.Display()
	// we return it cuze maby return error
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the data is failed.")
		return err
	}
	fmt.Println("Saving the data is successful")
	return nil
}

func getNoteData() (string, string, string) {
	title := input.GetInput("Note Title:")

	content := input.GetInput("Note content:")

	todo := input.GetInput("Todo text:")
	return title, content, todo
}
