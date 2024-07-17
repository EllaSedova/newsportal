package db

import (
	"github.com/go-pg/pg/v10"
	"testing"
)

func setupTestDB() *pg.DB {
	opts := &pg.Options{
		User:     "postgres",
		Password: "postgres",
		Addr:     "localhost:5432",
		Database: "newsportal",
	}
	return pg.Connect(opts)
}

func TestAddAndReadNews(t *testing.T) {
	// подключение к тестовой базе данных
	testDB := setupTestDB()
	defer testDB.Close()
	//dbRepo := NewNewsRepo(testDB)

}
