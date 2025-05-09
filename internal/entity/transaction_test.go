package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewTransaction(t *testing.T) {
	client1, err := NewClient("Caio", "caio@rocha.me")
	assert.NoError(t, err)

	accountFrom, err := NewAccount(client1)
	assert.NoError(t, err)

	client2, err := NewClient("Caio", "caio@rocha.me")
	assert.NoError(t, err)

	accountTo, err := NewAccount(client2)
	assert.NoError(t, err)

	accountFrom.Deposit(1000)

	transaction, err := NewTransaction(accountFrom, accountTo, 10)
	assert.NoError(t, err)
	assert.Equal(t, transaction.AccountFrom, accountFrom)
	assert.Equal(t, transaction.AccountTo, accountTo)
	assert.Equal(t, transaction.Amount, 10.0)
	assert.Equal(t, accountFrom.Balance, 990.0)
	assert.Equal(t, accountTo.Balance, 10.0)
	assert.NotEmpty(t, transaction.ID)
	assert.NotEmpty(t, transaction.CreatedAt)
}
