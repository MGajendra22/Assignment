package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type task struct {
	ID     int    `json:"id"`
	Desc   string `json:"desc"`
	Status bool   `json:"status"`
}

type handle struct {
	tasks []task
}

var idGen = func() func() int {
	id := 0
	return func() int {
		id++
		return id
	}
}()

// Handler : To add Task
func (h *handle) handleAddTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var t task
	err = json.Unmarshal(body, &t)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	t.ID = idGen()
	h.tasks = append(h.tasks, t)

	resp, _ := json.Marshal(t)

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	_, err1 := w.Write(resp)

	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}
}

// Handler : To Get Task by Id
func (h *handle) handleGetTaskByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for _, t := range h.tasks {
		if t.ID == id {
			resp, err := json.Marshal(t)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			w.Header().Set("Content-Type", "application/json")

			w.WriteHeader(http.StatusOK)
			_, err = w.Write(resp)
			if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return 
		}
			return }
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}

// Handler : To update (Complete the particular task) [Mark status : true]
func (h *handle) handleCompleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i := range h.tasks {
		if h.tasks[i].ID == id {
			h.tasks[i].Status = true

			resp, err := json.Marshal(h.tasks[i])
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			w.Header().Set("Content-Type", "application/json")

			w.WriteHeader(http.StatusOK)
			_, err = w.Write(resp)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return 
			}
				return }
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}

// Handler : To Delete Task by Id
func (h *handle) handleDeleteTask(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i := range h.tasks {

		if h.tasks[i].ID == id {
			h.tasks = append(h.tasks[:i], h.tasks[i+1:]...)

			w.WriteHeader(http.StatusCreated)

			str := fmt.Sprintf("Task with id : %d is deleted from task manager", id)
			_, err = w.Write([]byte(str))

			if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
			return}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}

// Handler : To get all the enrolled Tasks
func (h *handle) handleAllTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()

	resp, err := json.Marshal(h.tasks)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusFound)
	_, err = w.Write(resp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	h := &handle{}
	h.tasks = []task{
		{ID: idGen(), Desc: "Learn", Status: false},
		{ID: idGen(), Desc: "Write", Status: false},
		{ID: idGen(), Desc: "Test", Status: true},
	}

	http.HandleFunc("/taskall", h.handleAllTasks)
	http.HandleFunc("/taskadd", h.handleAddTask)
	http.HandleFunc("/taskget/{id}", h.handleGetTaskByID)
	http.HandleFunc("/task/{id}", h.handleCompleteTask)
	http.HandleFunc("/taskdlt/{id}", h.handleDeleteTask)

	fmt.Println("Server running at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
