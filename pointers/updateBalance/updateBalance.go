package main

import (
	"errors"
)

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

// Don't touch above this line

// ?
func updateBalance (cust *customer, trans transaction) error {
	switch trans.transactionType {
	case transactionDeposit:
		cust.balance += trans.amount
		return nil
	case transactionWithdrawal:
		if cust.balance < trans.amount{
			return errors.New("insufficient funds")
		}
		cust.balance -= trans.amount
		return nil
	default:
		return errors.New("unknown transaction type")
	}
}