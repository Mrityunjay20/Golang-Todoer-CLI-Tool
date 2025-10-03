package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// Todo is a struct that represents a single todo item.
// We've expanded it to be more descriptive.

type Todo struct {
	Title       string
	Description string
	Completed   bool
	Deadline    time.Time
}

func addTodo(CurrentTodos []Todo) []Todo {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Add a todo - Title\n")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
	}

	// Trim the newline and any extra spaces
	input = strings.TrimSpace(input)

	newTodos := Todo{
		Title:       input,
		Description: "coco",
		Completed:   false,
		Deadline:    time.Now(),
	}

	return append(CurrentTodos, newTodos)
}

func printAllTodo(todos []Todo) {
	for i, todo := range todos {
		status := "[ ]"
		if todo.Completed {
			status = "[✓]"
		}
		formattedDeadline := todo.Deadline.Format("Jan 2, 2006")

		fmt.Printf("%d. %s %s (Due: %s)\n", i+1, status, todo.Title, formattedDeadline)
		fmt.Println("")
	}
}

func main() {
	filename := "todos.json"
	var todos []Todo

	// Try to open the file, create it if it doesn't exist
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error opening/creating file:", err)
		return
	}
	defer file.Close()

	fmt.Println("File is ready:", filename)

	writer := bufio.NewWriterSize(file, 1024*1024)
	writer.WriteString("Line 1\n")
	writer.WriteString("Line n2\n")
	writer.Flush()

	for {
		fmt.Println("Choose an option:")
		fmt.Println("(1) Show Current Todos")
		fmt.Println("(2) Add a Todo")
		fmt.Println("(q) Quit")

		reader := bufio.NewReader(os.Stdin)
		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(strings.ToLower(option))

		switch option {
		case "1":
			printAllTodo(todos)
		case "2":
			todos = addTodo(todos)
		case "q":
			return
		case "Q":
			return
		default:
			fmt.Println("Invalid command. Please try again.")
		}
	}

}
