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

func main() {
	myCheckingAccount := checkingAccount{id: 5}
	mySavingsAccount := savingsAccount{id: 6}

	var accountSlice []bankAccount = []bankAccount{myCheckingAccount, mySavingsAccount}
	for _, account := range accountSlice {
		fmt.Println(account.getId())
	}
}
