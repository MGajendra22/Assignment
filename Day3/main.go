package main

import (
	"fmt"
	"os"
)

type task struct {
	id          int
	description string
	status      bool
}

var cont []task

func (t task) String() string {
	return fmt.Sprintf("Task %v : %v | Is Done? : %v", t.id, t.description, t.status)
}

func idgen() func() int {
	id := 0

	return func() int {
		id++
		return id
	}
}

var tempID = idgen()

func addtask(s string) {
	t1 := task{tempID(), s, false}
	cont = append(cont, t1)
}

func ListPendingTask() {
	for _, val := range cont {
		if !val.status {
			fmt.Println(val)
		}
	}
}

func CompleteTask(curid int) {
	for i := range cont {
		if cont[i].id == curid {
			cont[i].status = true

			fmt.Printf("Task %d marked as completed.", curid)
			
			return
		}
	}

	fmt.Println("Invalid Task ID ")
}

func ListAlltasks() {
	for i:= range cont {
		fmt.Println(cont[i])
	}
}

const (
	case0 = 0
	case1 = 1
	case2 = 2
	case3 = 3
	case4 = 4
)

func main() {
	fmt.Println("----- Welcome to Zopdev Task Manager -----")

	for {
		fmt.Println("To add task : 1\nList Pending Task : 2\nTo complete task : 3\nTo show all Tasks : 4\nTo Exit : 0")

		var inp int

		fmt.Print("------- Enter choice: ")

		_, err := fmt.Scanf("%d", &inp)
		if err != nil {
			fmt.Println("Oops! Invalid Choice, Enter again")
			continue
		}

		switch inp {
		case case1:
			var t string

			fmt.Print("Enter Task: ")

			_, err := fmt.Scanf("%s", &t)
			if err != nil {
				fmt.Println("Error")
			}

			addtask(t)
		case case2:
			fmt.Println("\nList of Pending Tasks : - ")
			ListPendingTask()

		case case3:
			var tempID int

			fmt.Print("Enter Task ID to complete: ")

			_,err:=fmt.Scanf("%d\n", &tempID)
			if err != nil {
				fmt.Println("Error")
			}

			CompleteTask(tempID)

		case case4:
			fmt.Println("All tasks :- ")
			ListAlltasks()

		case case0:
			fmt.Println("Exit")
			os.Exit(0)

		default:
			fmt.Println("Oops ! Invalid option. , Try again ")
		}
	}
}
