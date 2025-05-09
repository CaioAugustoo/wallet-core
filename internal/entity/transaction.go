package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID          string
	AccountFrom *Account
	AccountTo   *Account
	Amount      float64
	CreatedAt   time.Time
}

func NewTransaction(accountFrom, accountTo *Account, amount float64) (*Transaction, error) {
	transaction := &Transaction{
		ID:          uuid.NewString(),
		AccountFrom: accountFrom,
		AccountTo:   accountTo,
		Amount:      amount,
		CreatedAt:   time.Now(),
	}

	if err := transaction.Validate(); err != nil {
		return nil, err
	}

	return transaction, transaction.Commit()
}

func (t *Transaction) Validate() error {
	if t.AccountFrom == nil {
		return errors.New("account from is empty")
	}
	if t.AccountTo == nil {
		return errors.New("account to is empty")
	}
	if t.Amount <= 0 {
		return errors.New("amount must be greater than 0")
	}
	return nil
}

func (t *Transaction) Commit() error {
	if err := t.AccountFrom.Withdraw(t.Amount); err != nil {
		return err
	}
	t.AccountTo.Deposit(t.Amount)
	return nil
}
