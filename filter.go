package main

import (
	"fmt"
)

type Task struct {
	Description string
	Price       float64
	Rating      int
	Status      string
	Priority    string
}

func main() {
	// Create a slice of tasks
	tasks := []Task{
		{"Research and select e-commerce platform", 0, 3, "Not Started", "Medium"},
		{"Hire web developer and designer", 5000, 4, "Not Started", "High"},
		{"Create a product catalog and set up inventory management", 0, 5, "Not Started", "High"},
		{"Develop website design and user experience", 8000, 4, "Not Started", "High"},
		{"Implement payment gateway and shipping options", 2000, 3, "Not Started", "High"},
		{"Integrate social media marketing and SEO strategies", 3000, 5, "Not Started", "Medium"},
		{"Launch website and test functionality", 0, 4, "Not Started", "High"},
		{"Monitor website analytics and optimize for performance", 5000, 3, "Not Started", "Medium"},
	}

	// Define filter function
	filterTasks := func(minPrice, maxPrice float64, minRating, maxRating int, tasks []Task) []Task {
		filteredTasks := []Task{}
		for _, task := range tasks {
			if task.Price >= minPrice && task.Price <= maxPrice && task.Rating >= minRating && task.Rating <= maxRating {
				filteredTasks = append(filteredTasks, task)
			}
		}
		return filteredTasks
	}

	// Filter tasks by price and rating
	filteredTasks := filterTasks(0, 5000, 4, 5, tasks)

	// Print filtered tasks
	for _, task := range filteredTasks {
		fmt.Printf("%s ($%.2f, %d stars) - %s, %s\n", task.Description, task.Price, task.Rating, task.Status, task.Priority)
	}
}
