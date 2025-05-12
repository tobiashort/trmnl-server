package main

import (
	"net/http"
)

type CatchAllHandler struct {
}

func (CatchAllHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}
