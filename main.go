package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/Noxoin/golink/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	http.HandleFunc("/_ah/health", healthCheckHandler)
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}
