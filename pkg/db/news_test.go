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
var nr NewsRepo

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
		ID:          11,
		Title:       "Новость1",
		CategoryID:  1,
		Foreword:    "Преамбула",
		Content:     ptrs("Контент"),
		TagIDs:      []int{1, 2},
		Author:      "Автор",
		PublishedAt: time.Date(2024, time.July, 17, 18, 25, 28, 10745000, time.Local),
		StatusID:    1,
		Category: &Category{
			ID:          1,
			Title:       "рр",
			OrderNumber: nil,
			Alias:       "к",
			StatusID:    1,
		},
	}

	nr = NewNewsRepo(db)
	os.Exit(m.Run())
}

func TestNewsByID(t *testing.T) {
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

func TestNewsWithPagination(t *testing.T) {
	qb := &nr.QB
	page := 3
	pageSize := 2
	categoryID := 3
	tagID := 1
	sortTitle := false
	qb.AddFilter(`t."categoryId"`, categoryID)
	qb.AddNewFilter(`ANY (t."tagIds")`, tagID)
	qb.AddSort("t.title", sortTitle)

	// get news by tag
	actualNews, err := nr.NewsWithPagination(page, pageSize, qb)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(actualNews), "there is no news with this tagId")
}

func TestTagsByIDs(t *testing.T) {
	ids := []int{1, 2}
	// get news by tag
	actualTags, err := nr.TagsByIDs(ids)
	log.Println(actualTags)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(actualTags), "there is no news with this tagId")
}

func TestNewsRepo_NewsWithPagination(t *testing.T) {
	type fields struct {
		QB QueryBuilder
	}
	type args struct {
		page     int
		pageSize int
		qb       *QueryBuilder
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []News
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "invalid filters",
			args: args{
				page:     2,
				pageSize: -2,
				qb:       &nr.QB,
			},
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//tt.args.qb.AddFilter(`"categoryID"`, 3)
			got, err := nr.NewsWithPagination(tt.args.page, tt.args.pageSize, tt.args.qb)
			if !tt.wantErr(t, err, fmt.Sprintf("NewsWithPagination(%v, %v, %v)", tt.args.page, tt.args.pageSize, tt.args.qb)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewsWithPagination(%v, %v, %v)", tt.args.page, tt.args.pageSize, tt.args.qb)
		})
	}
}
