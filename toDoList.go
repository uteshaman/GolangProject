package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Todo struct {
	title       string
	description string
	completed   bool
}

type TodoList struct {
	todos []*Todo
}

func (t *Todo) String() string {
	status := " "
	if t.completed {
		status = "X"
	}
	return fmt.Sprintf("[%s] %s: %s", status, t.title, t.description)
}

func (tl *TodoList) add(todo *Todo) {
	tl.todos = append(tl.todos, todo)
}

func (tl *TodoList) list() {
	for _, todo := range tl.todos {
		fmt.Println(todo)
	}
}

func (tl *TodoList) complete(index int) {
	if index < 0 || index >= len(tl.todos) {
		fmt.Println("Invalid todo index")
		return
	}
	tl.todos[index].completed = true
	fmt.Println("Todo completed:", tl.todos[index])
}

func main() {
	todoList := &TodoList{}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Please choose an option:")
		fmt.Println("1. Add todo")
		fmt.Println("2. List todos")
		fmt.Println("3. Complete todo")
		fmt.Println("4. Exit")
		scanner.Scan()
		option, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil {
			fmt.Println("Invalid option")
			continue
		}
		switch option {
		case 1:
			fmt.Println("Enter title:")
			scanner.Scan()
			title := scanner.Text()
			fmt.Println("Enter description:")
			scanner.Scan()
			description := scanner.Text()
			todo := &Todo{title: title, description: description}
			todoList.add(todo)
			fmt.Println("Todo added")
		case 2:
			fmt.Println("Todo List:")
			todoList.list()
		case 3:
			fmt.Println("Enter the index of the todo to complete:")
			scanner.Scan()
			index, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("Invalid todo index")
				continue
			}
			todoList.complete(index)
		case 4:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid option")
		}
	}
}
