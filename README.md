# go-todo-cli

A Go language todo cli built using go built-in libraries to practice JSON serialization and go language best practices 
## Features

* Add, list, complete, and delete tasks directly from your terminal
* Persistent storage via local `todolist.json`
* Stable auto-incrementing task IDs
* Zero external dependencies (built with Go standard library)

## Installation

Ensure you have [Go](https://go.dev/) installed on your machine.

```bash
# Clone the repository
git clone https://github.com/your-username/go-todo-cli.git
cd go-todo-cli

# Build the executable
go build -o todo

```

## Usage
```bash
./todo
```
You will be prompted to enter a number from 0 to 4, where each does the following:
1. View tasks
2. Add new task
3. Change task status
4. Delete task
0. Exit

Adding a new task will automatically set its status to pending. 
Also, the task id is equal to its slice index make the app as simple as possible 
## License

[MIT](https://www.google.com/search?q=LICENSE)
