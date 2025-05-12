package main

import (
	"log"
	"net/http"
)

type AuthMiddleware struct {
	AccessToken string
	Handler     http.Handler
}

func (m AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Access-Token") != m.AccessToken {
		log.Println("Invalid access token")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	m.Handler.ServeHTTP(w, r)
}
