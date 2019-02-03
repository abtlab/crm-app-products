package products

import (
	"../types/statuses"
	"encoding/json"
	"net/http"
)

// Get :
func Get(writer http.ResponseWriter, request *http.Request) {
	id := request.URL.Query()["id"][0]
	productStatus := products.ProductStatus{ID: id, Name: "Product status 1"}

	json.NewEncoder(writer).Encode(productStatus)
}

// GetList :
func GetList(writer http.ResponseWriter, request *http.Request) {
	productStatuses := []products.ProductStatus{
		{ID: "1", Name: "Product status 1"},
		{ID: "2", Name: "Product status 2"},
	}

	json.NewEncoder(writer).Encode(productStatuses)
}
