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
func getBankAccountID(account bankAccount) int {
	return account.getId()
}

func main() {
	myCheckingAccount := checkingAccount{id: 5}
	fmt.Println(getBankAccountID(myCheckingAccount))

	mySavingsAccount := savingsAccount{id: 6}
	fmt.Println(getBankAccountID(mySavingsAccount))

}
