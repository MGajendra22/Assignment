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

func Test_addtask_ViaRun(t *testing.T) {
	t.Cleanup(func() {
		reset()
	})

	tests := []struct {
		name string
		desc string
		exp  string
	}{
		{"case1","Buymilk", "Buymilk"},
		{"case2","Writecode", "Write code"},
	}

	for i,val:=range tests{
		t.Run(val.name, func(t *testing.T){
			addtask(val.desc)
            
			if(cont[i].description!=val.exp){
				t.Errorf("Test Case (%d)=%s wants %s",i+1,cont[i].description,val.exp)
			}

		})
	}
}



func Benchmark_addtask(b *testing.B){
	for i:=0;i<b.N;i++ {
		addtask("Hello")
	}
}

func Benchmark_CompleteTask(b *testing.B){

	for i:=0;i<b.N;i++ {
		CompleteTask(2)
	}
}

func Test_CompleteTask(t *testing.T) {
	reset()

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
