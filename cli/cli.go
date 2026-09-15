package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"todocli/helper"
	"todocli/models"
	"todocli/store"
	"todocli/tasks"
)

func StartCli(taskList []*models.Todo) {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n=========================")
		fmt.Println("    Todo List Cli     ")
		fmt.Println("=========================")
		fmt.Println("1. View tasks")
		fmt.Println("2. Add new task")
		fmt.Println("3. Change task status")
		fmt.Println("4. Delete task")
		fmt.Println("0. Exit")
		fmt.Print("Choose an option: ")

		if !scanner.Scan() {
			break
		}

		choise := strings.TrimSpace(scanner.Text())

		switch choise {
		case "0":
			return
		case "1":
			tasks.ViewTasks(taskList)
		case "2":
			fmt.Println("Please enter task name:")
			taskList = tasks.AppendTask(taskList, helper.ReadLimitedString(scanner, 30), 0)
			store.SaveTodos(taskList)
			fmt.Println("Task created successfully.")
		case "3":
			tasks.ViewTasks(taskList)
			fmt.Println("Please enter task ID you want to change its status:")
			taskList = tasks.ChangeTaskStatus(taskList, helper.ReadTaskID(scanner, len(taskList)))
			store.SaveTodos(taskList)
			fmt.Println("Changed task status successfully.")
		case "4":
			fmt.Println("Please enter the task ID you want to delete:")
			taskList = tasks.DeleteTask(taskList, helper.ReadTaskID(scanner, len(taskList)))
			store.SaveTodos(taskList)
			fmt.Println("Deleted task successfully")
		default:
			fmt.Println("wrong option")
		}
	}
}
