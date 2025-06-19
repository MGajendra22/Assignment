package main

import (
	"testing"
)

func Test_circle_cal(t *testing.T){

	 test:=[]struct {
		rad float64
		exp float64

	 }{
        {0,0},
		{1,3.14},
		{2,12.56},
	 }

     for i,val:=range test{
		 c:=circle{float64(val.rad)}
		 out:=cal(&c)

		 if(out!=val.exp){
            t.Errorf("Test Case :%v failed",test[i])
		 }

	 }
	
	
}

func Test_rectangle_cal(t *testing.T){

	test:=[]struct{
		l,b float64
		exp float64

	}{
       {0,0,0},
	   {5,10,50},
	   {1,1,1},
	   {2,3,6},
	}

	for i,val:=range test{

		r:=rectangle{val.l,val.b}
		out:=cal(&r)

		if(out!=val.exp){
            t.Errorf("Test Case :%v failed",test[i])
		 }
	}
}