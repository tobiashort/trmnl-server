package main

import (
	"net/http"
)

type ImageHandler struct{}

func (ImageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "image/bmp")
	http.ServeFile(w, r, "test.bmp")
}
