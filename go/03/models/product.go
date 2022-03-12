package models

import (
	"log"

	"github.com/ismael3s/alura/go/03/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func FindAllProducts() []Product {
	var products []Product

	db := db.ConnectToDB()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM products")

	if err != nil {
		return []Product{}
	}

	for rows.Next() {
		var p Product

		err := rows.Scan(&p.Id, &p.Name, &p.Description, &p.Price, &p.Quantity)

		if err != nil {
			log.Fatal("Error on scan product: ", err.Error())
		}

		products = append(products, p)
	}

	return products
}

func InsertProduct(name, decription string, price float64, quantity int) {
	db := db.ConnectToDB()
	defer db.Close()

	query, err := db.Prepare("INSERT INTO products (name, description, price, quantity) VALUES ($1, $2, $3, $4)")

	if err != nil {
		log.Fatal("Error on insert product: ", err.Error())
	}

	_, err = query.Exec(name, decription, price, quantity)

	if err != nil {
		log.Fatal("Failed to insert product: ", err.Error())
	}

}

func DeleteProduct(id int) {
	db := db.ConnectToDB()

	defer db.Close()

	row, err := db.Prepare("DELETE FROM products WHERE id = $1")

	if err != nil {
		log.Fatal("Error on delete product: ", err.Error())
	}

	_, err = row.Exec(id)

	if err != nil {
		log.Fatal("Failed to delete product: ", err.Error())
	}
}

func GetProductById(id string) Product {
	db := db.ConnectToDB()
	defer db.Close()

	row, err := db.Query("SELECT * FROM products WHERE id = $1", id)

	if err != nil {
		log.Fatalln("Error on get product by id: ", err.Error())
	}

	product := Product{}

	for row.Next() {
		err := row.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Quantity)

		if err != nil {
			log.Fatal("Error on scan product: ", err.Error())
		}
	}
	return product

}

func UpdateProduct(id int, name, description string, price float64, quantity int) {
	db := db.ConnectToDB()

	defer db.Close()

	query, err := db.Prepare("UPDATE products SET name = $2, description = $3, price = $4, quantity = $5 WHERE id = $1 ")

	if err != nil {
		log.Fatal("Failed to prepare query", err.Error())
	}

	_, err = query.Exec(id, name, description, price, quantity)

	if err != nil {
		log.Fatalln("Failed to execute query", err.Error())
	}

}
