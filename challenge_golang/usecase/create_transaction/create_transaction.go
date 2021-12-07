package create_transaction

import "github.com/Daniel-Vinicius/challenge-golang/entity"

type CreateTransaction struct {
	Repository entity.TransactionRepository
}

func NewCreateTransactionUseCase(repository entity.TransactionRepository) *CreateTransaction {
	return &CreateTransaction{Repository: repository}
}

func (createTransaction *CreateTransaction) Execute(input TransactionDtoInput) error {
	errorTryingInsertIntoDatabase := createTransaction.Repository.Insert(input.ID, input.AccountID, input.Amount)

	if errorTryingInsertIntoDatabase != nil {
		return errorTryingInsertIntoDatabase
	}

	return nil
}
