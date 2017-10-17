package main

import (
	"fmt"
	"flag"
	"log"
	"net/http"

	"github.com/noxoin/golink/handlers"
)

var FLAG_port string

func initFlags() {
	flag.StringVar(&FLAG_port, "port", "8080", "Serving Port")
	flag.Parse()
}

func main() {
	initFlags()
	handlers.InitHandlers()

	http.HandleFunc("/_ah/health", healthCheckHandler)
	log.Printf("Server listening on port %s", FLAG_port)
	log.Fatal(http.ListenAndServe(":" + FLAG_port, nil))
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}
