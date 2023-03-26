package main

import "fmt"
import "your-project/todo"


package todo

type TodoItem struct {
    TaskName string
    Price    float64
    Rating   float64
}



type TodoList struct {
    Items []TodoItem
}


func (list TodoList) FilterByPrice(maxPrice float64) TodoList {
    var filteredItems []TodoItem
    for _, item := range list.Items {
        if item.Price <= maxPrice {
            filteredItems = append(filteredItems, item)
        }
    }
    return TodoList{Items: filteredItems}
}

func (list TodoList) FilterByRating(minRating float64) TodoList {
    var filteredItems []TodoItem
    for _, item := range list.Items {
        if item.Rating >= minRating {
            filteredItems = append(filteredItems, item)
        }
    }
    return TodoList{Items: filteredItems}
}


func main() {
    // Create a new TodoList
    list := todo.TodoList{
        Items: []todo.TodoItem{
            {TaskName: "Buy groceries", Price: 10.99, Rating: 4.5},
            {TaskName: "Clean the house", Price: 0, Rating: 3.2},
            {TaskName: "Finish the project", Price: 50, Rating: 4.8},
        },
    }

    // Filter the list by price
    expensiveItems := list.FilterByPrice(30.0)
    fmt.Println("Expensive items:")
    for _, item := range expensiveItems.Items {
        fmt.Printf("- %s ($%.2f)\n", item.TaskName, item.Price)
    }

    // Filter the list by rating
    highRatedItems := list.FilterByRating(4.0)
    fmt.Println("High-rated items:")
    for _, item := range highRatedItems.Items {
        fmt.Printf("- %s (%.1f stars)\n", item.TaskName, item.Rating)
    }
}
