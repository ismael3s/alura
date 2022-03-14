package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ismael3s/alura/go/04/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func ShowAllPersonalities(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalities)
}

func ShowPersonality(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Fprint(w, "Id inválido", http.StatusBadRequest)
		return
	}

	for _, personality := range models.Personalities {

		if personality.Id == id {
			json.NewEncoder(w).Encode(personality)
			return
		}
	}

	Error("Not found", w)
}

func Error(message string, w http.ResponseWriter) {
	response := make(map[string]string)

	response["error"] = message

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	w.Write(jsonResponse)
}
