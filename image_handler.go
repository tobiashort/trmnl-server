package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/tobiashort/trmnl-server/watch"
)

type ImageHandler struct{}

func (ImageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "image/bmp")
	t := time.Now()
	d := watch.Image(t)
	w.Header().Add("Content-Length", fmt.Sprintf("%d", len(d)))
	w.WriteHeader(200)
	_, err := w.Write(d)
	if err != nil {
		log.Println(err)
	}
}
