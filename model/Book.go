package model

import "time"

type Transaction struct {
	ID              uint      `json:"id"`
	BookID          uint      `json:"book_id"`
	ClientID        uint      `json:"client_id"`
	TransactionType string    `json:"transaction_type"`
	Amount          float64   `json:"amount"`
	TakenAt         time.Time `json:"taken_at"`
	ReturnedAt      time.Time `json:"returned_at"`
}

type Statistics struct {
	BookId      int
	TotalIncome float64
}
