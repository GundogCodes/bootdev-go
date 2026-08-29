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

func updateBalance(cust *customer, trans transaction) error {
	if (*cust).balance < trans.amount {
		return errors.New("insufficient funds")

	} else if trans.transactionType != "deposit" && trans.transactionType != "withdrawal" {
		return errors.New("unknown transaction type")
		
	} else if trans.transactionType == "deposit" {
		(*cust).balance = (*cust).balance + trans.amount
		return nil

	} else if trans.transactionType == "withdrawal" {
		(*cust).balance = (*cust).balance - trans.amount
		return nil

	} else {
		return nil
	}
}

