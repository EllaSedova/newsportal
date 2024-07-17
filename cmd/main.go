package main

import (
	"context"
	"github.com/go-pg/pg/v10"
	"log"
	"newsportal/pkg/app"
)

func main() {
	//todo add config

	// connect db
	opts := &pg.Options{
		User:     "postgres",
		Password: "postgres",
		Addr:     "localhost:5432",
		Database: "newsportal",
	}
	dbc := pg.Connect(opts)
	err := dbc.Ping(context.Background())
	if err != nil {
		log.Fatalf("DB error: %v", err)
	}
	// create app
	a, err := app.New(dbc)
	if err != nil {
		return
	}

	// run application
	err = a.Run()
	if err != nil {
		log.Fatalf("Application run error: %v", err)
	}

	// запуск сервера
}
