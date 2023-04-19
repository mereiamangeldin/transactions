package service

import (
	"github.com/mereiamangeldin/transaction/model"
	"github.com/mereiamangeldin/transaction/repostiory"
)

type ITransactionService interface {
	ServeTransaction(transaction model.Transaction) error
	GetTransaction() ([]model.Transaction, error)
	GetStatistics() ([]model.Statistics, error)
}

type Manager struct {
	Transaction ITransactionService
}

func NewManager(repo *repostiory.Repository) (*Manager, error) {
	transaction := NewTransactionService(repo)
	return &Manager{Transaction: transaction}, nil
}
