package server

import (
	"context"
	"github.com/go-pg/pg/v10"
	"log"
	"newsportal/pkg/db"
	"newsportal/pkg/newsportal"
	"os"
	"testing"
)

var dbc *pg.DB
var nr db.NewsRepo
var nm *newsportal.Manager
var ss *ServerService

func TestMain(m *testing.M) {
	opts := &pg.Options{
		User:     "postgres",
		Password: "postgres",
		Addr:     "localhost:5432",
		Database: "newsportal",
	}

	dbc = pg.Connect(opts)
	err := dbc.Ping(context.Background())
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}

	nr = db.NewNewsRepo(dbc)
	nm = newsportal.NewManager(nr)
	ss = NewServerService(nm)

	os.Exit(m.Run())
}
