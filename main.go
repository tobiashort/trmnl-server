package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/tobiashort/cfmt"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetPrefix(cfmt.Sprint("#b{LOG: }"))

	var debug bool
	flag.BoolVar(&debug, "debug", false, "enable debug mode")
	flag.Parse()

	log.Println("Debug:", debug)

	http.Handle("/", CatchAllHandler{})
	http.Handle("POST /api/log", LogHandler{})
	http.Handle("GET /api/display", DisplayHandler{})
	http.Handle("GET /image", ImageHandler{})

	var mux http.Handler
	mux = http.DefaultServeMux
	mux = DebugMiddleware{Active: debug, Handler: mux}

	addr := ":8080"
	log.Println("Start listening on", addr)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatalln(err)
	}
}
