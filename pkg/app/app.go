package app

import (
	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
	"newsportal/pkg/db"
	"newsportal/pkg/newsportal"
	"newsportal/pkg/server"
)

type App struct {
	nr db.NewsRepo
	nm *newsportal.Manager
	ss *server.ServerService
	e  *echo.Echo
}

func New(dbo *pg.DB) *App {
	a := &App{
		nr: db.NewNewsRepo(dbo),
		e:  echo.New(),
	}
	a.nm = newsportal.NewManager(a.nr)
	a.ss = server.NewServerService(a.nm)
	return a
}

func (a *App) Run() error {
	// register handlers
	a.registerHandlers()
	// run HTTP server
	return a.e.Start(":8080")
}

func (a *App) registerHandlers() {
	a.e.GET("/news", a.ss.NewsWithFilters)
	a.e.GET("/news/:id", a.ss.NewsByID)
	a.e.GET("/categories", a.ss.Categories)
	a.e.GET("/tags", a.ss.Tags)

}
