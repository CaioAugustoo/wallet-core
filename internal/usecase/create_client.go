package usecase

import (
	"github.com/CaioAugustoo/wallet-core/internal/entity"
	"github.com/CaioAugustoo/wallet-core/internal/gateway"
)

type createClient struct {
	clientGateway gateway.Client
}

type CreateClient interface {
	Execute(input CreateClientInput) (*CreateClientOutput, error)
}

type CreateClientInput struct {
	Name  string
	Email string
}

type CreateClientOutput struct {
	Client *entity.Client
}

func NewCreateClient(clientGateway gateway.Client) CreateClient {
	return &createClient{
		clientGateway: clientGateway,
	}
}

func (uc *createClient) Execute(input CreateClientInput) (*CreateClientOutput, error) {
	client, err := entity.NewClient(input.Name, input.Email)
	if err != nil {
		return nil, err
	}

	if err := uc.clientGateway.Save(client); err != nil {
		return nil, err
	}

	return &CreateClientOutput{
		Client: client,
	}, nil
}
