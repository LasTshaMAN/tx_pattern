// Package domain defines the domain model of this service.
package domain

import (
	"context"
	"database/sql"
)

// TxSQL provides methods for reading from / writing to SQL-based storage (e.g. Postgres).
//
// TxSQL mustn't define any methods related to commit / rollback action.
type TxSQL interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	// ...
}

// Tx is responsible for transaction management - creation, commit / rollback action upon finishing executing target function fn.
type Tx interface {
	// Wrap creates and otherwise manages a transaction tx.
	// Wrap passes transaction tx to function fn for fn to be able to leverage this transaction tx.
	//
	// Wrap returns any errors it receives from fn.
	Wrap(ctx context.Context, fn func(ctx context.Context, tx TxSQL) error) error
}

// Service1 is some service in the domain package.
type Service1 interface {
	// Method1 is some method on Service1.
	Method1(ctx context.Context) error
}

// Repo1 is some repository abstraction to store some domain entities in Postgres.
type Repo1 interface {
	// Method1 is some method on Repo1.
	Method1(ctx context.Context, tx TxSQL) error
	// Method2 is some method on Repo1.
	Method2(ctx context.Context, tx TxSQL) error
}

// Repo2 is some repository abstraction to store some domain entities in Postgres.
type Repo2 interface {
	// Method1 is some method on Repo2.
	Method1(ctx context.Context, tx TxSQL) error
	// Method1 is some method on Repo2 that doesn't require to be executed in a transaction.
	Method2(ctx context.Context) error
}
