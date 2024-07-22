package newsportal

import (
	"context"
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
		News: &db.News{ID: 15,
			Title:       "AНовость5",
			CategoryID:  1,
			Foreword:    "Преамбула",
			Content:     ptrs("Контент"),
			TagIDs:      []int{1, 2, 3},
			Author:      "Автор",
			PublishedAt: time.Date(2024, time.July, 17, 18, 25, 28, 10745000, time.Local),
			StatusID:    1},
		Category: &Category{
			ID:          1,
			Title:       "рр",
			OrderNumber: nil,
			Alias:       "к",
			StatusID:    1,
		},
		Tags: []Tag{
			{
				ID:       1,
				Title:    "заголовок1",
				StatusID: 1,
			},
			{
				ID:       2,
				Title:    "заголовок2",
				StatusID: 1,
			},
			{
				ID:       3,
				Title:    "заголовок3",
				StatusID: 1,
			},
		},
	}
	os.Exit(m.Run())
}

func TestGetNewsByID(t *testing.T) {
	// get wrong news by id
	wrongNews, err := nm.NewsByID(context.Background(), 6)
	assert.Nil(t, wrongNews)

	// get true news by id
	actualNews, err := nm.NewsByID(context.Background(), 15)
	assert.NoError(t, err)
	assert.Equal(t, realNews.Category, actualNews.Category)
	assert.Equal(t, realNews.Tags, actualNews.Tags)
	assert.Equal(t, realNews.PublishedAt, actualNews.PublishedAt)

}

func TestFillTags(t *testing.T) {
	// Инициализация тестовых данных
	newsListTrue := []News{
		{News: &db.News{ID: 1, TagIDs: []int{1, 2}}},
		{News: &db.News{ID: 2, TagIDs: []int{2, 3}}},
	}
	newsListWrong := []News{
		{News: &db.News{ID: 1, TagIDs: []int{}}},
		{News: &db.News{ID: 2, TagIDs: []int{2, 13}}},
	}
	// Вызов метода
	err := nm.FillTags(context.Background(), newsListTrue)
	assert.NoError(t, err)

	// Проверка результатов
	trueTags := map[int][]Tag{
		1: {{ID: 1, Title: "заголовок1", StatusID: 1}, {ID: 2, Title: "заголовок2", StatusID: 1}},
		2: {{ID: 2, Title: "заголовок2", StatusID: 1}, {ID: 3, Title: "заголовок3", StatusID: 1}},
	}
	err = nm.FillTags(context.Background(), newsListWrong)
	assert.NoError(t, err)
	// Проверка результатов
	wrongTags := map[int][]Tag{
		1: nil,
		2: {{ID: 2, Title: "заголовок2", StatusID: 1}},
	}
	for _, news := range newsListTrue {
		assert.Equal(t, trueTags[news.ID], news.Tags)
	}
	for _, news := range newsListWrong {
		assert.Equal(t, wrongTags[news.ID], news.Tags)
	}
}

func TestNews(t *testing.T) {
	categoryID := 1
	tagID := 3
	page := 1
	pageSize := 10

	newNewsList, err := nm.News(context.Background(), &categoryID, &tagID, &page, &pageSize)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, newNewsList)
	assert.Len(t, newNewsList, 1)
	assert.Equal(t, realNews.Tags, newNewsList[0].Tags)
	assert.Equal(t, realNews.Category, newNewsList[0].Category)
	assert.Equal(t, realNews.PublishedAt, newNewsList[0].PublishedAt)
	assert.Equal(t, realNews.ID, newNewsList[0].ID)

	wrongCategoryID := 17
	var emptyTagID int
	newNewsList2, err := nm.News(context.Background(), &wrongCategoryID, &emptyTagID, &page, &pageSize)
	assert.NoError(t, err)
	assert.Nil(t, newNewsList2)
}

func TestTagsByIDs(t *testing.T) {
	ids := []int{1, 2}
	tags, err := nm.TagsByIDs(context.Background(), ids)
	assert.NoError(t, err)
	// Проверка результатов
	trueTags := []Tag{{ID: 1, Title: "заголовок1", StatusID: 1}, {ID: 2, Title: "заголовок2", StatusID: 1}}

	assert.Equal(t, trueTags, tags)

}
