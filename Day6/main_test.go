package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"fmt"
)

func Test_handlePendingTasks(t *testing.T) {

	h := handle{
		ptr: &tasks{
			cont: []task{
				{Id: 1, Desc: "Sample Task", Status: false},
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(h.handlePendingTasks))

	resp, err := http.Get(server.URL)

	if err != nil {
		t.Error(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Unexpected Results ")
	}
}

func Test_handlePendingTask1(t *testing.T) {

	h := handle{
		ptr: &tasks{
			cont: []task{
				{Id: 3, Desc: "Grocery", Status: false},
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(h.handleAdd))
	body1, _ := json.Marshal(h.ptr.cont[0])

	req, err := http.NewRequest(http.MethodPost, server.URL, bytes.NewReader(body1))

	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	var t1 task

	err = json.Unmarshal(body, &t1)

	if err != nil {
		t.Fatal(err)
	}

	if t1 != h.ptr.cont[0] {

		t.Errorf("Unmatched Output")
	}

}

func Test_handleAdd(t *testing.T) {

	h := handle{
		ptr: &tasks{
			cont: []task{
				{Id: 3, Desc: "Grocery", Status: false},
			},
		},
	}

	req:=httptest.NewRequest("POST","http://localhost:8000/task1",nil)
	w := httptest.NewRecorder()

	h.handleAdd(w,req)
    
	resp:=w.Result()
	body,_:=io.ReadAll(resp.Body)


	fmt.Println(resp.StatusCode,string(body))



}