package pg

import (
	"context"

	domain "github.com/LasTshaMAN/tx_pattern"
)

// repo1 implements Repo1 from the domain package.
type repo1 struct {
}

// NewRepo1 returns new Repo1 object.
func NewRepo1() domain.Repo1 {
	return &repo1{}
}

// Method1 requires transaction.
func (r *repo1) Method1(ctx context.Context, tx domain.TxSQL) error {
	_, err := tx.ExecContext(ctx, "query")
	if err != nil {
		return err
	}

	// ...

	return nil
}

// Method2 requires transaction.
func (r *repo1) Method2(ctx context.Context, tx domain.TxSQL) error {
	// ...

	return nil
}
