package app

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/gorilla/mux"
	"net/http"
	"newsportal/pkg/db"
	"newsportal/pkg/newsportal"
	"newsportal/pkg/server"
)

type App struct {
	nr     db.NewsRepo
	nm     *newsportal.Manager
	ss     *server.ServerService
	router *mux.Router
}

func New(dbo *pg.DB) (*App, error) {
	a := &App{
		nr:     db.NewNewsRepo(dbo),
		router: mux.NewRouter(),
	}
	a.nm = newsportal.NewManager(a.nr)
	a.ss = server.NewServerService(a.nm)
	return a, nil
}

func (a *App) Run() error {
	// register handlers
	a.registerHandlers()

	// run HTTP server
	err := a.runHTTPServer(":8080")
	if err != nil {
		return fmt.Errorf("[Run]: %s\n", err)
	}

	return nil
}

func (a *App) registerHandlers() {
	a.router.HandleFunc("/news/{id:[0-9]+}", a.ss.NewsByID).Methods("GET")
	a.router.HandleFunc("/news/filtered", a.ss.NewsWithFilters).Methods("GET")
	a.router.HandleFunc("/categories", a.ss.Categories).Methods("GET")
	a.router.HandleFunc("/tags", a.ss.Tags).Methods("GET")

}

func (a *App) runHTTPServer(port string) error {
	return http.ListenAndServe(port, a.router)

}
