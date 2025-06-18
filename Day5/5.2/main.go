package main

import (
	"fmt"
)

type Logger interface{
	Log(message string)
}

type ConsoleLogger struct{
	message string
}

type FileLogger struct{
	message string
}



type RemoteLogger struct{
	message string
}

func (c *ConsoleLogger) Log(msg string) {
	c.message=msg
}

func (f *FileLogger) Log(msg string){
	f.message=msg
}

func (r *RemoteLogger) Log(msg string){
	r.message=msg
}

func LogAll(lg []Logger,message string){
     
	for i,_:=range lg{
		lg[i].Log(message)
	}

}

func main(){
   
   var sl []Logger

   c1:=&ConsoleLogger{}
   f1:=&FileLogger{}
   r1:=&RemoteLogger{}

   sl=append(sl,c1)
   sl=append(sl,f1)
   sl=append(sl,r1)

   var msg string

   fmt.Println("Write message you want to send ; ")
   fmt.Scanln(&msg)
   LogAll(sl,msg)

   fmt.Println("Console : ",c1.message)
   fmt.Println("FIle : ",f1.message)
   fmt.Println("Remote : ",r1.message)

   

}
