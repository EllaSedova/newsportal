package newsportal

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/stretchr/testify/assert"
	"log"
	"newsportal/pkg/db"
	"os"
	"testing"
	"time"
)

func ptrs(r string) *string { return &r }

var dbc *pg.DB
var nr db.NewsRepo
var nm *Manager
var realNews News

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
	nm = NewManager(nr)

	realNews = News{
		News: &db.News{ID: 11,
			Title:       "Новость1",
			CategoryID:  1,
			Foreword:    "Преамбула",
			Content:     ptrs("Контент"),
			TagIDs:      []int{1, 2},
			Author:      "Автор",
			PublishedAt: time.Date(2024, time.July, 17, 18, 25, 28, 10745000, time.Local),
			StatusID:    1},
		Category: &Category{
			ID:          1,
			Title:       "pp",
			OrderNumber: nil,
			Alias:       "к",
			StatusID:    1,
		},
	}
	os.Exit(m.Run())
}

func TestGetNewsByID(t *testing.T) {
	// get wrong news by id
	//wrongNews, err := nm.NewsByID(6)
	//assert.Nil(t, wrongNews)
	// get true news by id
	actualNews, err := nm.NewsByID(11)
	fmt.Println(actualNews.Category)
	fmt.Println(actualNews.Title)
	assert.NoError(t, err)
	assert.Equal(t, &realNews, actualNews)
}
