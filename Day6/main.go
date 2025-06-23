package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"log"
)

type task struct {
	id     int
	desc   string
	status bool
}

type tasks struct {
	cont []task
}

type handle struct {
	ptr *tasks
}

// var t1 tasks

var temp_id = idgen()

func idgen() func() int {
	id := 0
	return func() int {
		id++
		return id
	}
}

func GetTaskById(id int, h *handle) *task {

	for _, val := range h.ptr.cont {

		if val.id == id {
			return &val
		}
	}
	return nil
}

func (h *handle) handleAllTasks(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		_, err := w.Write([]byte("List of all tasks :- \n"))
		if err != nil {
			fmt.Printf("Error")
		}
		for _, val := range h.ptr.cont {
			str := fmt.Sprintf("Task %d : %s is done? %v \n", val.id, val.desc, val.status)
			_, err := w.Write([]byte(str))
			if err != nil {
				fmt.Printf("Error")
			}
		}
	}
}

func (h *handle) handleAddTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		bodyBy, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, err := w.Write([]byte("invalid body"))
			if err != nil {
				fmt.Printf("Error")
			}
			return
		}

		desc := string(bodyBy)
		if desc == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		} else {
			t := task{temp_id(), desc, false}
			h.ptr.cont = append(h.ptr.cont, t)
			str := fmt.Sprintf("Task %s added succesfully with id : %d", desc, t.id)
			w.WriteHeader(http.StatusCreated)
			_, err := w.Write([]byte(str))
			if err != nil {
				fmt.Printf("Error")
			}
		}
	}
}

func (h *handle) handleGetTaskById(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		idstr := r.PathValue("id")
		id, err := strconv.Atoi(idstr)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, err := w.Write([]byte(`invalid id`))
			if err != nil {
				fmt.Printf("Error")
			}
			return
		}

		t := GetTaskById(id, h)
		str := fmt.Sprintf("Task of id : %d is %s", id, t.desc)
		_, err = w.Write([]byte(str))
		if err != nil {
			fmt.Printf("Error")
		}
	}
}

func (h *handle) handleCompleteTask(w http.ResponseWriter, r *http.Request) {

	if r.Method == "PUT" {

		idstr := r.PathValue("id")
		id, err := strconv.Atoi(idstr)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if h.ptr.cont[id].status {
			_, err := w.Write([]byte("Task is already Completed"))
			if err != nil {
				fmt.Printf("Error")
			}
		} else {
			h.ptr.cont[id].status = true
			str := fmt.Sprintf("Task %d is completed ", id)
			_, err := w.Write([]byte(str))
			if err != nil {
				fmt.Printf("Error")
			}

		}

	}
}

func (h *handle) handlePendingTasks(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("List of Pending tasks :- \n"))
		if err != nil {
			fmt.Printf("Error")
		}
		for _, val := range h.ptr.cont {
			if !val.status {
				str := fmt.Sprintf("Task %d : %s is Not completed yet \n", val.id, val.desc)
				_, err := w.Write([]byte(str))
				if err != nil {
					fmt.Printf("Error")
				}
			}
		}

	}
}

func main() {

	a1 := task{temp_id(), " Grocery ", false}
	a2 := task{temp_id(), " Shopping ", false}

	var h handle

	h.ptr = &tasks{}

	h.ptr.cont = append(h.ptr.cont, a1, a2)

	http.HandleFunc("/showall", h.handleAllTasks)
	http.HandleFunc("/task/{id}", h.handleGetTaskById)
	http.HandleFunc("/add", h.handleAddTask)
	http.HandleFunc("/showpending", h.handlePendingTasks)
	http.HandleFunc("/do/{id}", h.handleCompleteTask)

	fmt.Println("Server Started running on http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", nil))

}
