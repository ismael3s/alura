package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/ismael3s/alura/go/03/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.FindAllProducts()

	temp.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		stringPrice := r.FormValue("preco")
		stringQuantity := r.FormValue("quantidade")

		price, err := strconv.ParseFloat(stringPrice, 64)
		if err != nil {
			http.Error(w, "Error on string to int convert", http.StatusInternalServerError)
			return
		}

		quantity, err := strconv.Atoi(stringQuantity)
		if err != nil {
			http.Error(w, "Error on string to int convert", http.StatusInternalServerError)
			return
		}

		fmt.Println(name, description, price, quantity)

		models.InsertProduct(name, description, price, quantity)

	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, "Error on string to int convert", http.StatusInternalServerError)
			return
		}

		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		stringPrice := r.FormValue("preco")
		stringQuantity := r.FormValue("quantidade")

		price, err := strconv.ParseFloat(stringPrice, 64)
		if err != nil {
			http.Error(w, "Error on string to int convert", http.StatusInternalServerError)
			return
		}

		quantity, err := strconv.Atoi(stringQuantity)
		if err != nil {
			http.Error(w, "Error on string to int convert", http.StatusInternalServerError)
			return
		}

		models.UpdateProduct(id, name, description, price, quantity)

	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, "Error on string to int convert", http.StatusInternalServerError)
	}

	models.DeleteProduct(id)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	product := models.GetProductById(idString)
	temp.ExecuteTemplate(w, "Edit", product)
}
