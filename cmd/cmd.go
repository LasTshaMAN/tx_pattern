package main

import (
	"context"

	"github.com/LasTshaMAN/tx_pattern/internal/pg"
	"github.com/LasTshaMAN/tx_pattern/internal/service"
)

func main() {
	client, err := pg.NewClient("")
	if err != nil {
		panic(err)
	}
	tx := pg.NewTx(client)
	repo1 := pg.NewRepo1()
	repo2 := pg.NewRepo2()
	srv := service.NewService1(tx, repo1, repo2)

	err = srv.Method1(context.Background())
	if err != nil {
		panic(err)
	}
}
