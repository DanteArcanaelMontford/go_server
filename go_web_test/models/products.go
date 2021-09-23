package models

import (
	"web/server/database"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetAllProducts() []Product {
	db := database.ConnectDatabase()

	selectAllProducts, err := db.Query("select * from products order by id asc")

	if err != nil {
		panic(err.Error())
	}

	product := Product{}

	products := []Product{}

	for selectAllProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err := selectAllProducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity

		products = append(products, product)
	}
	defer db.Close()
	return products
}

func CreateNewProduct(name, description string, price float64, quantity int) {
	db := database.ConnectDatabase()

	insertionOnDatabase, err := db.Prepare("insert into products(product_name, product_description, product_price, product_quantity) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertionOnDatabase.Exec(name, description, price, quantity)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := database.ConnectDatabase()
	deleteProduct, err := db.Prepare("delete from products where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(id)
	defer db.Close()
}

func EditProduct(id string) Product {
	db := database.ConnectDatabase()
	productSelected, err := db.Query("select * from products where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	updatedProduct := Product{}

	for productSelected.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = productSelected.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		updatedProduct.Id = id
		updatedProduct.Name = name
		updatedProduct.Description = description
		updatedProduct.Price = price
		updatedProduct.Quantity = quantity

	}
	defer db.Close()
	return updatedProduct
}

func UpdateProduct(id int, name string, description string, price float64, quantity int) {
	db := database.ConnectDatabase()

	UpdateProduct, err := db.Prepare("update products set product_name=$1, product_description=$2, product_price=$3, product_quantity=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	UpdateProduct.Exec(name, description, price, quantity, id)
	defer db.Close()
}
