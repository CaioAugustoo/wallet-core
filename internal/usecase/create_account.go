package usecase

import (
	"github.com/CaioAugustoo/wallet-core/internal/entity"
	"github.com/CaioAugustoo/wallet-core/internal/gateway"
)

type createAccount struct {
	accountGateway gateway.Account
	clientGateway  gateway.Client
}

type CreateAccount interface {
	Execute(input CreateAccountInput) (*CreateAccountOutput, error)
}

type CreateAccountInput struct {
	ClientID string
}

type CreateAccountOutput struct {
	ID string
}

func NewCreateAccount(
	accountGateway gateway.Account,
	clientGateway gateway.Client,
) CreateAccount {
	return &createAccount{
		accountGateway: accountGateway,
		clientGateway:  clientGateway,
	}
}

func (uc *createAccount) Execute(input CreateAccountInput) (*CreateAccountOutput, error) {
	client, err := uc.clientGateway.FindById(input.ClientID)
	if err != nil {
		return nil, err
	}

	account, err := entity.NewAccount(client)
	if err != nil {
		return nil, err
	}

	if err := uc.accountGateway.Save(account); err != nil {
		return nil, err
	}

	return &CreateAccountOutput{
		ID: account.ID,
	}, nil
}
