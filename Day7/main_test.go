package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var errTest = errors.New("error")

type errWriter struct {
	Code int
}

func (*errWriter) Header() http.Header {
	return http.Header{}
}
func (*errWriter) Write(_ []byte) (int, error) {
	return 0, errTest
}
func (e *errWriter) WriteHeader(statusCode int) {
	e.Code = statusCode
}

// GetTaskByID Method test
func Test_handleGetTaskById(t *testing.T) {

	h := &handle{tasks: []task{
		{ID: 1, Desc: "Learn", Status: false},
		{ID: 2, Desc: "Shopping", Status: false},
		{ID: 3, Desc: "Dance", Status: false},
	}}

	//  1. Wrong Method Case
	req1 := httptest.NewRequest(http.MethodPost, "/task/{id}", nil)
	req1.SetPathValue("id", "1")

	w1 := httptest.NewRecorder()

	h.handleGetTaskByID(w1, req1)

	if w1.Result().StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected %d but got %d (Wrong Method)", http.StatusMethodNotAllowed, w1.Result().StatusCode)
	}

	// 2. ID Not Passed Case
	req2 := httptest.NewRequest(http.MethodGet, "/task/{id}", nil)

	w2 := httptest.NewRecorder()

	h.handleGetTaskByID(w2, req2)

	if w2.Result().StatusCode != http.StatusBadRequest {
		t.Errorf("Expected %d but got %d (ID Not Passed)", http.StatusBadRequest, w2.Result().StatusCode)
	}

	// 3. Success Case
	req3 := httptest.NewRequest(http.MethodGet, "/task/{id}", nil)
	req3.SetPathValue("id", "1")

	w3 := httptest.NewRecorder()

	h.handleGetTaskByID(w3, req3)

	resp3 := w3.Result()
	body, err := io.ReadAll(resp3.Body)
	if err != nil {
		t.Error(err)
	}

	if resp3.StatusCode != http.StatusOK {
		t.Errorf("Expected %d, got %d", http.StatusOK, resp3.StatusCode)
	}

	var tsk task
	err = json.Unmarshal(body, &tsk)
	if err != nil {
		t.Error(err)
	}

	// 4. ID Not Found Case
	req4 := httptest.NewRequest(http.MethodGet, "/task/{id}", nil)
	req4.SetPathValue("id", "99")

	w4 := httptest.NewRecorder()

	h.handleGetTaskByID(w4, req4)

	if w4.Result().StatusCode != http.StatusNotFound {
		t.Errorf("Expected %d but got %d (ID Not Found)", http.StatusNotFound, w4.Result().StatusCode)
	}

	// 5. Write error
	req5 := httptest.NewRequest(http.MethodGet, "/task/{id}", nil)
	req5.SetPathValue("id", "1")

	w5 := errWriter{0}
	h.handleGetTaskByID(&w5, req5)

	if w5.Code != http.StatusInternalServerError {
		t.Errorf("Expected %d but got %d (ID Not Found)", http.StatusInternalServerError, w5.Code)
	}

	
}

// AddTask Method Test
type errReader int

