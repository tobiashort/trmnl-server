package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

type DebugMiddleware struct {
	Active  bool
	Handler http.Handler
}

func (m DebugMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if m.Active {
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Println(string(dump))
	} else {
		log.Printf("%s %s", r.Method, r.URL)
	}
	m.Handler.ServeHTTP(w, r)
}
