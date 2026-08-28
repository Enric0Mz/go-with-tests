package main

import (
	"errors"
	"fmt"
)

var withdrawErr = errors.New("Insuficient balance to withdraw")

type Stringer interface {
	String() string
}

type Wallet struct {
	balance Bitcoin
}

type Bitcoin int

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return withdrawErr
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func ProcessWithdraw(amount Bitcoin, accountID string, wallet *Wallet) error {
	err := wallet.Withdraw(amount)

	if err != nil {
		return fmt.Errorf("account %s has insufficient funds in wallet: %w", accountID, err)
	}
	return nil
}

func main() {
	wallet := Wallet{Bitcoin(10)}

	err := ProcessWithdraw(100, "acc-xyz", &wallet)

	if err == withdrawErr {
		fmt.Print("Conventional IF WORKS")
	} else {
		fmt.Printf("Conventional IF DOES NOT WORK")
	}

	if errors.Is(err, withdrawErr) {
		fmt.Printf("This is the correct way to check wrapped errors")
	}
}
