package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

const (
	ErrNameIsEmpty  = "name is empty"
	ErrEmailIsEmpty = "email is empty"
)

type Client struct {
	ID        string
	Name      string
	Email     string
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
