package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

type App struct {
	Router     *mux.Router
	Middleware *Middleware
	Config     *Env
}

type shortenReq struct {
	URL                 string `json:"url" validate:"nonzero"`
	ExpirationInMinutes int64  `json:"expiration_in_minutes" validate:"min=0"`
}

type shortLinkResp struct {
	ShortLink string `json:"shortLink"`
}

// initialization of app
func (a *App) Initialize(e *Env) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	a.Router = mux.NewRouter()
	a.Config = e
	a.Middleware = &Middleware{}

	a.InitializeRouter()
}

func (a *App) InitializeRouter() {
	commonHandler := alice.New(a.Middleware.LoggingHandler, a.Middleware.RecoverHandler)

	a.Router.Handle("/api/shorten", commonHandler.ThenFunc(a.createShortLink)).Methods("POST")
	a.Router.Handle("/api/info", commonHandler.ThenFunc(a.getShortLinkInfo)).Methods("GET")
	a.Router.Handle("/{shortlink:[1-9]+}", commonHandler.ThenFunc(a.redirect)).Methods("GET")

	// a.Router.HandleFunc("/api/shorten", a.createShortLink).Methods("POST")
	// a.Router.HandleFunc("/api/info", a.getShortLinkInfo).Methods("GET")
	// a.Router.HandleFunc("/{shortlink:[1-9]+}", a.redirect).Methods("GET")
}

// Run starts listen and server
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
