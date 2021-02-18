package main

import (
	"log"
	"net/http"
	"time"
)

type Middleware struct {
}

func (m Middleware) LoggingHandler(handler http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		handler.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("\n[%s] %q %v \n\n", r.Method, r.URL.String(), t2.Sub(t1))
	}
	return http.HandlerFunc(fn)
}

func (m Middleware) RecoverHandler(handler http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("\nrecover from panic %v \n\n", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()
		handler.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
