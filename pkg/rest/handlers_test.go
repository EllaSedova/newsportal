package rest

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"newsportal/pkg/db"
	"newsportal/pkg/newsportal"
	"os"
	"testing"
)

var dbc *pg.DB
var nr db.NewsRepo
var nm *newsportal.Manager
var ss *NewsService
var e *echo.Echo

const trueNews = `{"newsId":15,"title":"AНовость5","categoryId":{"categoryId":1,"title":"рр","orderNumber":null,"alias":"к"},"foreword":"Преамбула","content":"Контент","tags":[{"tagId":1,"title":"заголовок1"},{"tagId":2,"title":"заголовок2"},{"tagId":3,"title":"заголовок3"}],"author":"Автор","publishedAt":"2024-07-17T18:25:28.010745+03:00"}
`
const trueNewsList = `[{"newsId":15,"title":"AНовость5","categoryId":{"categoryId":1,"title":"рр","orderNumber":null,"alias":"к"},"foreword":"Преамбула","tags":[{"tagId":1,"title":"заголовок1"},{"tagId":2,"title":"заголовок2"},{"tagId":3,"title":"заголовок3"}],"author":"Автор","publishedAt":"2024-07-17T18:25:28.010745+03:00"}]
`
const wrongNews = `null
`

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

	os.Exit(m.Run())
}
func TestNewsById(t *testing.T) {

	// регистрация маршрута
	e.GET("/news/:id", func(c echo.Context) error {
		return ss.NewsByID(c) // Замените на вашу функцию-обработчик
	})
	// создание нового HTTP запроса
	reqTrue := httptest.NewRequest(http.MethodGet, "/news/15", nil)
	recTrue := httptest.NewRecorder()
	// создание нового HTTP запроса
	reqWrong := httptest.NewRequest(http.MethodGet, "/news/0", nil)
	recWrong := httptest.NewRecorder()
	// выполнение запроса
	e.ServeHTTP(recTrue, reqTrue)
	e.ServeHTTP(recWrong, reqWrong)

	// проверки
	assert.Equal(t, http.StatusOK, recTrue.Code)
	assert.Equal(t, trueNews, recTrue.Body.String())
	assert.Equal(t, http.StatusOK, recWrong.Code)
	assert.Equal(t, wrongNews, recWrong.Body.String())

}

func TestNewsWithFilters(t *testing.T) {
	type args struct {
		c echo.Context
	}
	// описываем тестовые случаи
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
		want    string
	}{
		{
			name: "valid filters",
			args: args{
				c: validEchoContext(),
			},
			want:    trueNewsList,
			wantErr: assert.NoError,
		},
		{
			name: "invalid filters",
			args: args{
				c: invalidEchoContext(),
			},
			want:    wrongNews,
			wantErr: assert.NoError,
		},
	}
	// тестируем
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			tt.args.c.Response().Writer = rec

			err := ss.NewsWithFilters(tt.args.c)
			if !tt.wantErr(t, err, fmt.Sprintf("NewsWithFilters(%v)", tt.args.c)) {
				return
			}
			// вывод результата
			got := rec.Body.String()
			assert.Equal(t, tt.want, got)
		})
	}
}

// создание echo.Context с валидными фильтрами
func validEchoContext() echo.Context {
	req := httptest.NewRequest(http.MethodGet, "/news", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// добавление параметров в запрос
	q := req.URL.Query()
	q.Add("categoryID", "1")
	q.Add("tagID", "3")
	req.URL.RawQuery = q.Encode()

	return c
}

// создание echo.Context с невалидными фильтрами
func invalidEchoContext() echo.Context {
	req := httptest.NewRequest(http.MethodGet, "/news", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// добавление параметров в запрос
	q := req.URL.Query()
	q.Add("categoryID", "13")
	req.URL.RawQuery = q.Encode()

	return c
}
