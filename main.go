package main

import (
	"todocli/cli"
	"todocli/store"
)

func main() {

	jsonFile := store.LoadTodos("todolist.json")

	cli.StartCli(jsonFile)

}
