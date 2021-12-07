package entity

import "errors"

const (
	REJECTED = "rejected"
	APPROVED = "approved"
)

type Transaction struct {
	ID           string
	AccountID    string
	Amount       float64
	CreditCard   CreditCard
	Status       string
	ErrorMessage string
}

func NewTransaction() *Transaction {
	return &Transaction{}
}

func (transaction *Transaction) IsValid() error {
	if transaction.Amount > 1000 {
		return errors.New("you dont have limit for this transaction")
	}

	if transaction.Amount < 1 {
		return errors.New("the amount must be greater than 0")
	}

	return nil
}

func (transaction *Transaction) SetCreditCard(card CreditCard) {
	transaction.CreditCard = card
}
