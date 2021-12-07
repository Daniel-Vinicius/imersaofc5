package entity

type TransactionRepository interface {
	Insert(id string, accountId string, amount float64) error
}
