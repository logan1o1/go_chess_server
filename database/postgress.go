package database

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Client struct {
	db *sqlx.DB
}

func (client *Client) Execute(query string, args ...interface{}) (sql.Result, error) {
	return client.db.Exec(query, args...)
}

func (client *Client) Get(dest interface{}, query string, args ...interface{}) error {
	return client.db.Get(dest, query, args...)
}

func (client *Client) Select(dest interface{}, query string, args ...interface{}) error {
	return client.db.Select(dest, query, args...)
}

func NewPgClient(pgConnectionString string) (*Client, error) {
	db, err := sqlx.Connect("postgres", pgConnectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Client{db: db}, nil
}
