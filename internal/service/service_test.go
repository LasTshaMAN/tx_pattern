package service

import (
	"context"
	"testing"
)

func TestService1_Method1(t *testing.T) {
	t.Run("transactional logic should still be easily testable", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			txMock := newTxMock()
			repo1Mock := newRepo1Mock()
			repo2Mock := newRepo2Mock()

			srv := NewService1(txMock, repo1Mock, repo2Mock)

			err := srv.Method1(context.Background())
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !repo2Mock.m2Called {
				t.Fatalf("Method2 on repo2Mock should've been called")
			}
			if !repo1Mock.m1Called {
				t.Fatalf("Method1 on repo1Mock should've been called")
			}
			if !repo2Mock.m1Called {
				t.Fatalf("Method1 on repo2Mock should've been called")
			}
			if !repo1Mock.m2Called {
				t.Fatalf("Method2 on repo1Mock should've been called")
			}
		})
		t.Run("failure", func(t *testing.T) {
			txMock := newTxMock()
			repo1Mock := newRepo1ErrMock()
			repo2Mock := newRepo2Mock()

			srv := NewService1(txMock, repo1Mock, repo2Mock)

			err := srv.Method1(context.Background())
			if err == nil {
				t.Fatalf("expected error, but got: %v", err)
			}
			if !repo2Mock.m2Called {
				t.Fatalf("Method2 on repo2Mock should've been called")
			}
			if !repo1Mock.m1Called {
				t.Fatalf("Method1 on repo1Mock should've been called")
			}
			if repo2Mock.m1Called {
				t.Fatalf("Method1 on repo2Mock shouldn't have been called")
			}
			if repo1Mock.m2Called {
				t.Fatalf("Method2 on repo1Mock shouldn't have been called")
			}
		})
	})
}
