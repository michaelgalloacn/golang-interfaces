package main

import "fmt"

type bankAccount interface {
	getId() int
}

type checkingAccount struct {
	owner string
	id    int
}

func (account checkingAccount) getId() int {
	return account.id
}

func checkingAccountOnlyFunction(account checkingAccount) int {
	return account.id * account.id
}

func accountList[accountType bankAccount](accounts ...accountType) []accountType {
	var result []accountType
	for _, account := range accounts {
		result = append(result, account)
	}

	return result

}

func main() {
	accountOne := checkingAccount{id: 5}
	accountTwo := checkingAccount{id: 6}

	checkingAccountList := accountList(accountOne, accountTwo)

	for _, account := range checkingAccountList {
		fmt.Println(checkingAccountOnlyFunction(account))
	}

}
