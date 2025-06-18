package main

import (
	"fmt"
)
var amount int =50000

type PaymentMethod interface{
	pay (amount float64) string

}

type creditCard struct {
	cardNumber int 
	otp bool
    
}

type Paypal struct {
	id string
	otp bool
}

type Upi struct {
	id string
	otp bool
}


func (c *creditCard) pay(money float64) string{

	amount-=int(money)
	return fmt.Sprintf("Paid Rs. %v using card end with %v",money,c.cardNumber)

}

	
func (p *Paypal) pay(money float64) string{
    
    amount-=int(money)
	return fmt.Sprintf("Paid Rs. %v Paypal account : %v",money,p.id)
}


func (u* Upi) pay(money float64) string{

	amount-=int(money)
	return fmt.Sprintf("Paid Rs. %v using UPI : %v",money,u.id)

}


func GenerateOTP() string{

	return fmt.Sprintln("OTP sent to registered number")
}


func main(){

	fmt.Println("---------- Payment Methods ---------- \n")

	for {
        
		var input string
	
		var spendamount float64
		exit:=false


		fmt.Println("Mode of payment : \nCreditcard : 1\nPayPal ; 2\nUpi : 3\nLeft amount : 4\nExit : 0 ")
		fmt.Scanf("%s",&input)


		switch input {
		case "1":
			fmt.Println("Enter Amount : ")
			fmt.Scanln(&spendamount)
			var id int
			fmt.Println("Enter Card Number : ")
			fmt.Scanln(&id)
			var p PaymentMethod =&creditCard{id,true}
			str:=p.pay((spendamount))

		    p1,_:=p.(*creditCard)

			if(p1.otp){
				str:=GenerateOTP()
				fmt.Println(str)
			}

			
            
			fmt.Println(str,"\n")
		
		case "2":
			fmt.Println("Enter Amount : ")
			fmt.Scanln(&spendamount)
			var id string
			fmt.Println("Enter id : ")
			fmt.Scanln(&id)
			var p PaymentMethod=&Paypal{id,false}
			str:=p.pay((spendamount))

			p1,_:=p.(*Paypal)
			if(p1.otp){
				str:=GenerateOTP()
				fmt.Println(str)
			}
			fmt.Println(str,"\n")

	    case "3":
			fmt.Println("Enter Amount : ")
			fmt.Scanln(&spendamount)
			var id string
			fmt.Println("Enter id : ")
			fmt.Scanln(&id)
			var p PaymentMethod=&Upi{id,true}
			str:=p.pay((spendamount))

			p1,_:=p.(*Upi)
			if(p1.otp){
				str:=GenerateOTP()
				fmt.Println(str)
			}
			fmt.Println(str,"\n")

			fmt.Println(str,"\n")

		case "4":
			fmt.Println("Remaining Amount : ",amount,"\n")	


		case "0":
			 exit=true

		}
		
		if(exit){
			break
		}
	}
	
}

