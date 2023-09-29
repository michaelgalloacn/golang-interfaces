package main

const minimumBalance = 100.00
const minimumBalanceFee = 35

type minimumFeeEligibleBankAccount interface {
	bankAccount
	checkingAccount
}

type bankAccount interface {
	getBalance() float32
}

// Legacy Checking Account From OB Acquisition - Business requirement not to assess fees
type obCheckingAccount struct {
	id      int
	balance float32
}

type checkingAccount struct {
	id      int
	balance float32
}

func (account obCheckingAccount) getBalance() float32 {
	return account.balance

}

func (account checkingAccount) getBalance() float32 {
	return account.balance
}

func assessMinimumBalanceFee[accountType minimumFeeEligibleBankAccount](account accountType) float32 {
	if account.getBalance() < minimumBalance {
		return 35
	}
	return 0

}

func main() {
	// Legacy Account does not meet the type constraint, so compiler will not allow it
	// To be used in assessMinimumBalanceFee
	legacyAccount := obCheckingAccount{balance: 90}
	assessMinimumBalanceFee(legacyAccount)

	//
	newAccount := checkingAccount{balance: 90}
	assessMinimumBalanceFee(newAccount)

}
