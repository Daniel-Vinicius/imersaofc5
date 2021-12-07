package entity

type Transaction struct {
	ID        string
	AccountID string
	Amount    string
}

func NewTransaction() *Transaction {
	return &Transaction{}
}
