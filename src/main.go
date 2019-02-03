package main

import (
	"./controllers"
	"./migrations"
	"net/http"
)

func main() {
	migrator.Run("postgres://postgres:postgres@localhost:5432?sslmode=disable")

	http.HandleFunc("/", products.Index)
	http.HandleFunc("/products/statuses", products.GetList)

	http.ListenAndServe(":9000", nil)
}
