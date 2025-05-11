package database

import (
	"database/sql"
	"testing"

	"github.com/CaioAugustoo/wallet-core/internal/entity"
	"github.com/stretchr/testify/suite"

	_ "github.com/mattn/go-sqlite3"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDB *AccountDB
	client    entity.Client
}

func (s *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Require().NoError(err)
	s.db = db

	db.Exec("CREATE TABLE clients (id TEXT PRIMARY KEY, name TEXT, email TEXT, created_at DATE)")
	db.Exec("CREATE TABLE accounts (id TEXT PRIMARY KEY, client_id TEXT, balance FLOAT, created_at DATE)")

	c, _ := entity.NewClient("John Doe", "john@doe.com")
	s.client = *c

	s.accountDB = NewAccountDB(s.db)
}

func (s *AccountDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (s *AccountDBTestSuite) TestSave() {
	account, _ := entity.NewAccount(&s.client)

	err := s.accountDB.Save(account)

	s.Require().NoError(err)
}

func (s *AccountDBTestSuite) TestGet() {
	s.db.Exec("INSERT INTO clients (id, name, email, created_at) VALUES ($1, $2, $3, $4)", s.client.ID, s.client.Name, s.client.Email, s.client.CreatedAt)

	account, _ := entity.NewAccount(&s.client)
	err := s.accountDB.Save(account)
	s.Require().NoError(err)

	result, err := s.accountDB.FindById(account.ID)

	s.Require().NoError(err)
	s.Require().NotNil(result)
	s.Require().Equal(account.ID, result.ID)
	s.Require().Equal(account.Client.ID, result.Client.ID)
	s.Require().Equal(account.Balance, result.Balance)
}
