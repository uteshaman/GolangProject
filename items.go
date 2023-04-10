package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TodoItem struct {
	text   string
	done   bool
	rating int
}

func main() {
	todoList := make([]TodoItem, 0)

	for {
		fmt.Print("Enter a command (add, complete, rate, list, exit): ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "add":
			fmt.Print("Enter task: ")
			text, _ := reader.ReadString('\n')
			text = strings.TrimSpace(text)
			item := TodoItem{text, false, 0}
			todoList = append(todoList, item)
			fmt.Println("Task added!")
		case "complete":
			fmt.Print("Enter task number: ")
			indexStr, _ := reader.ReadString('\n')
			indexStr = strings.TrimSpace(indexStr)
			index, _ := strconv.Atoi(indexStr)
			if index < 1 || index > len(todoList) {
				fmt.Println("Invalid task number")
				continue
			}
			todoList[index-1].done = true
			fmt.Println("Task completed!")
		case "rate":
			fmt.Print("Enter task number: ")
			indexStr, _ := reader.ReadString('\n')
			indexStr = strings.TrimSpace(indexStr)
			index, _ := strconv.Atoi(indexStr)
			if index < 1 || index > len(todoList) {
				fmt.Println("Invalid task number")
				continue
			}
			fmt.Print("Enter rating (1-5): ")
			ratingStr, _ := reader.ReadString('\n')
			ratingStr = strings.TrimSpace(ratingStr)
			rating, _ := strconv.Atoi(ratingStr)
			if rating < 1 || rating > 5 {
				fmt.Println("Invalid rating")
				continue
			}
			todoList[index-1].rating = rating
			fmt.Println("Rating added!")
		case "list":
			if len(todoList) == 0 {
				fmt.Println("No tasks in list")
			} else {
				for i, item := range todoList {
					status := " "
					if item.done {
						status = "X"
					}
					fmt.Printf("%d. [%s] %s (rating: %d)\n", i+1, status, item.text, item.rating)
				}
			}
		case "exit":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}
