package usecase

import (
	"errors"
	"testing"

	"github.com/CaioAugustoo/wallet-core/internal/entity"
	"github.com/CaioAugustoo/wallet-core/internal/gateway"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTransaction(t *testing.T) {
	accountGateway := new(gateway.AccountGatewayMock)
	transactionGateway := new(gateway.TransactionGatewayMock)
	usecase := NewCreateTransaction(accountGateway, transactionGateway)

	accountGateway.On("FindById", mock.Anything).Return(&entity.Account{
		ID: "1ff8eaea-fc57-49ef-961e-1a83f0a84950",
	}, nil)

	transactionGateway.On("Create", mock.Anything).Return(nil)

	input := CreateTransactionInput{
		AccountIDFrom: "1ff8eaea-fc57-49ef-961e-1a83f0a84950",
		AccountIDTo:   "1ff8eaea-fc57-49ef-961e-1a83f0a84950",
		Amount:        100.0,
	}

	output, err := usecase.Execute(input)

	assert.NoError(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.TransactionID)
	transactionGateway.AssertNumberOfCalls(t, "Create", 1)
}

func TestCreateTransactionAccountGatewayError(t *testing.T) {
	accountGateway := new(gateway.AccountGatewayMock)
	transactionGateway := new(gateway.TransactionGatewayMock)
	usecase := NewCreateTransaction(accountGateway, transactionGateway)

	transactionGateway.On("Create", mock.Anything).Return(errors.New("error"))

	input := CreateTransactionInput{
		AccountIDFrom: "1ff8eaea-fc57-49ef-961e-1a83f0a84950",
		AccountIDTo:   "1ff8eaea-fc57-49ef-961e-1a83f0a84950",
		Amount:        100.0,
	}

	accountGateway.On("FindById", input.AccountIDFrom).Return(&entity.Account{}, errors.New("error"))

	_, err := usecase.Execute(input)

	assert.Error(t, err)
}

func TestCreateTransactionGatewayError(t *testing.T) {
	accountGateway := new(gateway.AccountGatewayMock)
	transactionGateway := new(gateway.TransactionGatewayMock)
	usecase := NewCreateTransaction(accountGateway, transactionGateway)

	input := CreateTransactionInput{
		AccountIDFrom: "1ff8eaea-fc57-49ef-961e-1a83f0a84950",
		AccountIDTo:   "1ff8eaea-fc57-49ef-961e-1a83f0a84950",
		Amount:        100.0,
	}

	accountGateway.On("FindById", input.AccountIDFrom).Return(&entity.Account{}, nil)
	transactionGateway.On("Create", mock.Anything).Return(errors.New("error"))

	_, err := usecase.Execute(input)

	assert.Error(t, err)
}
