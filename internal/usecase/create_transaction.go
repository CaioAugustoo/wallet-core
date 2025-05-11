package usecase

import (
	"github.com/CaioAugustoo/wallet-core/internal/entity"
	"github.com/CaioAugustoo/wallet-core/internal/gateway"
)

type createTransaction struct {
	accountGateway     gateway.Account
	transactionGateway gateway.Transaction
}

type CreateTransactionInput struct {
	AccountIDFrom string
	AccountIDTo   string
	Amount        float64
}

type CreateTransactionOutput struct {
	TransactionID string
}

type CreateTransaction interface {
	Execute(input CreateTransactionInput) (*CreateTransactionOutput, error)
}

func NewCreateTransaction(
	accountGateway gateway.Account,
	transactionGateway gateway.Transaction,
) CreateTransaction {
	return &createTransaction{
		accountGateway:     accountGateway,
		transactionGateway: transactionGateway,
	}
}

func (uc *createTransaction) Execute(
	input CreateTransactionInput,
) (*CreateTransactionOutput, error) {
	accountFrom, err := uc.accountGateway.FindById(input.AccountIDFrom)
	if err != nil {
		return nil, err
	}

	accountTo, err := uc.accountGateway.FindById(input.AccountIDTo)
	if err != nil {
		return nil, err
	}

	transaction, err := entity.NewTransaction(
		accountFrom,
		accountTo,
		input.Amount,
	)
	if err != nil {
		return nil, err
	}

	if err := uc.transactionGateway.Create(transaction); err != nil {
		return nil, err
	}

	return &CreateTransactionOutput{
		TransactionID: transaction.ID,
	}, nil
}
