package main

import "fmt"

type bankAccount interface {
	getId() int
}

type checkingAccount struct {
	id int
}
type savingsAccount struct {
	id int
}

func (c checkingAccount) getId() int {
	return c.id
}

func (s savingsAccount) getId() int {
	return s.id
}

func printAccountStatus(account bankAccount) {
	// this will fail
	// account.id

	// Convert into a checking account if possible and print
	checkingAccount, isCheckingAccount := account.(checkingAccount)

	if isCheckingAccount {
		fmt.Printf("This is a checking account with id %v\n", checkingAccount.id)
	}

	// Convert into a savings account if possible and print
	savingsAccount, isSavingsAccount := account.(savingsAccount)
	if isSavingsAccount {
		fmt.Printf("This is a savings account with id %v\n", savingsAccount.id)
	}

	// Converting a checkingAccount back into a bank account
	var acc bankAccount = &checkingAccount
	fmt.Println(acc.getId())
}

func main() {
	myCheckingAccount := checkingAccount{id: 5}
	printAccountStatus(myCheckingAccount)

	mySavingsAccount := savingsAccount{id: 6}
	printAccountStatus(mySavingsAccount)
}
