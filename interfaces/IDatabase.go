package interfaces

import (
	"database/sql"
)

type IDatabase interface {
	Execute(query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
}
