package repostiory

import (
	"github.com/mereiamangeldin/transaction/config"
	"github.com/mereiamangeldin/transaction/model"
	"github.com/mereiamangeldin/transaction/repostiory/postgres"
)

type ITransactionRepository interface {
	ServeTransaction(transaction model.Transaction) error
	GetTransaction() ([]model.Transaction, error)
	GetStatistics() ([]model.Statistics, error)
}

type Repository struct {
	Transaction ITransactionRepository
}

func New(cfg *config.Config) (*Repository, error) {
	pgDb, err := postgres.Dial(cfg.PgURL)
	if err != nil {
		return nil, err
	}
	transaction := postgres.NewTransactionRepository(pgDb)
	return &Repository{Transaction: transaction}, nil
}
