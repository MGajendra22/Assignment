package main

import(
	"fmt"
	"os"     // For exception Handling 
	"bufio"  // For opening , closing , reading file elements
	"time"   // To provide Current time 
)

func main(){

    var filename string // to store fileName

	//if we do not give any file path then it will throw error
	if len(os.Args)<2 {
		fmt.Println("Please give the file name ")
		os.Exit(1)
	}else {
        filename=os.Args[1]
	}

	fmt.Println("------ Log Analysis of file: log.txt ------\n")
	
	// Log file is opened via this and if not exist then again throw error
	readfile,err1:=os.Open(filename)
	if err1!=nil {
		os.Exit(1)
	}
    
	// To close the file at the end
	defer readfile.Close()
    
	// To scan every word one by one of file : log.txt
	Scanner:=bufio.NewScanner(readfile)
	Scanner.Split(bufio.ScanWords)

    // Counting of every type of Messages
	var errorcnt =0
	var warningcnt=0
	var infocnt=0 


    // Loop to check 
	for Scanner.Scan() {
       text:=Scanner.Text() // Space seperated Word

	   // Switch case for : every single message count
       switch text{
	   case "[INFO]": infocnt++
	   case "[WARNING]": warningcnt++
	   case "[ERROR]": errorcnt++
	   }
	}
    
	total_log_lines:=infocnt+warningcnt+errorcnt

     // print the output 
	fmt.Println("Error count : ",errorcnt," \n Info count : ",infocnt,"\n Warning count : ",warningcnt)
	fmt.Println("Total Numbers of log Lines : ",total_log_lines,"\n Error Percentage : ",float32((errorcnt*100)/total_log_lines),"\n Warning Percentage : ",float64((warningcnt*100)/total_log_lines),"\n Info percentage : ",float64((infocnt*100)/total_log_lines))
    
	// time variable to store current time and show the current time
	tm:=time.Now()
	fmt.Println("Analyzed at:",tm)

}