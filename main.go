package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Todo struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	Deadline    time.Time `json:"deadline"`
}

func addTodo(currentTodos []Todo) []Todo {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Add a todo - Title: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return currentTodos
	}

	input = strings.TrimSpace(input)

	newTodo := Todo{
		Title:       input,
		Description: "coco",
		Completed:   false,
		Deadline:    time.Now(),
	}

	return append(currentTodos, newTodo)
}

func printAllTodo(todos []Todo) {
	fmt.Println()
	for i, todo := range todos {
		status := "[ ]"
		if todo.Completed {
			status = "[✓]"
		}
		formattedDeadline := todo.Deadline.Format("Jan 2, 2006")
		fmt.Printf("%d. %s %s (Due: %s)\n", i+1, status, todo.Title, formattedDeadline)
	}
	fmt.Println()
}

func writeToFile(filename string, todos []Todo) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(todos); err != nil {
		fmt.Println("Error writing JSON:", err)
	}
}

func completeTask(currentTodos []Todo) ([]Todo, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Complete a todo - number: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return currentTodos, nil
	}

	input = strings.TrimSpace(input)
	todoNum, _ := strconv.Atoi(input)
	if todoNum < 1 {
		return nil, errors.New("Enter a valid index of todo")
	}

	for i, _ := range currentTodos {
		if i == todoNum-1 {
			currentTodos[i].Completed = true
		}
	}

	return currentTodos, nil
}

func main() {
	filename := "todos.json"
	var todos []Todo

	// Load existing todos if file exists
	data, err := os.ReadFile(filename)

	if err != nil && !os.IsNotExist(err) {
		fmt.Println("Error reading file:", err)
		panic("cant read from existing storage file, file does not exist")
	}
	if err == nil && len(data) > 0 {
		if err := json.Unmarshal(data, &todos); err != nil {
			fmt.Println("Error decoding JSON:", err)
		}
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Choose an option:")
		fmt.Println("(1) Show Current Todos")
		fmt.Println("(2) Add a Todo")
		fmt.Println("(3) Complete a Todo")
		fmt.Println("(q) Quit")

		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(strings.ToLower(option))

		switch option {
		case "1":
			printAllTodo(todos)
		case "2":
			todos = addTodo(todos)
			writeToFile(filename, todos) // Save immediately
		case "3":
			todos, err := completeTask(todos)
			if err != nil {
				fmt.Println(err)
			} else {
				writeToFile(filename, todos)
				fmt.Println("Todo marked as complete!")
			}

		case "q":
			writeToFile(filename, todos) // Save before quitting
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid command. Please try again.")
		}
	}
}
