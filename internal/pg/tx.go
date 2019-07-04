package pg

import (
	"context"
	"database/sql"
	"fmt"
	"runtime/debug"

	"github.com/pkg/errors"

	domain "github.com/LasTshaMAN/tx_pattern"
)

// Tx manages SQL transactions.
type tx struct {
	client *Client
}

// NewTx returns new Tx object.
func NewTx(client *Client) domain.Tx {
	return &tx{
		client: client,
	}
}

// Wrap creates sql.Tx transaction tx, calls fn with this transaction tx and depending on the result from fn
// Wrap performs commit / rollback for the transaction tx.
func (t *tx) Wrap(ctx context.Context, fn func(ctx context.Context, tx domain.TxSQL) error) error {
	tx, err := t.client.db.BeginTx(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "failed to open new transaction")
	}

	defer func() {
		if rec := recover(); rec != nil {
			err = rollbackTx(tx, fmt.Errorf("panic during transaction execution: %v\n%s", rec, debug.Stack()))
		}
	}()

	if err := fn(ctx, tx); err != nil {
		return rollbackTx(tx, err)
	}

	return tx.Commit()
}

func rollbackTx(tx *sql.Tx, cause error) error {
	if rollErr := tx.Rollback(); rollErr != nil {
		return errors.Wrap(cause, fmt.Sprintf("rollback failed: %v", rollErr))
	}
	return cause
}
