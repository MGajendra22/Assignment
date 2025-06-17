package main

import (
	"fmt"
	"os"
)

// Task Struct
type task struct {
	id          int
	description string
	status      bool
}

// Container to contain tasks
var cont []task

// Global ID generator closure
var temp_id func() int = idgen()
func idgen() func() int {
	id := 0
	return func() int {
		id++
		return id
	}
}

// Function to add Task
func addtask(s string) {
	t1 := task{temp_id(), s, false}
	cont = append(cont, t1)
}

// List all pending tasks
func ListTask() {
	for _, val := range cont {
		if val.status==false {
			fmt.Printf("Task %d : %s\n", val.id, val.description)
		}
	}
}

// Mark task as completed
func CompleteTask(curid int) {

	for i := range cont {
		if cont[i].id == curid {
			cont[i].status = true
			fmt.Printf("Task %d marked as completed.\n", curid)
			return
		}
	}
	fmt.Println("Invalid Task ID ")
}

func main() {

	fmt.Println("----- Welcome to Zopdev Task Manager -----")

	for {
		fmt.Println("\nTo add task : 1\n To Show list : 2\n To complete task : 3\n Exit : 0 ")

		var inp int

		 fmt.Print("Enter choice: ")
		    _, err := fmt.Scanf("%d\n", &inp)
		if err != nil {
			fmt.Println("Invalid Choice, Enter again")
			continue
		}

		switch inp {
		case 1:
			var t string
			fmt.Print("Enter Task: ")
			fmt.Scanf("%s\n", &t) 
			addtask(t)

		case 2:
			ListTask()

		case 3:
			var temp_id int
			fmt.Print("Enter Task ID to complete: ")
			fmt.Scanf("%d\n", &temp_id)
			CompleteTask(temp_id)

		case 0:
			fmt.Println("Exit")
			os.Exit(0)

		default:
			fmt.Println("Invalid option. , Try again ")
		}
	}
}
