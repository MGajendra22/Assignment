package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

type errReader int

func (errReader) Read([]byte) (int, error) {
	return 0, io.ErrUnexpectedEOF
}

// type errWriter struct {
// 	Code int
// }

// func (*errWriter) Header() http.Header { 
// 	return http.Header{} 
// }
// func (*errWriter) Write([]byte) (int, error) { 
// 	return 0, io.ErrUnexpectedEOF
//  }
// func (e *errWriter) WriteHeader(statusCode int) {
// 	e.Code = statusCode
// }

func reset(t *testing.T) *sql.DB {
	db, err := sql.Open("mysql", "root:root123@tcp(localhost:3306)/test_db?parseTime=true")
	if err != nil {
		t.Fatalf("failed to connect to DB: %v", err)
	}
	_, _ = db.Exec("DROP TABLE IF EXISTS tasks")
	_, err = db.Exec(`
		CREATE TABLE tasks (
			id INT AUTO_INCREMENT PRIMARY KEY,
			description VARCHAR(255),
			status BOOLEAN
		)`)
	if err != nil {
		t.Fatalf("table creation failed: %v", err)
	}
	return db
}

func insertTestTask(t *testing.T, db *sql.DB, desc string, status bool) int {
	res, err := db.Exec("INSERT INTO tasks (description, status) VALUES (?, ?)", desc, status)
	if err != nil {
		t.Fatalf("insert failed: %v", err)
	}
	id, _ := res.LastInsertId()
	return int(id)
}

func Test_handleAddTask_DB(t *testing.T) {
	db := reset(t)
	defer db.Close()
	h := &handle{db: db}

	// 1. Wrong HTTP method
	req1 := httptest.NewRequest(http.MethodGet, "/taskadd", nil)
	res1 := httptest.NewRecorder()
	h.handleAddTask(res1, req1)
	if res1.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected 405, got %d", res1.Code)
	}

	// 2. Invalid JSON
	req2 := httptest.NewRequest(http.MethodPost, "/taskadd", bytes.NewReader([]byte("invalid-json")))
	res2 := httptest.NewRecorder()
	h.handleAddTask(res2, req2)
	if res2.Code != http.StatusBadRequest {
		t.Errorf("Expected 400, got %d", res2.Code)
	}

	// 3. Body read error
	req3 := httptest.NewRequest(http.MethodPost, "/taskadd", errReader(0))
	res3 := httptest.NewRecorder()
	h.handleAddTask(res3, req3)
	if res3.Code != http.StatusBadRequest {
		t.Errorf("Expected 400 for body read error, got %d", res3.Code)
	}

	// 4. Empty description (should trigger validation failure)
	bodyEmptyDesc, _ := json.Marshal(task{Desc: "", Status: true})
	req4 := httptest.NewRequest(http.MethodPost, "/taskadd", bytes.NewReader(bodyEmptyDesc))
	res4 := httptest.NewRecorder()
	h.handleAddTask(res4, req4)
	if res4.Code != http.StatusBadRequest {
		t.Errorf("Expected 400 for empty description, got %d", res4.Code)
	}

	// 5. Valid task insert (success case)
	validTask := task{Desc: "Test Add", Status: true}
	body, _ := json.Marshal(validTask)
	req5 := httptest.NewRequest(http.MethodPost, "/taskadd", bytes.NewReader(body))
	res5 := httptest.NewRecorder()
	h.handleAddTask(res5, req5)
	if res5.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", res5.Code)
	}


}


func Test_handleCompleteTask_DB(t *testing.T) {
	db := reset(t)
	defer db.Close()
	h := &handle{db: db}

	id := insertTestTask(t, db, "Complete Me", false)

	r := httptest.NewRequest(http.MethodPut, "/task/", nil)
	r.SetPathValue("id", strconv.Itoa(id))
	w := httptest.NewRecorder()
	h.handleCompleteTask(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}

	// invalid id
	r2 := httptest.NewRequest(http.MethodPut, "/task/", nil)
	r2.SetPathValue("id", "abc")
	w2 := httptest.NewRecorder()
	h.handleCompleteTask(w2, r2)
	if w2.Code != http.StatusBadRequest {
		t.Errorf("Expected 400, got %d", w2.Code)
	}

	// non-existing ID
	r3 := httptest.NewRequest(http.MethodPut, "/task/", nil)
	r3.SetPathValue("id", "999")
	w3 := httptest.NewRecorder()
	h.handleCompleteTask(w3, r3)
	if w3.Code != http.StatusNotFound {
		t.Errorf("Expected 404, got %d", w3.Code)
	}

	// wrong method
	r4 := httptest.NewRequest(http.MethodGet, "/task/", nil)
	w4 := httptest.NewRecorder()
	h.handleCompleteTask(w4, r4)
	if w4.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected 405 for wrong method, got %d", w4.Code)
	}
}

