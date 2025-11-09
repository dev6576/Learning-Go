package bank

import (
	"sync"
)

type Bank struct {
	accounts map[string]*Account
	mu       sync.RWMutex
}

func NewBank() *Bank {
	return &Bank{
		accounts: make(map[string]*Account),
	}
}

func (b *Bank) CreateAccount(id string) *Account {
	b.mu.Lock()
	defer b.mu.Unlock()
	acc := &Account{ID: id}
	b.accounts[id] = acc
	return acc
}

func (b *Bank) GetAccount(id string) (*Account, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	acc, exists := b.accounts[id]
	if !exists {
		return nil, ErrAccountNotFound
	}
	return acc, nil
}

func (b *Bank) Transfer(from, to string, amount int64) error {
	accFrom, err := b.GetAccount(from)
	if err != nil {
		return err
	}

	accTo, err := b.GetAccount(to)
	if err != nil {
		return err
	}

	if err := accFrom.Withdraw(amount); err != nil {
		return err
	}

	accTo.Deposit(amount)
	return nil
}

func (b *Bank) Deposit(id string, amount int64) error {
	acc, err := b.GetAccount(id)
	if err != nil {
		return err
	}

	acc.Deposit(amount)
	return nil
}
