package main

import (
	"fmt"
	// "os"
)

type task struct {
	id          int
	description string
	status      bool
}

type tasks struct {
	cont []task
	id   int
}

type handle struct {
	ptr *tasks
}


func addtask(h *handle, s string) {
	h.ptr.id++
	t1 := task{h.ptr.id, s, false}
	h.ptr.cont = append(h.ptr.cont, t1)
}


func CompleteTask(h *handle, curid int) {
	for i := range h.ptr.cont {
		if h.ptr.cont[i].id == curid {
			h.ptr.cont[i].status = true

			fmt.Printf("Task %d marked as completed.", curid)

			return
		}
	}

	fmt.Println("Invalid Task ID ")
}

// func ListAlltasks(h *handle) {
// 	for i := range h.ptr.cont {
// 		fmt.Println(h.ptr.cont[i])
// 	}
// }

