package factory

import "github.com/sSchmidtT/imersaofc5/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
