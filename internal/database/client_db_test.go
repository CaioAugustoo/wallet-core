package database

import (
	"database/sql"
	"testing"

	"github.com/CaioAugustoo/wallet-core/internal/entity"
	"github.com/stretchr/testify/suite"

	_ "github.com/mattn/go-sqlite3"
)

type ClientDBTestSuite struct {
	suite.Suite
	db       *sql.DB
	clientDB *ClientDB
}

func (s *ClientDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Require().NoError(err)
	s.db = db

	db.Exec("CREATE TABLE clients (id TEXT PRIMARY KEY, name TEXT, email TEXT, created_at DATE)")
	s.clientDB = NewClientDB(s.db)
}

func (s *ClientDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
}

func TestClientDBTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDBTestSuite))
}

func (s *ClientDBTestSuite) TestSave() {
	client, _ := entity.NewClient("John Doe", "john@doe.com")

	err := s.clientDB.Save(client)

	s.Require().NoError(err)
}

func (s *ClientDBTestSuite) TestGet() {
	client, _ := entity.NewClient("John Doe", "john@doe.com")
	s.clientDB.Save(client)

	result, err := s.clientDB.Get(client.ID)

	s.Require().NoError(err)
	s.Require().NotNil(result)
	s.Require().Equal(client.ID, result.ID)
	s.Require().Equal(client.Name, result.Name)
	s.Require().Equal(client.Email, result.Email)
}
