package routes

import (
	mux "github.com/gorilla/mux"
	"github.com/ismael3s/alura/go/04/controllers"
	"github.com/ismael3s/alura/go/04/middlewares"
)

func HandlerRequests() *mux.Router {
	router := mux.NewRouter()

	router.Use(middlewares.ContentTypeMiddleware)

	router.HandleFunc("/", controllers.Home).Methods("GET")
	router.HandleFunc("/personalities", controllers.ShowAllPersonalities).Methods("GET")
	router.HandleFunc("/personalities/{id}", controllers.ShowPersonality).Methods("GET")
	router.HandleFunc("/personalities/{id}", controllers.RemovePersonality).Methods("DELETE")
	router.HandleFunc("/personalities/{id}", controllers.UpdatePersonality).Methods("PUT")
	router.HandleFunc("/personalities", controllers.CreatePersonality).Methods("POST")
	return router
}
