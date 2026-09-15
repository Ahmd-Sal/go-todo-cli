package tasks

import (
	"fmt"
	"slices"
	"todocli/helper"
	"todocli/models"
)

func ViewTasks(todolist []*models.Todo) {

	for index, item := range todolist {
		fmt.Printf("ID: %d | Task: %s | Status: %s\n", index+1, item.Name, helper.ReturnStatus(item.Status))
	}
}
func DeleteTask(todolist []*models.Todo, index int) []*models.Todo {

	i := index - 1
	todolist = slices.Delete(todolist, i, index)

	return todolist
}

func ChangeTaskStatus(todolist []*models.Todo, index int) []*models.Todo {
	ChoosenTask := todolist[index-1]
	switch ChoosenTask.Status {
	case 0:
		ChoosenTask.Status = 1
	case 1:
		ChoosenTask.Status = 0
	}
	return todolist
}

func AppendTask(todolist []*models.Todo, name string, status int) []*models.Todo {

	todolist = append(todolist, &models.Todo{Name: name, Status: status})

	return todolist

}