func (errReader) Read([]byte) (n int, err error) {
	return 0, errors.New("test error")
}
func Test_handleAddTask(t *testing.T) {

	h := &handle{tasks: []task{
		{ID: 1, Desc: "Learn", Status: false},
		{ID: 2, Desc: "Shopping", Status: false},
		{ID: 3, Desc: "Dance", Status: false},
	}}

	// 1. Wrong Method Case
	req1 := httptest.NewRequest(http.MethodGet, "/task", nil)

	w1 := httptest.NewRecorder()

	h.handleAddTask(w1, req1)

	if w1.Result().StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected %d but got %d (Wrong Method)", http.StatusMethodNotAllowed, w1.Result().StatusCode)
	}

	// 2. Invalid Body Type Case
	t1 := []struct {
		id   int
		name string
	}{
		{id: 1, name: "ravi"},
	}
	bodyInvalid, _ := json.Marshal(t1)
	newReaderInvalid := bytes.NewReader(bodyInvalid)

	req2 := httptest.NewRequest(http.MethodPost, "/task", newReaderInvalid)

	w2 := httptest.NewRecorder()

	h.handleAddTask(w2, req2)

	if w2.Result().StatusCode != http.StatusBadRequest {
		t.Errorf("Expected %d but got %d (Invalid Body Type)", http.StatusBadRequest, w2.Result().StatusCode)
	}

	//  3. Body Read Error Case
	req3 := httptest.NewRequest(http.MethodPost, "/task", errReader(0))

	w3 := httptest.NewRecorder()

	h.handleAddTask(w3, req3)

	if w3.Result().StatusCode != http.StatusBadRequest {
		t.Errorf("Expected %d but got %d (Body Read Error)", http.StatusBadRequest, w3.Result().StatusCode)
	}

	//  4. Success Case
	body, _ := json.Marshal(h.tasks[0])
	newReader := bytes.NewReader(body)

	req4 := httptest.NewRequest(http.MethodPost, "/task", newReader)

	w4 := httptest.NewRecorder()

	h.handleAddTask(w4, req4)

	resp := w4.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected %d but got %d (Success)", http.StatusCreated, resp.StatusCode)
	}

	var tsk task
	err = json.Unmarshal(body, &tsk)
	if err != nil {
		t.Error(err)
	}
    
	// 5. Write error
	body5, _ := json.Marshal(task{ID: 100, Desc: "Write fail test", Status: false})
	req5 := httptest.NewRequest(http.MethodPost, "/taskadd/", bytes.NewReader(body5))

	w5 := errWriter{0}
	h.handleAddTask(&w5, req5)

	if w5.Code != http.StatusInternalServerError {
		t.Errorf("Expected %d but got %d ", http.StatusInternalServerError, w5.Code)
	}
}

// UpdateTask Method test
func Test_handleCompleteTask(t *testing.T) {

	h := &handle{tasks: []task{
		{ID: 1, Desc: "Learn", Status: false},
	}}

	// 1. Wrong Method
	req1 := httptest.NewRequest(http.MethodGet, "/task", nil)

	w1 := httptest.NewRecorder()

	h.handleCompleteTask(w1, req1)

	if w1.Result().StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected %d but got %d (Wrong Method)", http.StatusMethodNotAllowed, w1.Result().StatusCode)
	}

	// 2. ID Not Passed
	req2 := httptest.NewRequest(http.MethodPut, "/task/{id}", nil)

	w2 := httptest.NewRecorder()

	h.handleCompleteTask(w2, req2)

	if w2.Result().StatusCode != http.StatusBadRequest {
		t.Errorf("Expected %d but got %d (ID Not Passed)", http.StatusBadRequest, w2.Result().StatusCode)
	}

	//  3. Invalid ID (Not Found)
	req3 := httptest.NewRequest(http.MethodPut, "/task/{id}", nil)
	req3.SetPathValue("id", "5")

	w3 := httptest.NewRecorder()

	h.handleCompleteTask(w3, req3)

	if w3.Result().StatusCode != http.StatusNotFound {
		t.Errorf("Expected %d but got %d (Invalid ID)", http.StatusNotFound, w3.Result().StatusCode)
	}

	// 4. Success Case
	req5 := httptest.NewRequest(http.MethodPut, "/task/{id}", nil)
	req5.SetPathValue("id", "1")

	w5 := httptest.NewRecorder()

	h.handleCompleteTask(w5, req5)

	if w5.Result().StatusCode != http.StatusOK {
		t.Errorf("Expected %d but got %d (Success)", http.StatusOK, w5.Result().StatusCode)
	}

	// 5. Write error
	req6 := httptest.NewRequest(http.MethodPut, "/task/{id}", nil)
	req6.SetPathValue("id", "1")

	w6:= errWriter{0}
	h.handleCompleteTask(&w6, req6)

	if w6.Code != http.StatusInternalServerError {
		t.Errorf("Expected %d but got %d ", http.StatusInternalServerError, w6.Code)
	}
}

