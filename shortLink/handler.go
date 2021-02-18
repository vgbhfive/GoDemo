package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
)

// ShortRequest for request: url, expiration_in_minutes
type ShortRequest struct {
	URL                 string `json:"url" validate:"required"`
	ExpirationInMinutes int64  `json:"expiration_in_minutes" validate:"min=0"`
}

// ShortlinkRequest :shortlink
type ShortlinkRequest struct {
	Shortlink string `json:"shortlink"`
}

func (a *App) createShortLink(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("create short link! \n")
	var req ShortRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Printf("Req: %v\n", req)
		respondWithError(w, StatusError{Code: http.StatusBadRequest, Err: fmt.Errorf("json decode err %v", r.Body)})
		return
	}
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		respondWithError(w, StatusError{Code: http.StatusBadRequest, Err: fmt.Errorf("json validate err %v", req)})
		return
	}

	defer r.Body.Close()
	fmt.Printf("Req: %v\n", req)
	s, err := a.Config.S.Shorten(req.URL, req.ExpirationInMinutes)
	if err != nil {
		respondWithError(w, err)
	} else {
		respondWithJSON(w, http.StatusCreated, ShortlinkRequest{Shortlink: s})
	}

}

func (a *App) getShortLinkInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("get short link info! \n")
	url := r.URL.Query()
	s := url.Get("shortlink")

	d, err := a.Config.S.ShortLinkInfo(s)
	if err != nil {
		respondWithError(w, err)
	} else {
		respondWithJSON(w, http.StatusOK, d)
	}

	fmt.Printf("Response: %v\n", d)
}

func (a *App) redirect(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("redirect short link! \n")
	fmt.Printf("Req: %v\n", r)
	vars := mux.Vars(r)
	fmt.Printf("shortlink: %v\n", vars["shortlink"])
	u, err := a.Config.S.Unshorten(vars["shortlink"])
	fmt.Printf("u: %v\n", u)
	if err != nil {
		respondWithError(w, err)
	} else {
		http.Redirect(w, r, u, http.StatusAccepted)
		return
	}
}
