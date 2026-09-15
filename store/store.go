package store

import (
	"encoding/json"
	"fmt"
	"os"
	"todocli/models"
)

func SaveTodos(todolist []*models.Todo) {

	fileData, err := json.Marshal(todolist)
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}

	err = os.WriteFile("todolist.json", fileData, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

}

func LoadTodos(fileName string) []*models.Todo {
	todolist := []*models.Todo{}

	fileData, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error reading file. A new json file will be created instead")
		return todolist
	}

	if err := json.Unmarshal(fileData, &todolist); err != nil {
		fmt.Println("Error Unmarshaling json. A new json file will be created instead")
		return todolist
	}
	fmt.Println("File loaded successfully.")
	return todolist

}
