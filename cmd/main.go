package main

import (
	"context"
	"github.com/BurntSushi/toml"
	"github.com/go-pg/pg/v10"
	"log"
	"newsportal/pkg/app"
	"os"
)

type Config struct {
	Database *pg.Options
}

func main() {
	// read TOML file
	tomlData, err := os.ReadFile("./cfg/local.toml")
	if err != nil {
		log.Fatalf("Error reading local.toml file: %v", err)
	}

	// decode TOML file to struct Config
	var cfg Config
	_, err = toml.Decode(string(tomlData), &cfg)
	if err != nil {
		log.Fatalf("Error decoding TOML data: %v", err)
	}

	// connect db
	opts := &pg.Options{
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		Addr:     cfg.Database.Addr,
		Database: cfg.Database.Database,
	}
	dbc := pg.Connect(opts)
	err = dbc.Ping(context.Background())
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}

	// create app
	a, err := app.New(dbc)
	if err != nil {
		log.Fatalf("Application creation error: %v", err)
	}

	// run application
	err = a.Run()
	if err != nil {
		log.Fatalf("Application run error: %v", err)
	}
}
