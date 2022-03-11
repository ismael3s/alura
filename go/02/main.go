package main

import (
	"fmt"

	"github.com/ismael3s/alura/go/02/accounts"
	"github.com/ismael3s/alura/go/02/customers"
)

func Pay(account VerifyAccount, value float64) {
	if value < 0 {
		fmt.Println("Invalid value")
		return
	}
	account.Withdraw(value)
}

type VerifyAccount interface {
	Withdraw(value float64) error
}

func main() {
	customerOne := customers.Customer{Name: "Ismael", CPF: "123.456.789-00", Job: "Developer"}

	savingAccount := accounts.SavingAccount{Customer: customerOne, Balance: 1000}

	Pay(&savingAccount, 100)

	fmt.Println(savingAccount.GetBalance())
}