func Test_handleDeleteTask_DB(t *testing.T) {
	db := reset(t)
	defer db.Close()
	h := &handle{db: db}

	id := insertTestTask(t, db, "Delete Me", true)

	r := httptest.NewRequest(http.MethodDelete, "/taskdlt/", nil)
	r.SetPathValue("id", strconv.Itoa(id))
	w := httptest.NewRecorder()
	h.handleDeleteTask(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}

	// invalid id
	r2 := httptest.NewRequest(http.MethodDelete, "/taskdlt/", nil)
	r2.SetPathValue("id", "xyz")
	w2 := httptest.NewRecorder()
	h.handleDeleteTask(w2, r2)
	if w2.Code != http.StatusBadRequest {
		t.Errorf("Expected 400, got %d", w2.Code)
	}

	// non-existent ID
	r3 := httptest.NewRequest(http.MethodDelete, "/taskdlt/", nil)
	r3.SetPathValue("id", "888")
	w3 := httptest.NewRecorder()
	h.handleDeleteTask(w3, r3)
	if w3.Code != http.StatusNotFound {
		t.Errorf("Expected 404, got %d", w3.Code)
	}

	// wrong method
	r4 := httptest.NewRequest(http.MethodPost, "/taskdlt/", nil)
	w4 := httptest.NewRecorder()
	h.handleDeleteTask(w4, r4)
	if w4.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected 405 for wrong method, got %d", w4.Code)
	}
}

func Test_handleGetTaskByID_DB(t *testing.T) {
	db := reset(t)
	defer db.Close()
	h := &handle{db: db}

	id := insertTestTask(t, db, "Get Me", false)

	r := httptest.NewRequest(http.MethodGet, "/taskget/", nil)
	r.SetPathValue("id", strconv.Itoa(id))
	w := httptest.NewRecorder()
	h.handleGetTaskByID(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}

	// invalid ID
	r2 := httptest.NewRequest(http.MethodGet, "/taskget/", nil)
	r2.SetPathValue("id", "x")
	w2 := httptest.NewRecorder()
	h.handleGetTaskByID(w2, r2)
	if w2.Code != http.StatusBadRequest {
		t.Errorf("Expected 400, got %d", w2.Code)
	}

	// not found
	r3 := httptest.NewRequest(http.MethodGet, "/taskget/", nil)
	r3.SetPathValue("id", "999")
	w3 := httptest.NewRecorder()
	h.handleGetTaskByID(w3, r3)
	if w3.Code != http.StatusNotFound {
		t.Errorf("Expected 404, got %d", w3.Code)
	}

	// wrong method
	r4 := httptest.NewRequest(http.MethodPost, "/taskget/", nil)
	w4 := httptest.NewRecorder()
	h.handleGetTaskByID(w4, r4)
	if w4.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected 405 for wrong method, got %d", w4.Code)
	}
}

func Test_handleAllTasks_DB(t *testing.T) {
	db := reset(t)
	defer db.Close()
	h := &handle{db: db}

	insertTestTask(t, db, "Task1", false)
	insertTestTask(t, db, "Task2", true)

	r := httptest.NewRequest(http.MethodGet, "/taskall", nil)
	w := httptest.NewRecorder()
	h.handleAllTasks(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}

	body, err := io.ReadAll(w.Body)
	if err != nil {
		t.Fatalf("Failed to read response: %v", err)
	}

	var tasks []task
	err = json.Unmarshal(body, &tasks)
	if err != nil || len(tasks) < 2 {
		t.Errorf("Expected task list with >= 2 items, got %d and err: %v", len(tasks), err)
	}

	// wrong method
	r2 := httptest.NewRequest(http.MethodPost, "/taskall", nil)
	w2 := httptest.NewRecorder()
	h.handleAllTasks(w2, r2)
	if w2.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected 405 for wrong method, got %d", w2.Code)
	}
}
