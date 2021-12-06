package factory

import "github.com/Daniel-Vinicus/imersaofc5/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
