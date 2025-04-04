package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// first create a type of Task by using struct and createa slice for store it and id

// creating type of tasks
type Task struct {
	ID    int
	Title string
	Done  bool
}

// create slice for store tasks  and id
var tasks []Task
var nextId int = 1

// creating function add , list, mark

// AddTask
func addTask(title string) {
	task := Task{
		ID:    nextId,
		Title: title,
		Done:  false,
	}
	nextId++
	tasks = append(tasks, task)
	fmt.Println("✅ Task added successfully!")
}

// List all tasks
func listTasks() {
	// first check there task are not
	if len(tasks) == 0 {
		fmt.Println("📭 No tasks found!")
		return
	}
	for _, task := range tasks {
		status := "❌"
		if task.Done {
			status = "✅"
		}
		fmt.Printf("\n[%d] %s - %s", task.ID, task.Title, status)

	}

}

// Mark a task as completed
func markTaskDone(id int) {
	for i := range tasks {
		if tasks[i].ID == id {
			if tasks[i].Done {
				fmt.Println("⚠️ Task is already completed!")
				return
			}
			tasks[i].Done = true
			fmt.Println("✔️ Task marked as completed!")
			return
		}
	}
	fmt.Println("Task not found.")
}

// Delete a task
func deleteTask(id int) {
	index := -1
	for i, task := range tasks {
		if task.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("❌ Task not found.")
		return
	}
	tasks = append(tasks[:index], tasks[index+1:]...)
	fmt.Println("🗑️ Task deleted successfully!")

}

// make main function
func main() {
	fmt.Println("📋 Welcome to Go Task Manager!")
	reader := bufio.NewReader(os.Stdin)

	// loop
	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1 - Add Task")
		fmt.Println("2 - List Tasks")
		fmt.Println("3 - Mark Task as Done")
		fmt.Println("4 - Delete Task")
		fmt.Println("5 - Exit")

		// get value
		var choice int
		fmt.Print("Enter choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter task title: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)
			if title != "" {
				addTask(title)
			} else {
				fmt.Println("Title cannot be empty")
			}

		case 2:
			listTasks()

		case 3:
			var id int
			fmt.Print("Enter task ID to mark as done: ")
			fmt.Scanln(&id)
			markTaskDone(id)

		case 4:
			var id int
			fmt.Print("Enter task ID to delete it: ")
			fmt.Scanln(&id)
			deleteTask(id)

		case 5:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid option. Please choose 1-4.")
		}

	}
}
