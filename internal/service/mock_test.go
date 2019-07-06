package service

import (
	"context"

	"github.com/pkg/errors"

	domain "github.com/LasTshaMAN/tx_pattern"
)

type txMock struct {
}

func newTxMock() *txMock {
	return &txMock{}
}

func (t *txMock) Wrap(ctx context.Context, fn func(ctx context.Context, tx domain.TxSQL) error) error {
	return fn(ctx, nil)
}

type repo1Mock struct {
	m1Called bool
	m2Called bool
}

func newRepo1Mock() *repo1Mock {
	return &repo1Mock{}
}

func (t *repo1Mock) Method1(ctx context.Context, tx domain.TxSQL) error {
	t.m1Called = true
	return nil
}

func (t *repo1Mock) Method2(ctx context.Context, tx domain.TxSQL) error {
	t.m2Called = true
	return nil
}

type repo2Mock struct {
	m1Called bool
	m2Called bool
}

func newRepo2Mock() *repo2Mock {
	return &repo2Mock{}
}

func (t *repo2Mock) Method1(ctx context.Context, tx domain.TxSQL) error {
	t.m1Called = true
	return nil
}

func (t *repo2Mock) Method2(ctx context.Context) error {
	t.m2Called = true
	return nil
}

type repo1ErrMock struct {
	m1Called bool
	m2Called bool
}

func newRepo1ErrMock() *repo1ErrMock {
	return &repo1ErrMock{}
}

func (t *repo1ErrMock) Method1(ctx context.Context, tx domain.TxSQL) error {
	t.m1Called = true
	return errors.New("unexpected error")
}

func (t *repo1ErrMock) Method2(ctx context.Context, tx domain.TxSQL) error {
	t.m2Called = true
	return nil
}
