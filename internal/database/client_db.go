package database

import (
	"database/sql"

	"github.com/CaioAugustoo/wallet-core/internal/entity"
)

type ClientDB struct {
	DB *sql.DB
}

func NewClientDB(db *sql.DB) *ClientDB {
	return &ClientDB{
		DB: db,
	}
}

func (c *ClientDB) Get(id string) (*entity.Client, error) {
	client := &entity.Client{}
	err := c.DB.QueryRow("SELECT id, name, email, created_at FROM clients WHERE id = $1", id).Scan(
		&client.ID,
		&client.Name,
		&client.Email,
		&client.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (c *ClientDB) Save(client *entity.Client) error {
	_, err := c.DB.Exec(
		"INSERT INTO clients (id, name, email, created_at) VALUES ($1, $2, $3, $4)",
		client.ID, client.Name, client.Email, client.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
