package app

import (
	//"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/itsnitigya/go-store/app/handler"
)

type App struct {
	Router *mux.Router
}

// Initialize initializes the app with predefined configuration
func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {
	a.Post("/set", a.handleRequest(handler.CreatePair))
	a.Get("/get/{key}", a.handleRequest(handler.GetValue))
	a.Get("/searchPrefix/{prefix}", a.handleRequest(handler.SearchPrefixValue))
	a.Get("/searchSuffix/{suffix}", a.handleRequest(handler.SearchSuffixValue))
}

func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

type RequestHandlerFunction func(w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}
}
