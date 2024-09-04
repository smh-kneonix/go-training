package manager

import (
	"fmt"
	"os"

	"example.com/note/note"
)

func WriteFile(note note.Note) {
	data := fmt.Sprint(note)
	os.WriteFile("learn_go.json", []byte(data), 0644)
}

func ReadFile() (string, error) {
	var data, err = os.ReadFile("learn_go.json")
	if err != nil {
		return "", err
	}
	result := string(data)
	return result, nil
}
