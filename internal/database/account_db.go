package database

import (
	"database/sql"

	"github.com/CaioAugustoo/wallet-core/internal/entity"
)

type AccountDB struct {
	DB *sql.DB
}

func NewAccountDB(db *sql.DB) *AccountDB {
	return &AccountDB{
		DB: db,
	}
}

func (a *AccountDB) FindById(id string) (*entity.Account, error) {
	var account entity.Account
	var client entity.Client
	account.Client = &client

	err := a.DB.QueryRow("SELECT a.id, a.client_id, a.balance, a.created_at, c.id, c.name, c.email, c.created_at FROM accounts a INNER JOIN clients c ON a.client_id = c.id WHERE a.id = $1", id).Scan(
		&account.ID,
		&account.Client.ID,
		&account.Balance,
		&account.CreatedAt,
		&client.ID,
		&client.Name,
		&client.Email,
		&client.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (a *AccountDB) Save(account *entity.Account) error {
	_, err := a.DB.Exec(
		"INSERT INTO accounts (id, client_id, balance, created_at) VALUES ($1, $2, $3, $4)",
		account.ID, account.Client.ID, account.Balance, account.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
