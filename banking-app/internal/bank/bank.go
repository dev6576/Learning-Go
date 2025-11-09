package bank

import (
	"errors"
	"sync"
)

type Bank struct {
	accounts     map[string]*Account
	Mu           sync.RWMutex
	Transactions chan Transaction
	TxStore      map[string]*Transaction
}

func NewBank() *Bank {
	b := &Bank{
		accounts:     make(map[string]*Account),
		Transactions: make(chan Transaction, 100), // buffered
		TxStore:      make(map[string]*Transaction),
	}

	return b
}

func (b *Bank) StartWorkers(n int) {
	for i := 0; i < n; i++ {
		go b.processTransactions()
	}
}

func (b *Bank) processTransactions() {
	for tx := range b.Transactions {
		err := b.transfer(tx.From, tx.To, tx.Amount)
		b.Mu.Lock()
		if err != nil {
			tx.Status = Failed
			tx.Error = err.Error()
		} else {
			tx.Status = Success
		}
		b.TxStore[tx.ID] = &tx
		b.Mu.Unlock()
	}
}

func (b *Bank) transfer(from, to string, amount float64) error {
	fromAcc, ok := b.accounts[from]
	if !ok {
		return errors.New("source account does not exist")
	}
	toAcc, ok := b.accounts[to]
	if !ok {
		return errors.New("destination account does not exist")
	}
	if fromAcc.Balance < amount {
		return errors.New("insufficient balance")
	}

	fromAcc.Balance -= amount
	toAcc.Balance += amount
	return nil
}

func (b *Bank) CreateAccount(id string) *Account {
	b.Mu.Lock()
	defer b.Mu.Unlock()
	acc := &Account{ID: id}
	b.accounts[id] = acc
	return acc
}

func (b *Bank) GetAccount(id string) (*Account, error) {
	b.Mu.RLock()
	defer b.Mu.RUnlock()

	acc, exists := b.accounts[id]
	if !exists {
		return nil, ErrAccountNotFound
	}
	return acc, nil
}

func (b *Bank) Deposit(id string, amount float64) error {
	acc, err := b.GetAccount(id)
	if err != nil {
		return err
	}

	acc.Deposit(amount)

	return nil
}

func (b *Bank) Withdraw(id string, amount float64) error {
	acc, err := b.GetAccount(id)
	if err != nil {
		return err
	}

	acc.Withdraw(amount)

	return nil
}
