package mocks

import (
	"database/sql"

	"github.com/logan1o1/go_chess_server/interfaces"
)

type MockDatabase struct {
	ExecuteFunc func(query string, args ...interface{}) (sql.Result, error)
	GetFunc     func(dest interface{}, query string, args ...interface{}) error
	SelectFunc  func(dest interface{}, query string, args ...interface{}) error
}

func (m *MockDatabase) Execute(query string, args ...interface{}) (sql.Result, error) {
	if m.ExecuteFunc != nil {
		return m.ExecuteFunc(query, args...)
	}
	return nil, nil
}

func (m *MockDatabase) Get(dest interface{}, query string, args ...interface{}) error {
	if m.GetFunc != nil {
		return m.GetFunc(dest, query, args...)
	}
	return nil
}

func (m *MockDatabase) Select(dest interface{}, query string, args ...interface{}) error {
	if m.SelectFunc != nil {
		return m.SelectFunc(dest, query, args...)
	}
	return nil
}

var _ interfaces.IDatabase = (*MockDatabase)(nil)
