package db

import (
	"context"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/go-pg/pg/v10"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
	"time"
)

func ptrs(r string) *string { return &r }

type Config struct {
	Database *pg.Options
}

var db *pg.DB
var realNews News

func TestMain(m *testing.M) {
	tomlData, err := os.ReadFile("../../cfg/local.toml")
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

	db = pg.Connect(opts)
	err = db.Ping(context.Background())
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}
	realNews = News{
		Title:       "Новость1",
		CategoryID:  1,
		Foreword:    "Преамбула",
		Content:     ptrs("Контент"),
		TagIDs:      []int{1, 2},
		Author:      "Автор",
		PublishedAt: time.Date(2024, time.July, 17, 18, 25, 28, 10745000, time.Local),
		StatusID:    1,
	}
	os.Exit(m.Run())
}

func TestGetNewsByID(t *testing.T) {
	nr := NewNewsRepo(db)
	fmt.Println(realNews)

	// add news
	_, err := db.Model(&realNews).Insert()

	// get wrong news by id
	wrongNews, err := nr.NewsByID(6)
	assert.Nil(t, wrongNews)
	// get true news by id
	actualNews, err := nr.NewsByID(realNews.ID)
	assert.NoError(t, err)
	assert.Equal(t, &realNews, actualNews)

	// delete news by id
	_, err = db.Model(&realNews).WherePK().Delete()
	assert.NoError(t, err)
}

func TestGetNewsByTagID(t *testing.T) {
	nr := NewNewsRepo(db)

	// get news by tag
	actualNews, err := nr.NewsByTagID(1)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(actualNews), "there is no news with this tagId")
	fmt.Println(actualNews)
}

func TestGetNewsByCategoryID(t *testing.T) {
	nr := NewNewsRepo(db)

	// get news by category
	actualNews, err := nr.NewsByCategoryID(1)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(actualNews), "there is no news with this categoryId")
	fmt.Println(actualNews)
}

func TestGetNewsByCategoryIDWithPagination(t *testing.T) {
	nr := NewNewsRepo(db)

	// get news by category with pagination
	actualNews, err := nr.NewsByCategoryIDWithPagination(1, 3, 3)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(actualNews))
	fmt.Println(actualNews)
}

func TestGetNewsByTagIDWithPagination(t *testing.T) {
	nr := NewNewsRepo(db)

	// get news by category with pagination
	actualNews, err := nr.NewsByTagIDWithPagination(1, 3, 3)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(actualNews))
	fmt.Println(actualNews)
}
