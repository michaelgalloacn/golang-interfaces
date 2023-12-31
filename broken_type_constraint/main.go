package main

import "fmt"

type bankAccount interface {
	checkingAccount | savingsAccount
	getId() int
}

type checkingAccount struct {
	id int
}
type savingsAccount struct {
	id int
}

type faceBookAccount struct {
	id int
}

func (f faceBookAccount) getId() int {
	return f.id
}
func (c checkingAccount) getId() int {
	return c.id
}

func getBankAccountID[T bankAccount](b T) int {
	return b.getId()
}

func main() {
	myAccount := faceBookAccount{id: 5}
	fmt.Println(getBankAccountID(myAccount))
}
