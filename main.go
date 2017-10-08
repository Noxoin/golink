package main

import (
	"fmt"
	"flag"
	"log"
	"net/http"

	"google.golang.org/appengine"
	"github.com/noxoin/golink"
)

var FLAG_port string
var FLAG_host string

func initFlags() {
	flag.StringVar(&FLAG_port, "port", "80", "Serving Port")
	flag.StringVar(&FLAG_host, "host", "localhost", "Serving Hostname")
	flag.Parse()
}

func main() {
	initFlags()

	http.HandleFunc("/", golink.Handler)
	http.HandleFunc("/_ah/health", healthCheckHandler)
	log.Printf("Server listening on port %s", FLAG_port)
	appengine.Main()
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}
