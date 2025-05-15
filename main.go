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
	var baseUrl string
	var accessToken string
	var port string
	flag.BoolVar(&debug, "debug", false, "enable debug mode")
	flag.StringVar(&baseUrl, "baseUrl", "", "the base URL")
	flag.StringVar(&accessToken, "accessToken", "", "access token")
	flag.StringVar(&port, "port", "8080", "port")
	flag.Parse()

	log.Println("Debug:", debug)
	log.Println("Base URL:", baseUrl)
	if accessToken != "" {
		log.Println("Access Token: ****")
	} else {
		log.Println(cfmt.Sprint("Access Token: #r{not set}"))
	}

	http.Handle("/", CatchAllHandler{})
	http.Handle("POST /api/log", LogHandler{})
	http.Handle("GET /api/display", DisplayHandler{BaseUrl: baseUrl})
	http.Handle("GET /image", ImageHandler{})

	var mux http.Handler
	mux = http.DefaultServeMux
	mux = AuthMiddleware{AccessToken: accessToken, Handler: mux}
	mux = DebugMiddleware{Active: debug, Handler: mux}

	addr := ":" + port
	log.Println("Start listening on", addr)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatalln(err)
	}
}
