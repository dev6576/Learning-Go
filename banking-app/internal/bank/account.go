package bank

import "sync"

type Account struct {
	ID      string
	Balance int64
	mu      sync.Mutex
}

func (a *Account) Deposit(amount int64) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Balance += amount
}

func (a *Account) Withdraw(amount int64) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.Balance < amount {
		return ErrInsufficientFunds
	}
	a.Balance -= amount
	return nil
}
