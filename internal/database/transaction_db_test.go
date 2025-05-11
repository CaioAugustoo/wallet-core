package database

import (
	"database/sql"
	"testing"

	"github.com/CaioAugustoo/wallet-core/internal/entity"
	"github.com/stretchr/testify/suite"

	_ "github.com/mattn/go-sqlite3"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	client        entity.Client
	client2       entity.Client
	accountFrom   entity.Account
	accountTo     entity.Account
	transactionDB *TransactionDB
}

func (s *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Require().NoError(err)
	s.db = db

	db.Exec("CREATE TABLE clients (id TEXT PRIMARY KEY, name TEXT, email TEXT, created_at DATE)")
	db.Exec("CREATE TABLE accounts (id TEXT PRIMARY KEY, client_id TEXT, balance FLOAT, created_at DATE)")
	db.Exec("CREATE TABLE transactions (id TEXT PRIMARY KEY, account_from_id TEXT, account_to_id TEXT, amount FLOAT, created_at DATE)")

	client, _ := entity.NewClient("John Doe", "john@doe.com")
	s.client = *client

	client2, _ := entity.NewClient("John Doe2", "john2@doe.com")
	s.client2 = *client2

	accountFrom, _ := entity.NewAccount(&s.client)
	s.accountFrom = *accountFrom
	s.accountFrom.Balance = 100.0

	accountTo, _ := entity.NewAccount(&s.client2)
	s.accountTo = *accountTo
	s.accountTo.Balance = 100.0

	s.transactionDB = NewTransactionDB(s.db)
}

func (s *TransactionDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE transactions")
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (s *TransactionDBTestSuite) TestCreate() {
	transaction, err := entity.NewTransaction(&s.accountFrom, &s.accountTo, 100.0)
	s.Require().NoError(err)
	err = s.transactionDB.Create(transaction)

	s.Require().NoError(err)
}
