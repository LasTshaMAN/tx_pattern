package service

import (
	"context"

	domain "github.com/LasTshaMAN/tx_pattern"
)

// service1 implements Service1 from the domain package.
type service1 struct {
	tx    domain.Tx
	repo1 domain.Repo1
	repo2 domain.Repo2
}

// NewService1 returns new Service1 object.
func NewService1(tx domain.Tx, repo1 domain.Repo1, repo2 domain.Repo2) domain.Service1 {
	return &service1{
		tx:    tx,
		repo1: repo1,
		repo2: repo2,
	}
}

func (s *service1) Method1(ctx context.Context) error {
	// Throw some business logic into the mix.
	err := s.repo2.Method2(ctx)
	if err != nil {
		return err
	}

	// ...

	// Here goes our transaction.
	err = s.tx.Wrap(ctx, func(ctx context.Context, tx domain.TxSQL) error {
		err := s.repo1.Method1(ctx, tx)
		if err != nil {
			return err
		}

		err = s.repo2.Method1(ctx, tx)
		if err != nil {
			return err
		}

		err = s.repo1.Method2(ctx, tx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	// Send some notification.
	// ...

	return nil
}
