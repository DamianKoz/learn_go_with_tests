package main

import "fmt"

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(to_deposit Bitcoin) {
	w.balance += to_deposit
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
