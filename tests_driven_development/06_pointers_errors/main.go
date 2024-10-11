package main

import (
	"errors"
	"fmt"
)

// create a new type from existing one
type Bitcoin int

type Wallet struct {
	balance Bitcoin // private
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// function to deposit
func (w *Wallet) Deposit(amount Bitcoin) {
	// fmt.Printf("address of balance in code: %p\n", &w.balance)
	w.balance += amount
}

// function to withdraw
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds") // global variable

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

// function to check balance
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
