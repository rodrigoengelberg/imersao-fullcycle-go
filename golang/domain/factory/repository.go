package factory

import "github.com/rodrigoengelberg/imersao-fullcycle-go/domain/repository"

type RepositoryFactory interface {
	CreateTranscationRepository() repository.TransactionRepository
}
