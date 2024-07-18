package server

import (
	"context"
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
var ss *ServerService

const trueNews = `{"newsId":13,"title":"Новость3","categoryId":{"categoryId":3,"title":"h","orderNumber":1,"alias":"u"},"foreword":"Преамбула","content":"Контент","tags":[{"tagId":1,"title":"заголовок1"},{"tagId":2,"title":"заголовок2"}],"author":"Автор","publishedAt":"2024-07-17T18:25:28.010745+03:00"}
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
	ss = NewServerService(nm)

	os.Exit(m.Run())
}
func TestNewsById(t *testing.T) {
	e := echo.New()
	// регистрация маршрута
	e.GET("/news/:id", func(c echo.Context) error {
		return ss.NewsByID(c) // Замените на вашу функцию-обработчик
	})
	// создание нового HTTP запроса
	req := httptest.NewRequest(http.MethodGet, "/news/13", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	// установка параметров
	c.SetParamNames("id")
	c.SetParamValues("13")

	// выполнение запроса
	e.ServeHTTP(rec, req)

	// проверки
	assert.Equal(t, http.StatusOK, rec.Code)

	assert.Equal(t, trueNews, rec.Body.String())
}
