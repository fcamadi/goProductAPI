package entity

import (
	"encoding/json"
	"os"
)

/*
The json tags assigned to the struct field let’s the json Marshaller know that
this field is found with this particular name in a json file.
*/
type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	IsAvailable bool    `json:"isAvailable"`
}

func GetProducts() ([]byte, error) {
	// Read JSON file
	data, err := os.ReadFile("./data/data.json")
	if err != nil {
		return nil, err
	}
	return data, nil
}

func AddProduct(product Product) error {
	// Load existing products and append the data to product list
	var products []Product
	data, err := os.ReadFile("./data/data.json")
	if err != nil {
		return err
	}
	// Load our JSON file to memory using array of products
	err = json.Unmarshal(data, &products)
	if err != nil {
		return err
	}
	// Add new Product to our list
	products = append(products, product)

	// Write Updated JSON file
	updatedData, err := json.Marshal(products)
	if err != nil {
		return err
	}
	err = os.WriteFile("./data/data.json", updatedData, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
