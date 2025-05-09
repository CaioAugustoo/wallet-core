package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        string
	Client    *Client
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAccount(client *Client) (*Account, error) {
	account := &Account{
		ID:        uuid.NewString(),
		Client:    client,
		Balance:   0.0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return account, account.Validate()
}

func (a *Account) Validate() error {
	if a.Client == nil {
		return errors.New("client is empty")
	}
	return nil
}

func (a *Account) Update(balance float64) error {
	a.UpdatedAt = time.Now()
	a.Balance = balance
	return a.Validate()
}

func (a *Account) Deposit(amount float64) {
	a.UpdatedAt = time.Now()
	a.Balance += amount
}

func (a *Account) Withdraw(amount float64) error {
	if a.Balance < amount {
		return errors.New("insufficient balance")
	}
	a.UpdatedAt = time.Now()
	a.Balance -= amount
	return nil
}
