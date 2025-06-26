package main

import (
	"testing"
)

func newTestHandle() *handle {
	return &handle{
		ptr: &tasks{},
	}
}

func Test_addtask(t *testing.T) {
	h := newTestHandle()

	tests := []struct {
		desc string
		exp  string
	}{
		{"Buymilk", "Buymilk"},
		{"Writecode", "Writecode"},
	}

	for i, val := range tests {
		addtask(h, val.desc)

		if len(h.ptr.cont) != i+1 {
			t.Errorf("expected %d tasks, got %d", i+1, len(h.ptr.cont))
		}

		if h.ptr.cont[i].description != val.exp {
			t.Errorf("expected description %q, got %q", val.exp, h.ptr.cont[i].description)
		}
	}
}

func Test_addtask_ViaRun(t *testing.T) {
	h := newTestHandle()

	tests := []struct {
		name string
		desc string
		exp  string
	}{
		{"case1", "Buymilk", "Buymilk"},
		{"case2", "Writecode", "Writecode"},
	}

	for i, val := range tests {
		t.Run(val.name, func(t *testing.T) {
			addtask(h, val.desc)

			if h.ptr.cont[i].description != val.exp {
				t.Errorf("test %s failed: expected %q, got %q", val.name, val.exp, h.ptr.cont[i].description)
			}
		})
	}
}

func Test_CompleteTask(t *testing.T) {
	h := newTestHandle()

	addtask(h, "LearnGo")
	addtask(h, "Testapp")

	tests := []struct {
		id  int
		exp bool
	}{
		{1, true},
		{2, true},

	}

	for _, val := range tests {
		CompleteTask(h, val.id)

		found := false

		for _, task := range h.ptr.cont {
			if task.id == val.id {
				found = true

				if task.status != val.exp {
					t.Errorf("task %d: expected status %v, got %v", val.id, val.exp, task.status)
				}
			}
		}

		if !found {
			t.Errorf("task with id %d not found", val.id)
		}
	}
}

func Benchmark_addtask(b *testing.B) {
	h := newTestHandle()

	for i := 0; i < b.N; i++ {
		addtask(h, "Hello")
	}
}

func Benchmark_CompleteTask(b *testing.B) {
	h := newTestHandle()

	// Setup initial tasks
	for i := 0; i < 10; i++ {
		addtask(h, "Task")
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CompleteTask(h, 5)
	}
}
