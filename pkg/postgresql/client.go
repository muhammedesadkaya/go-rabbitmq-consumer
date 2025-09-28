package postgresql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gitlab.com/go-rabbitmq-consumer-app/pkg/config"
)

type PostgreSQLClient interface {
	OpenConnection() (*sql.DB, error)
	CloseConnection(db *sql.DB) error
}

type Client struct {
	config     config.ApplicationConfig
	schemaName string
}

func (self *Client) OpenConnection() (db *sql.DB, err error) {

	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", self.config.PostgreSQL.Host, self.config.PostgreSQL.Port, self.config.PostgreSQL.Username, self.config.PostgreSQL.Password, self.config.PostgreSQL.DbName)

	if db, err = sql.Open("postgres", conn); err != nil {
		return nil, err
	}

	return db, nil
}

func (self *Client) CloseConnection(db *sql.DB) error {
	return db.Close()
}

func NewClient(config config.ApplicationConfig, schemaName string) PostgreSQLClient {
	return &Client{config: config, schemaName: schemaName}
}
