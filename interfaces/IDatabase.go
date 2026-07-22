package interfaces

import (
	"database/sql"
)

type IDatabase interface {
	Execute(query string, args ...any) (sql.Result, error)
	Get(dest any, query string, args ...any) error
	Select(dest any, query string, args ...any) error
}
