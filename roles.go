package main

import (
	"errors"
	"fmt"
)

// Define the possible roles
const (
	RoleAdmin  = "admin"
	RoleClient = "client"
	RoleSeller = "seller"
)

// Define the ToDo struct
type ToDo struct {
	Title    string
	Complete bool
}

// Define the User struct
type User struct {
	Name     string
	Email    string
	Password string
	Role     string
}

// Define the ToDoList struct
type ToDoList struct {
	ToDos []ToDo
}

// Define the AddToDo method on the ToDoList struct
func (tl *ToDoList) AddToDo(todo ToDo) {
	tl.ToDos = append(tl.ToDos, todo)
}

// Define the GetToDos method on the ToDoList struct
func (tl *ToDoList) GetToDos() []ToDo {
	return tl.ToDos
}

// Define the RemoveToDo method on the ToDoList struct
func (tl *ToDoList) RemoveToDo(index int) error {
	if index < 0 || index >= len(tl.ToDos) {
		return errors.New("invalid index")
	}
	tl.ToDos = append(tl.ToDos[:index], tl.ToDos[index+1:]...)
	return nil
}

// Define the main function
func main() {
	// Create some sample ToDos
	todo1 := ToDo{Title: "Buy milk", Complete: false}
	todo2 := ToDo{Title: "Clean the house", Complete: true}
	todo3 := ToDo{Title: "Finish project", Complete: false}

	// Create some sample users
	admin := User{Name: "Admin", Email: "admin@example.com", Password: "admin", Role: RoleAdmin}
	client := User{Name: "Client", Email: "client@example.com", Password: "client", Role: RoleClient}
	seller := User{Name: "Seller", Email: "seller@example.com", Password: "seller", Role: RoleSeller}

	// Create a ToDoList and add the ToDos
	todoList := ToDoList{}
	todoList.AddToDo(todo1)
	todoList.AddToDo(todo2)
	todoList.AddToDo(todo3)

	// Display the ToDos for the client role
	if client.Role == RoleClient {
		fmt.Println("ToDos:")
		for _, todo := range todoList.GetToDos() {
			fmt.Printf("%s - %t\n", todo.Title, todo.Complete)
		}
	}

	// Remove a ToDo for the admin role
	if admin.Role == RoleAdmin {
		err := todoList.RemoveToDo(1)
		if err != nil {
			fmt.Println("Error removing ToDo:", err)
		}
	}

	// Display the ToDos for the seller role
	if seller.Role == RoleSeller {
		fmt.Println("ToDos:")
		for _, todo := range todoList.GetToDos() {
			fmt.Printf("%s - %t\n", todo.Title, todo.Complete)
		}
	}
}
