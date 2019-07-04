package pg

import (
	"context"

	domain "github.com/LasTshaMAN/tx_pattern"
)

// repo2 implements Repo2 from the domain package.
type repo2 struct {
}

// NewRepo2 returns new Repo2 object.
func NewRepo2() domain.Repo2 {
	return &repo2{}
}

// Method1 requires transaction.
func (r *repo2) Method1(ctx context.Context, tx domain.TxSQL) error {
	res, err := tx.QueryContext(ctx, "")
	if err != nil {
		return err
	}
	defer res.Close()

	// Inspect DB rows here
	// ...

	return nil
}

// Method2 doesn't require transaction.
func (r *repo2) Method2(ctx context.Context) error {
	return nil
}
