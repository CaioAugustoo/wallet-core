package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewAccount(t *testing.T) {
	client, err := NewClient("Caio", "caio@rocha.me")
	assert.NoError(t, err)
	account, err := NewAccount(client)
	assert.NoError(t, err)
	assert.Equal(t, "Caio", account.Client.Name)
	assert.Equal(t, "caio@rocha.me", account.Client.Email)
	assert.NotEmpty(t, account.ID)
	assert.NotEmpty(t, account.CreatedAt)
	assert.NotEmpty(t, account.UpdatedAt)
}

func TestCreateNewAccountWhenClientIsEmpty(t *testing.T) {
	_, err := NewAccount(nil)
	assert.Error(t, err)
}

func TestUpdateAccount(t *testing.T) {
	client, err := NewClient("Caio", "caio@rocha.me")
	assert.NoError(t, err)
	account, err := NewAccount(client)
	assert.NoError(t, err)

	assert.Equal(t, account.Balance, 0.0)

	err = account.Update(10)
	assert.NoError(t, err)
	assert.Equal(t, account.Balance, 10.0)
}

func TestDeposit(t *testing.T) {
	client, err := NewClient("Caio", "caio@rocha.me")
	assert.NoError(t, err)
	account, err := NewAccount(client)
	assert.NoError(t, err)

	assert.Equal(t, account.Balance, 0.0)
	account.Deposit(10)
	account.Deposit(10)
	account.Deposit(10)

	assert.Equal(t, account.Balance, 30.0)
}

func TestWithdraw(t *testing.T) {
	client, err := NewClient("Caio", "caio@rocha.me")
	assert.NoError(t, err)
	account, err := NewAccount(client)
	assert.NoError(t, err)

	assert.Equal(t, account.Balance, 0.0)
	account.Deposit(10)
	account.Deposit(10)

	assert.Equal(t, account.Balance, 20.0)

	err = account.Withdraw(10)
	assert.NoError(t, err)
	assert.Equal(t, account.Balance, 10.0)
}

func TestWithdrawError(t *testing.T) {
	client, err := NewClient("Caio", "caio@rocha.me")
	assert.NoError(t, err)
	account, err := NewAccount(client)
	assert.NoError(t, err)

	assert.Equal(t, account.Balance, 0.0)
	account.Deposit(10)
	account.Deposit(10)

	assert.Equal(t, account.Balance, 20.0)

	err = account.Withdraw(30)
	assert.Error(t, err)
}
