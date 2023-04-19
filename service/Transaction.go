package service

import (
	"github.com/mereiamangeldin/transaction/model"
	"github.com/mereiamangeldin/transaction/repostiory"
)

type TransactionService struct {
	Repo *repostiory.Repository
}

func (s *TransactionService) GetStatistics() ([]model.Statistics, error) {
	return s.Repo.Transaction.GetStatistics()
}

func (s *TransactionService) GetTransaction() ([]model.Transaction, error) {
	return s.Repo.Transaction.GetTransaction()
}

func (s *TransactionService) ServeTransaction(transaction model.Transaction) error {
	return s.Repo.Transaction.ServeTransaction(transaction)
}

func NewTransactionService(repo *repostiory.Repository) *TransactionService {
	return &TransactionService{Repo: repo}
}
