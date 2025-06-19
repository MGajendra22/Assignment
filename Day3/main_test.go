package main

import (
	"testing"
)

func reset() {
	cont = []task{}
	temp_id = idgen()
}

func Test_addtask(t *testing.T) {
	t.Cleanup(func() {
		reset()
	})

	tests := []struct {
		desc string
		exp  string
	}{
		{"Buymilk", "Buymilk"},
		{"Writecode", "Writecode"},
	}

	for i, val := range tests {
		addtask(val.desc)

		if len(cont) != i+1 {
			t.Errorf("Test Case %v Failed: ", val)
		}

		if cont[i].description != val.exp {
			t.Errorf("Test Case %v Failed: ", val)
		}
	}
}

func Test_CompleteTask(t *testing.T) {
	t.Cleanup(func() {
		reset()
	})

	addtask("LearnGo")
	addtask("Testapp")

	tests := []struct {
		id  int
		exp bool
	}{
		{1, true},
		{2, true},
	}

	for i, val := range tests {
		CompleteTask(val.id)
		if cont[i].status != val.exp {
			t.Errorf("Test Case %v Failed: ", val)
		}
	}
}

