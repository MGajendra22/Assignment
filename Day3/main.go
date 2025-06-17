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

//stringer

func (t task) String() string {

	return fmt.Sprintf("Task %v : %v | Is Done? : %v",t.id,t.description,t.status)
}



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
func ListPendingTask() {
	for _, val := range cont {
		if val.status==false {
			fmt.Println(val)
		}
	}

	fmt.Println("\n")
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

func ListAlltasks(){

	for i,_:=range cont{
        fmt.Println(cont[i])
	}

	fmt.Println("\n")
}

func main() {

	fmt.Println("----- Welcome to Zopdev Task Manager -----")

	for {
		fmt.Println("To add task : 1\nList Pending Task : 2\nTo complete task : 3\nTo show all Tasks : 4\nTo Exit : 0\n")

		var inp int

		 fmt.Print("------- Enter choice: ")
		    _, err := fmt.Scanf("%d\n", &inp)
		if err != nil {
			fmt.Println("Oops! Invalid Choice, Enter again\n")
			continue
		}

		switch inp {
		case 1:
			var t string
			fmt.Print("Enter Task: ")
			fmt.Scanf("%s\n", &t) 
			addtask(t)

		case 2:
			fmt.Println("\nList of Pending Tasks : - ")
			ListPendingTask()

		case 3:
			var temp_id int
			fmt.Print("Enter Task ID to complete: ")
			fmt.Scanf("%d\n", &temp_id)
			CompleteTask(temp_id)

		case 4:
			fmt.Println("All tasks :- \n")
			ListAlltasks()	

		case 0:
			fmt.Println("Exit")
			os.Exit(0)

		default:
			fmt.Println("Oops ! Invalid option. , Try again \n")
		}
	}
}
