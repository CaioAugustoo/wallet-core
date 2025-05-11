package usecase

import (
	"errors"
	"testing"

	"github.com/CaioAugustoo/wallet-core/internal/entity"
	"github.com/CaioAugustoo/wallet-core/internal/gateway"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateAccount(t *testing.T) {
	accountGateway := new(gateway.AccountGatewayMock)
	clientGateway := new(gateway.ClientGatewayMock)
	usecase := NewCreateAccount(accountGateway, clientGateway)

	clientID := uuid.NewString()

	accountGateway.On("Save", mock.Anything).Return(nil)
	clientGateway.On("FindById", mock.Anything).Return(&entity.Client{
		ID:    clientID,
		Name:  "John Doe",
		Email: "john@doe.com",
	}, nil)

	input := CreateAccountInput{
		ClientID: clientID,
	}

	output, err := usecase.Execute(input)

	assert.NoError(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	accountGateway.AssertNumberOfCalls(t, "Save", 1)
}

func TestCreateAccountClientGatewayError(t *testing.T) {
	accountGateway := new(gateway.AccountGatewayMock)
	clientGateway := new(gateway.ClientGatewayMock)
	usecase := NewCreateAccount(accountGateway, clientGateway)

	clientID := uuid.NewString()

	clientGateway.On("FindById", clientID).Return(&entity.Client{
		ID:    clientID,
		Name:  "John Doe",
		Email: "john@doe.com",
	}, errors.New("error"))

	input := CreateAccountInput{
		ClientID: clientID,
	}

	_, err := usecase.Execute(input)

	assert.Error(t, err)
}

func TestCreateAcountAccountGatewayError(t *testing.T) {
	accountGateway := new(gateway.AccountGatewayMock)
	clientGateway := new(gateway.ClientGatewayMock)
	usecase := NewCreateAccount(accountGateway, clientGateway)

	clientID := uuid.NewString()

	clientGateway.On("FindById", clientID).Return(&entity.Client{
		ID:    clientID,
		Name:  "John Doe",
		Email: "john@doe.com",
	}, nil)

	accountGateway.On("Save", mock.Anything).Return(errors.New("error"))

	input := CreateAccountInput{
		ClientID: clientID,
	}

	_, err := usecase.Execute(input)

	assert.Error(t, err)
}
