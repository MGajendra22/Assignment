package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)


// Task 
type task struct {
	id     int
	desc   string
	status bool
}

// Slice of task
type tasks struct {
	cont []task
}

// Pointer to the slice 
type handle struct {
    ptr * tasks
}

// Genereate id number 
var temp_id = idgen()

func idgen() func() int {
	id := 0
	return func() int {
		id++
		return id
	}
}

// Returns Task struct of particular id
func GetTaskById(id int,h *handle) *task {

	for _, val := range h.ptr.cont {

		if val.id == id {
			return &val
		}
	}
	return nil
}

// This request show All Tasks
func (h* handle)handleAllTasks(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET"{

    w.Write([]byte("List of all tasks :- \n"))
	for _, val := range h.ptr.cont {
        str:=fmt.Sprintf("Task %d : %s is done? %v \n",val.id,val.desc,val.status)
		 w.Write([]byte(str))
	}

}
}

// To add Task 
func (h* handle)handleAddTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		bodyBy, err := io.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte("invalid body"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		desc := string(bodyBy)
		if desc == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		} else {
			t := task{temp_id(), desc, false}
			h.ptr.cont = append(h.ptr.cont, t)
            str:=fmt.Sprintf("Task %s added succesfully with id : %d",desc,t.id)
			w.Write([]byte(str))
		}
	}
}

// Get particular task by id
func (h * handle)handleGetTaskById(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		idstr := r.PathValue("id")
		id, err := strconv.Atoi(idstr)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
            return
		}

		t := GetTaskById(id,h)
        str:=fmt.Sprintf("Task of id : %d is %s",id,t.desc)
		w.Write([]byte(str))

	}
}


// To complete particular task
func (h * handle) handleCompleteTask(w http.ResponseWriter, r *http.Request){

    if r.Method=="PUT" {

        idstr := r.PathValue("id")
		id, err := strconv.Atoi(idstr)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
            return
		}
        

        if(h.ptr.cont[id].status) {
          w.Write([]byte("Task is already Completed"))
        }else {
            h.ptr.cont[id].status=true
            str:=fmt.Sprintf("Task %d is completed ",id)
            w.Write([]byte(str))

        }
        
    }
}

// List all the pending tasks
func (h* handle)handlePendingTasks(w http.ResponseWriter, r *http.Request) {
    
    if r.Method =="GET" {
    w.Write([]byte("List of Pending tasks :- \n"))
	for _, val := range h.ptr.cont {
        if(!val.status){
        str:=fmt.Sprintf("Task %d : %s is Not completed yet \n",val.id,val.desc)
		 w.Write([]byte(str))
        } 
	}

 }
}


// Main Function
func main() {

	a1 := task{temp_id(), " Grocery ", false}
	a2 := task{temp_id(), " Shopping ", false}
    
    var h handle

    h.ptr=&tasks{}
    
	h.ptr.cont = append(h.ptr.cont, a1, a2)

	http.HandleFunc("/showall", h.handleAllTasks)
	http.HandleFunc("/task/{id}", h.handleGetTaskById)
	http.HandleFunc("/add", h.handleAddTask)
    http.HandleFunc("/showpending", h.handlePendingTasks)
    http.HandleFunc("/do/{id}", h.handleCompleteTask)


	fmt.Println("Server started ")
	http.ListenAndServe(":8000", nil)

}
