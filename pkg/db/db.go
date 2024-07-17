package db

import (
	"github.com/go-pg/pg/v10"
)

// Connect подключение к базе данных
func Connect() *pg.DB {
	//todo чтение данных о бд из конфига
	//подключение к базе данных
	opts := &pg.Options{
		User:     "postgres",
		Password: "postgres",
		Addr:     "localhost:5432",
		Database: "newsportal",
	}
	db := pg.Connect(opts)
	return db
}
