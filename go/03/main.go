package main

import (
	"log"
	"net/http"

	"github.com/ismael3s/alura/go/03/routes"
)

func main() {
	PORT := ":8080"

	routes.LoadRouter()
	log.Fatal(http.ListenAndServe(PORT, nil))
}
