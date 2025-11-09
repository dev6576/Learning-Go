package bank

import "time"

type TransactionStatus string

const (
	Pending TransactionStatus = "PENDING"
	Success TransactionStatus = "SUCCESS"
	Failed  TransactionStatus = "FAILED"
)

type Transaction struct {
	ID        string            `json:"id"`
	From      string            `json:"from"`
	To        string            `json:"to"`
	Amount    float64           `json:"amount"`
	Status    TransactionStatus `json:"status"`
	Error     string            `json:"error,omitempty"`
	CreatedAt time.Time         `json:"created_at"`
}
