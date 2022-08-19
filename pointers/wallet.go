package main

type Wallet struct {
	balance int
}

func (w Wallet) Deposit(to_deposit int) {
	w.balance += to_deposit
}

func (w Wallet) Balance() int {
	return w.balance
}
