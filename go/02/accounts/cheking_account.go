package accounts

import (
	"fmt"

	"github.com/ismael3s/alura/go/02/customers"
)

type CheckingAccount struct {
	Customer      customers.Customer
	AgencyNumber  int
	AccountNumber int
	Operation     int
	Balance       float64
}

func (a *CheckingAccount) Withdraw(value float64) error {
	canWithdraw := (a.Balance-value) >= 0 && value > 0

	if !canWithdraw {
		return fmt.Errorf("not enough balance")
	}

	a.Balance -= value

	return nil
}

func (a *CheckingAccount) Deposit(value float64) error {
	if value < 0 {
		return fmt.Errorf("invalid value")
	}

	a.Balance += value

	return nil
}

func (a *CheckingAccount) Transfer(value float64, account *CheckingAccount) error {
	if value < 0 {
		return fmt.Errorf("invalid value")
	}

	if a.Balance < value {
		return fmt.Errorf("not enough balance")
	}

	a.Balance -= value
	account.Deposit(value)

	return nil

}

func (a CheckingAccount) GetBalance() float64 {
	return a.Balance
}
