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

// This function would behave identically if the type parameter were removed
// And first argument's type was bankAccount
func getBankAccountID[accountType bankAccount](account accountType) int {
	return account.getId()
}

func main() {
	myCheckingAccount := checkingAccount{id: 5}
	fmt.Println(getBankAccountID(myCheckingAccount))

	mySavingsAccount := savingsAccount{id: 6}
	fmt.Println(getBankAccountID(mySavingsAccount))
}
