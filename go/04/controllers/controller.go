package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ismael3s/alura/go/04/database"
	"github.com/ismael3s/alura/go/04/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func ShowAllPersonalities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var personalities []models.Personality

	database.DB.Find(&personalities)

	json.NewEncoder(w).Encode(personalities)
}

func ShowPersonality(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		message := fmt.Sprintf("Id inválido: %s", params["id"])
		http.Error(w, message, http.StatusBadRequest)
		return
	}

	var personality models.Personality

	result := database.DB.First(&personality, id)

	if result.Error != nil {
		Error("Not found", w)
		return
	}

	json.NewEncoder(w).Encode(personality)
}

func CreatePersonality(w http.ResponseWriter, r *http.Request) {
	var personality models.Personality

	json.NewDecoder(r.Body).Decode(&personality)

	database.DB.Create(&personality)

	json.NewEncoder(w).Encode(personality)
}

func RemovePersonality(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Fprint(w, "Id inválido", http.StatusBadRequest)
		return
	}

	var personality models.Personality

	result := database.DB.First(&personality, id)

	if result.Error != nil {
		Error("Not found", w)
		return
	}

	database.DB.Delete(&personality)

	json.NewEncoder(w).Encode(personality)
}

func UpdatePersonality(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Fprint(w, "Id inválido", http.StatusBadRequest)
		return
	}

	var personality models.Personality
	var personalityUpdated models.Personality

	err = json.NewDecoder(r.Body).Decode(&personalityUpdated)
	if err != nil {
		fmt.Fprint(w, "Error", http.StatusInternalServerError)
	}

	if personalityUpdated.Id != id {
		http.Error(w, "Ids não correspondem", http.StatusBadRequest)
		return
	}

	result := database.DB.First(&personality, id)

	if result.Error != nil {
		log.Println(result.Error)
		Error("Not found", w)
		return
	}

	database.DB.Model(&personality).Updates(personalityUpdated)

	json.NewEncoder(w).Encode(personalityUpdated)

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

	w.WriteHeader(http.StatusBadRequest)
	w.Write(jsonResponse)
}
