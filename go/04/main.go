package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/ismael3s/alura/go/04/database"
	"github.com/ismael3s/alura/go/04/routes"
)

const PORT = ":8080"

func main() {
	fmt.Println("Server running at port " + PORT)

	database.ConnectTODB()

	router := routes.HandlerRequests()
	log.Fatal(http.ListenAndServe(PORT, handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(router)))
}
