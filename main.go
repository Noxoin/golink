package main

import (
	"fmt"
	"flag"
	"log"
	"net/http"

	_ "github.com/noxoin/golink/server"
	"google.golang.org/appengine"
)

var FLAG_port string

func init() {
	flag.StringVar(&FLAG_port, "port", "8080", "Serving Port")
	flag.Parse()
}

func main() {
	http.HandleFunc("/_ah/health", healthCheckHandler)
	log.Printf("Server listening on port %s", FLAG_port)
	appengine.Main()
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}
