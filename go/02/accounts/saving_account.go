package accounts

import (
	"fmt"

	"github.com/ismael3s/alura/go/02/customers"
)

type SavingAccount struct {
	Customer      customers.Customer
	AgencyNumber  int
	AccountNumber int
	Operation     int
	Balance       float64
}

func (s *SavingAccount) Withdraw(value float64) error {
	canWithdraw := (s.Balance-value) >= 0 && value > 0

	if !canWithdraw {
		return fmt.Errorf("not enough balance")
	}

	s.Balance -= value

	return nil
}

func (s *SavingAccount) Deposit(value float64) error {
	if value < 0 {
		return fmt.Errorf("invalid value")
	}

	s.Balance += value

	return nil
}

func (s *SavingAccount) GetBalance() float64 {
	return s.Balance
}
