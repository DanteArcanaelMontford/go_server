package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"web/server/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	products := models.GetAllProducts()

	templates.ExecuteTemplate(w, "index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "new", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")

		description := r.FormValue("description")

		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			log.Println("Error in the price conversion: ", err)
		}

		quantity, err := strconv.Atoi(r.FormValue("quantity"))
		if err != nil {
			log.Println("Error in the quantity conversion: ", err)
		}

		models.CreateNewProduct(name, description, price, quantity)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	models.DeleteProduct(productId)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	product := models.EditProduct(productId)
	templates.ExecuteTemplate(w, "edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			log.Println("Error while convert id:", err)
		}
		name := r.FormValue("name")

		description := r.FormValue("description")

		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			log.Println("Error while convert price:", err)
		}

		quantity, err := strconv.Atoi(r.FormValue("quantity"))
		if err != nil {
			log.Println("Error while conver quantity:", err)
		}

		models.UpdateProduct(id, name, description, price, quantity)
	}
	http.Redirect(w, r, "/", 301)
}
