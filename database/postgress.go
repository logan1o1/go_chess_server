package database

import (
	"database/sql"
	"fmt"

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

func NewPgClient(db_user, db_password, db_name string) (*Client, error) {
	pgConnectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", db_user, db_password, db_name)

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
