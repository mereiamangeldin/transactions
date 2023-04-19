package postgres

import (
	"github.com/mereiamangeldin/transaction/model"
	"gorm.io/gorm"
	"time"
)

type TransactionRepository struct {
	db *gorm.DB
}

func (r *TransactionRepository) GetStatistics() ([]model.Statistics, error) {
	var statistics []model.Statistics
	//err := r.db.Raw("select book_id, sum(amount)/count(*) as total_income from transactions group by book_id").Save(&statistics).Error
	err := r.db.Model(&model.Transaction{}).Select("book_id, sum(amount) as total_income").Group("book_id").Scan(&statistics).Error
	if err != nil {
		return nil, err
	}
	return statistics, nil
}

func (r *TransactionRepository) GetTransaction() ([]model.Transaction, error) {
	var transactions []model.Transaction
	err := r.db.Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, err
}

func (r *TransactionRepository) ServeTransaction(transaction model.Transaction) error {
	transaction.TakenAt = time.Now()
	err := r.db.Create(&transaction)
	if err.Error != nil {
		return err.Error
	}

	return nil
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}
