package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ismael3s/alura/go/04/models"
	"github.com/ismael3s/alura/go/04/routes"
)

const PORT = ":8080"

func main() {
	fmt.Println("Server running at port " + PORT)

	models.Personalities = []models.Personality{
		{
			Id:      1,
			Name:    "Ismael",
			History: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. ",
		},
		{
			Id:      2,
			Name:    "Leonardo",
			History: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. ",
		},
	}

	router := routes.HandlerRequests()
	log.Fatal(http.ListenAndServe(PORT, router))
}
