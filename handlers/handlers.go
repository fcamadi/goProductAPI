package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/goProductAPI/entity"
)

// GetProductHandler is used to get data inside the products defined on our product catalog
func GetProductHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		//Read JSON file
		data, err := entity.GetProducts()
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		//Write body with JSON data
		rw.Header().Add("content-type", "application/json")
		rw.WriteHeader(http.StatusFound)
		rw.Write(data)
	}
}

// CreateProductHandler is used to create a new product and add to our product store.
func CreateProductHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// Read incoming JSON from request body
		data, err := io.ReadAll(r.Body)

		// If no body is associated return with StatusBadRequest
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		// Check if data is proper JSON (data validation)
		var product entity.Product
		err = json.Unmarshal(data, &product)
		if err != nil {
			rw.WriteHeader(http.StatusExpectationFailed)
			rw.Write([]byte("Invalid Data Format"))
			return
		}

		err = entity.AddProduct(product)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		// return after writing Body
		rw.WriteHeader(http.StatusCreated)
		rw.Write([]byte("Added New Product"))
	}
}
