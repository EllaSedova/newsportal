package rpc

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"log"
	"newsportal/pkg/db"
	"newsportal/pkg/newsportal"
	"os"
	"testing"
	"time"
)

var dbc *pg.DB
var nr db.NewsRepo
var nm *newsportal.Manager
var ss *NewsService
var e *echo.Echo
var realNewsSummary NewsSummary
var wrongNewsSummary NewsSummary
var realNews News
var wrongNews News

func ptrs(r string) *string { return &r }
func ptri(r int) *int       { return &r }

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
	ss = NewNewsService(nm)
	e = echo.New()
	realNews = News{
		ID:    15,
		Title: "AНовость5",
		Category: Category{
			ID:          1,
			Title:       "рр",
			OrderNumber: nil,
			Alias:       "к",
		},
		Foreword: "Преамбула",
		Content:  ptrs("Контент"),
		Tags: []Tag{
			{
				ID:    1,
				Title: "заголовок1",
			},
			{
				ID:    2,
				Title: "заголовок2",
			},
			{
				ID:    3,
				Title: "заголовок3",
			},
		},
		Author:      "Автор",
		PublishedAt: time.Date(2024, time.July, 17, 18, 25, 28, 10745000, time.Local),
	}
	realNewsSummary = NewsSummary{
		ID:    15,
		Title: "AНовость5",
		Category: Category{
			ID:          1,
			Title:       "рр",
			OrderNumber: nil,
			Alias:       "к",
		},
		Foreword: "Преамбула",
		Tags: []Tag{
			{
				ID:    1,
				Title: "заголовок1",
			},
			{
				ID:    2,
				Title: "заголовок2",
			},
			{
				ID:    3,
				Title: "заголовок3",
			},
		},
		Author:      "Автор",
		PublishedAt: time.Date(2024, time.July, 17, 18, 25, 28, 10745000, time.Local),
	}
	os.Exit(m.Run())
}
func TestNewsById(t *testing.T) {
	news, err := ss.NewsByID(15)
	// проверки
	assert.NoError(t, err)
	assert.Equal(t, realNews, *news)
}
func TestFill(t *testing.T) {

	type args struct {
		CategoryID *int
		TagID      *int
		Page       *int
		PageSize   *int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "valid filters",
			args: args{CategoryID: ptri(1), TagID: ptri(3), Page: ptri(1), PageSize: ptri(2)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trueParams := &FilterParams{
				CategoryID: tt.args.CategoryID,
				TagID:      tt.args.TagID,
				Page:       tt.args.Page,
				PageSize:   tt.args.PageSize,
			}
			p := trueParams
			p.Fill(tt.args.CategoryID, tt.args.TagID, tt.args.Page, tt.args.PageSize)
			assert.Equal(t, trueParams, p)
		})
	}
}

func TestNewsWithFilters(t *testing.T) {
	type args struct {
		categoryID *int
		tagID      *int
		page       *int
		pageSize   *int
	}
	// описываем тестовые случаи
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
		want    []NewsSummary
	}{
		{
			name:    "valid filters",
			args:    args{categoryID: ptri(1), tagID: ptri(3), page: nil, pageSize: nil},
			want:    []NewsSummary{realNewsSummary},
			wantErr: assert.NoError,
		},
		{
			name:    "invalid page",
			args:    args{categoryID: ptri(1), tagID: ptri(1), page: ptri(100), pageSize: ptri(12)},
			want:    nil,
			wantErr: assert.NoError,
		},
		{
			name:    "invalid tagID",
			args:    args{categoryID: ptri(1), tagID: ptri(-1), page: ptri(1), pageSize: ptri(17)},
			want:    nil,
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ss.NewsWithFilters(tt.args.categoryID, tt.args.tagID, tt.args.page, tt.args.pageSize)
			if !tt.wantErr(t, err, fmt.Sprintf("NewsWithFilters(%v, %v, %v, %v)", tt.args.categoryID, tt.args.tagID, tt.args.page, tt.args.pageSize)) {
				return
			}
			log.Println(tt.name)
			assert.Equalf(t, tt.want, got, "NewsWithFilters(%v, %v, %v, %v)", tt.args.categoryID, tt.args.tagID, tt.args.page, tt.args.pageSize)
		})
	}
}
