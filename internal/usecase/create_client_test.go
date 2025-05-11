package usecase

import (
	"errors"
	"testing"

	"github.com/CaioAugustoo/wallet-core/internal/gateway"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateClient(t *testing.T) {
	clientGateway := new(gateway.ClientGatewayMock)
	usecase := NewCreateClient(clientGateway)

	clientGateway.On("Save", mock.Anything).Return(nil)

	input := CreateClientInput{
		Name:  "John Doe",
		Email: "john@doe.com",
	}

	output, err := usecase.Execute(input)

	assert.NoError(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, output.Client.Name, input.Name)
	clientGateway.AssertNumberOfCalls(t, "Save", 1)
}

func TestCreateClientError(t *testing.T) {
	clientGateway := new(gateway.ClientGatewayMock)
	usecase := NewCreateClient(clientGateway)

	input := CreateClientInput{
		Name:  "",
		Email: "john@doe.com",
	}

	_, err := usecase.Execute(input)

	assert.Error(t, err)
	clientGateway.AssertNumberOfCalls(t, "Save", 0)
}

func TestCreateClientGatewayError(t *testing.T) {
	clientGateway := new(gateway.ClientGatewayMock)
	usecase := NewCreateClient(clientGateway)

	clientGateway.On("Save", mock.Anything).Return(errors.New("error"))

	input := CreateClientInput{
		Name:  "John Doe",
		Email: "john@doe.com",
	}

	_, err := usecase.Execute(input)

	assert.Error(t, err)
	clientGateway.AssertNumberOfCalls(t, "Save", 1)
}
