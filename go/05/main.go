package main

import (
	"github.com/ismael3s/alura/go/05/database"
	"github.com/ismael3s/alura/go/05/routes"
)

func main() {
	database.ConnectToDB()

	routes.HandleRequests()
}