// AllTask Method testing
func Test_handleAllTasks(t *testing.T) {

	h := &handle{tasks: []task{
		{ID: 1, Desc: "Learn", Status: false},
	}}

	//  1. Wrong Method
	req1 := httptest.NewRequest(http.MethodPut, "/task/{id}", nil)
	req1.SetPathValue("id", "1")

	w1 := httptest.NewRecorder()

	h.handleAllTasks(w1, req1)

	if w1.Result().StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected %d but got %d (Wrong Method)", http.StatusMethodNotAllowed, w1.Result().StatusCode)
	}

	// 2. Success Case
	req2 := httptest.NewRequest(http.MethodGet, "/task", nil)

	w2 := httptest.NewRecorder()

	h.handleAllTasks(w2, req2)

	if w2.Result().StatusCode != http.StatusFound {
		t.Errorf("Expected %d but got %d (Success)", http.StatusFound, w2.Result().StatusCode)
	}

	// 3. Write error
	req6 := httptest.NewRequest(http.MethodGet, "/task/{id}", nil)
	req6.SetPathValue("id", "1")

	w6:= errWriter{0}
	h.handleAllTasks(&w6, req6)

	if w6.Code != http.StatusInternalServerError {
		t.Errorf("Expected %d but got %d ", http.StatusInternalServerError, w6.Code)
	}
}

// DeleteTask Method test
func Test_handleDeleteTask(t *testing.T) {

	h := &handle{tasks: []task{
		{ID: 1, Desc: "Learn", Status: false},
	    {2,"Split",true},},
		
	}

	// 1. Wrong Method
	req1 := httptest.NewRequest(http.MethodPut, "/taskdlt/{id}", nil)
	req1.SetPathValue("id", "1")

	w1 := httptest.NewRecorder()

	h.handleDeleteTask(w1, req1)

	if w1.Result().StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected %d but got %d (Wrong Method)", http.StatusMethodNotAllowed, w1.Result().StatusCode)
	}

	// 2. ID Not Passed
	req2 := httptest.NewRequest(http.MethodDelete, "/taskdlt/{id}", nil)

	w2 := httptest.NewRecorder()

	h.handleDeleteTask(w2, req2)

	if w2.Result().StatusCode != http.StatusBadRequest {
		t.Errorf("Expected %d but got %d (ID Not Passed)", http.StatusBadRequest, w2.Result().StatusCode)
	}

	// 3. Invalid ID
	req3 := httptest.NewRequest(http.MethodDelete, "/taskdlt/{id}", nil)
	req3.SetPathValue("id", "5")

	w3 := httptest.NewRecorder()

	h.handleDeleteTask(w3, req3)

	if w3.Result().StatusCode != http.StatusNotFound {
		t.Errorf("Expected %d but got %d (Invalid ID)", http.StatusNotFound, w3.Result().StatusCode)
	}

	// 4. Success Case
	req4 := httptest.NewRequest(http.MethodDelete, "/taskdlt/{id}", nil)
	req4.SetPathValue("id", "1")

	w4 := httptest.NewRecorder()

	h.handleDeleteTask(w4, req4)

	if w4.Result().StatusCode != http.StatusCreated {
		t.Errorf("Expected %d but got %d (Success)", http.StatusCreated, w4.Result().StatusCode)
	}

	// 5. Write error
	req6 := httptest.NewRequest(http.MethodDelete, "/taskdlt/{id}", nil)
	req6.SetPathValue("id", "2")

	w6:= errWriter{0}
	h.handleDeleteTask(&w6, req6)

	if w6.Code != http.StatusInternalServerError {
		t.Errorf("Expected %d but got %d ", http.StatusInternalServerError, w6.Code)
	}
}
