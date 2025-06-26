package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	// "strings"

	_ "github.com/go-sql-driver/mysql"
)

type task struct {
	ID     int    `json:"id"`
	Desc   string `json:"desc"`
	Status bool   `json:"status"`
}

type handle struct {
	db *sql.DB
}

// Handler: Add Task (POST)
func (h *handle) handleAddTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var t task
	if err := json.Unmarshal(body, &t); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if t.Desc == "" {
		http.Error(w, "Description cannot be empty", http.StatusBadRequest)
		return
	}

	result, err := h.db.Exec("INSERT INTO tasks (description, status) VALUES (?,?)", t.Desc, t.Status)
	if err != nil {
		http.Error(w, "Insert failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	t.ID = int(id)

	resp, _ := json.Marshal(t)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}

// Get Task by ID
func (h *handle) handleGetTaskByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var t task
	err = h.db.QueryRow("SELECT id, description, status FROM tasks WHERE id = ?", id).Scan(&t.ID, &t.Desc, &t.Status)
	if err == sql.ErrNoRows {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Query failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp, _ := json.Marshal(t)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

// Complete Task
func (h *handle) handleCompleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	result, err := h.db.Exec("UPDATE tasks SET status = true WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Update failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Task %d marked as complete", id)))
}

// Delete Task
func (h *handle) handleDeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	result, err := h.db.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Delete failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Task %d deleted", id)))
}

// Get All Tasks
func (h *handle) handleAllTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	rows, err := h.db.Query("SELECT id, description, status FROM tasks")
	if err != nil {
		http.Error(w, "Select failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tasks []task
	for rows.Next() {
		var t task
		if err := rows.Scan(&t.ID, &t.Desc, &t.Status); err != nil {
			http.Error(w, "Scan failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, t)
	}

	resp, _ := json.Marshal(tasks)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}



func main() {
	connString := "root:root123@tcp(localhost:3306)/test_db?parseTime=true"
	db, err := sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Connection failed:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Ping failed:", err)
	}

	db.Exec("DROP TABLE IF EXISTS tasks")
	_, err = db.Exec(`
		CREATE TABLE tasks (
			id INT AUTO_INCREMENT PRIMARY KEY,
			description VARCHAR(255),
			status BOOLEAN
		)
	`)
	if err != nil {
		log.Fatal("Table creation failed:", err)
	}

	h := &handle{db: db}

	http.HandleFunc("/taskall", h.handleAllTasks)
	http.HandleFunc("/taskadd", h.handleAddTask)
	http.HandleFunc("/taskget/{id}", h.handleGetTaskByID)
	http.HandleFunc("/task/{id}", h.handleCompleteTask)
	http.HandleFunc("/taskdlt/", h.handleDeleteTask)

	fmt.Println(" Server running at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
