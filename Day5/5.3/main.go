package main

import (
	"fmt"
)

var amount int = 50000

type PaymentMethod interface {
	pay(amount float64) string
	shouldRequireOTP() bool
}

type creditCard struct {
	cardNumber int
}

type Paypal struct {
	id string
}

type Upi struct {
	id string
}

func (c *creditCard) pay(money float64) string {

	amount -= int(money)
	return fmt.Sprintf("Paid Rs. %v using card end with %v", money, c.cardNumber)

}

func (c *creditCard) shouldRequireOTP() bool {
	return true
}

func (p *Paypal) pay(money float64) string {

	amount -= int(money)
	return fmt.Sprintf("Paid Rs. %v Paypal account : %v", money, p.id)
}

func (*Paypal) shouldRequireOTP() bool {
	return false
}

func (u *Upi) pay(money float64) string {

	amount -= int(money)
	return fmt.Sprintf("Paid Rs. %v using UPI : %v", money, u.id)

}

func (u *Upi) shouldRequireOTP() bool {
	return true
}

func GenerateOTP() string {

	return fmt.Sprintln("OTP sent to registered number")
}

func main() {

	fmt.Println("---------- Payment Methods ---------- ")

	for {

		var input string

		var spendamount float64
		exit := false

		fmt.Println("Mode of payment : \nCreditcard : 1\nPayPal ; 2\nUpi : 3\nLeft amount : 4\nExit : 0 ")
		fmt.Scanf("%s", &input)

		switch input {
		case "1":
			fmt.Println("Enter Amount : ")
			fmt.Scanln(&spendamount)
			var id int
			fmt.Println("Enter Card Number : ")
			fmt.Scanln(&id)
			var p PaymentMethod = &creditCard{id}
			str := p.pay((spendamount))

			if p.shouldRequireOTP() {
				str := GenerateOTP()
				fmt.Println(str)
			}

			fmt.Println(str)

		case "2":
			fmt.Println("Enter Amount : ")
			fmt.Scanln(&spendamount)
			var id string
			fmt.Println("Enter id : ")
			fmt.Scanln(&id)
			var p PaymentMethod = &Paypal{id}
			str := p.pay((spendamount))

			if p.shouldRequireOTP() {
				str := GenerateOTP()
				fmt.Println(str)
			}
			fmt.Println(str)

		case "3":
			fmt.Println("Enter Amount : ")
			fmt.Scanln(&spendamount)
			var id string
			fmt.Println("Enter id : ")
			fmt.Scanln(&id)
			var p PaymentMethod = &Upi{id}
			str := p.pay((spendamount))

			if p.shouldRequireOTP() {
				str := GenerateOTP()
				fmt.Println(str)
			}
			fmt.Println(str)

			fmt.Println(str)

		case "4":
			fmt.Println("Remaining Amount : ", amount)

		case "0":
			exit = true

		}

		if exit {
			break
		}
	}

}
