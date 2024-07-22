package app

import (
	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
	zm "github.com/vmkteam/zenrpc-middleware"
	"github.com/vmkteam/zenrpc/v2"
	"net/http"
	"newsportal/pkg/db"
	"newsportal/pkg/newsportal"
	"newsportal/pkg/rest"
	"newsportal/pkg/rpc"
)

type App struct {
	nr  db.NewsRepo
	nm  *newsportal.Manager
	ss  *rest.NewsService
	nss *rpc.NewsService
	e   *echo.Echo
	rpc zenrpc.Server
}

func New(dbo *pg.DB) *App {
	a := &App{
		nr: db.NewNewsRepo(dbo),
		e:  echo.New(),
	}
	a.nm = newsportal.NewManager(a.nr)
	a.ss = rest.NewNewsService(a.nm)
	a.nss = rpc.NewNewsService(a.nm)
	a.rpc = rpc.New(a.nm)

	return a
}

func (a *App) Run() error {
	// register handlers
	a.registerHandlers()
	// run HTTP rest
	return a.e.Start(":8080")
}

func (a *App) registerHandlers() {
	a.e.Any("/v1/rpc/", zm.EchoHandler(zm.XRequestID(a.rpc)))
	a.e.Any("/v1/rpc/doc/", echo.WrapHandler(http.HandlerFunc(zenrpc.SMDBoxHandler)))
	a.e.GET("/news", a.ss.NewsWithFilters)
	//a.e.GET("/news/count", a.ss.NewsCountWithFilters)
	a.e.GET("/news/:id", a.ss.NewsByID)
	a.e.GET("/categories", a.ss.Categories)
	a.e.GET("/tags", a.ss.Tags)

}
