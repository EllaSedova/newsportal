package app

import (
	"github.com/go-pg/pg/v10"
	"newsportal/pkg/db"
	"newsportal/pkg/newsportal"
)

type App struct {
	nr db.NewsRepo
	nm *newsportal.Manager
}

func New(dbo *pg.DB) (*App, error) {
	a := &App{
		nr: db.NewNewsRepo(dbo),
	}
	a.nm = newsportal.NewManager(a.nr)
	return a, nil
}

func (a *App) Run() error {
	// register handlers
	a.registerHandlers()

	// run HTTP server
	a.runHTTPServer("localhost", 8080)

	return nil
}

func (a *App) registerHandlers() {

}

func (a *App) runHTTPServer(host string, port int) {

}
