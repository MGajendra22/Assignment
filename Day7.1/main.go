package main

import (
	"fmt"
	"time"
)

func display(s string){

	for i:=0;i<5;i++{
        time.Sleep(500*time.Millisecond)
		fmt.Println(s)
	}
}

func main(){
	go display("Go routines")
    // display("Hello World")
    
	// //anonymous GoRoutine
	// go func (val int){
	// 	for i:=0;i<val;i++{
		  
    //        fmt.Println(i)
	// 	}
	// }(10)

	//  time.Sleep(2*time.Millisecond)

	fmt.Println("Main function is executed")
}