package routes

import (
	mux "github.com/gorilla/mux"
	"github.com/ismael3s/alura/go/04/controllers"
)

func HandlerRequests() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", controllers.Home).Methods("GET")
	router.HandleFunc("/personalities", controllers.ShowAllPersonalities).Methods("GET")
	router.HandleFunc("/personalities/{id}", controllers.ShowPersonality).Methods("GET")

	return router
}
