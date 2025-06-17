package main

import "fmt"

//Struct
type BankAccount struct {
	Owner   string
	Balance float64
}

// without reference
func (b BankAccount) DisplayBalance() {
	fmt.Printf("Owner: %s, Balance: %.2f\n", b.Owner, b.Balance)
}

// with reference
func (b *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		b.Balance += amount
		fmt.Printf("Deposited %.2f to %s's account\n", amount, b.Owner)
	}
}

// with reference
func (b *BankAccount) Withdraw(amount float64) {
	if amount <= b.Balance {
		b.Balance -= amount
		fmt.Printf("Withdrew %.2f from %s's account\n", amount, b.Owner)
	} else {
		fmt.Println("Insufficient funds")
	}
}

func main() {

	// Create account as value
	acc := BankAccount{Owner: "Alice", Balance: 100.0}

	acc.DisplayBalance() 

	
	acc.Deposit(50.0)       
	acc.DisplayBalance()    


	acc.Withdraw(30.0)      
	acc.DisplayBalance()    

	fmt.Println("\n")

	
	copyAcc := acc
	copyAcc.Deposit(100.0)  
	copyAcc.DisplayBalance()
	acc.DisplayBalance()   
}
