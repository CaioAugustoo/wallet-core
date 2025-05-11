package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

const (
	ErrNameIsEmpty                 = "name is empty"
	ErrEmailIsEmpty                = "email is empty"
	ErrAccountDoesntBelongToClient = "account does not belong to client"
	ErrClientNotFound              = "client not found"
)

type Client struct {
	ID        string
	Name      string
	Email     string
	Accounts  []Account
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewClient(name, email string) (*Client, error) {
	client := &Client{
		ID:        uuid.NewString(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return client, client.Validate()
}

func (c *Client) Validate() error {
	if c.Name == "" {
		return errors.New(ErrNameIsEmpty)
	}
	if c.Email == "" {
		return errors.New(ErrEmailIsEmpty)
	}
	return nil
}

func (c *Client) Update(name, email string) error {
	c.UpdatedAt = time.Now()
	c.Name = name
	c.Email = email
	return c.Validate()
}

func (c *Client) AddAccount(account *Account) error {
	if account.Client.ID != c.ID {
		return errors.New(ErrAccountDoesntBelongToClient)
	}

	c.Accounts = append(c.Accounts, *account)
	return nil
}
