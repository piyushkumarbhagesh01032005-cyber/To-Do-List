package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var tasks []string

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n ------------ TO-DO-List ------------")
		fmt.Println("1. Add a tasks")
		fmt.Println("2. Show all tasks")
		fmt.Println("3. Mark task as done")
		fmt.Println("4. Delete a task")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice:")

		scanner.Scan()
		Choice := scanner.Text()

		switch Choice {
		case "1":
			addTask(scanner)
		case "2":
			showTask()
		case "3":
			markDone(scanner)
		case "4":
			deleteTask(scanner)
		case "5":
			fmt.Println("--Goodbye!--")
			return
		default:
			fmt.Println("Invalid choice! Please try again.")
		}
	}

}

func addTask(scanner *bufio.Scanner) {
	fmt.Print("Enter task: ")
	scanner.Scan()
	task := (scanner.Text())
	/*if task == "" {
		fmt.Println("Task cannot be empty:")
		return
	}*/
	tasks = append(tasks, task)
	fmt.Println("Task added successfully!")
}

func showTask() {
	/*if len(tasks) == 0 {
		fmt.Println(" No Task:")
		return
	}*/
	fmt.Println("\nYour Tasks:")
	for index, value := range tasks {
		fmt.Printf("%d %s\n", index+1, value)
	}
}

func markDone(scanner *bufio.Scanner) {
	/*if len(tasks) == 0 {
		fmt.Println("No tasks to mark as done!")
		return
	}*/
	showTask()
	fmt.Print("Enter task number to mark as done: ")
	scanner.Scan()

	num, _ := strconv.Atoi(scanner.Text())
	tasks[num-1] = tasks[num-1] + ""
	fmt.Println("Task marked as done!")
}

func deleteTask(scanner *bufio.Scanner) {
	/*	if len(tasks) == 0 {
		fmt.Println("No Task number to delete:")
		return
	}*/
	showTask()
	fmt.Print("Enter the number to delete:")
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	tasks = append(tasks[:num-1], tasks[num:]...)
	fmt.Println("Tasks deleted successfully!")
}
