package products

import (
	"encoding/json"
	"net/http"
	"time"
)

type Product struct {
	ID          int `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       float32 `json:"price"`
	SKU         string `json:"sku"`
	CreatedOn   string `json:"-"` // - means ignore this field
	UpdatedOn   string `json:"-"`
	DeletedOn   string	`json:"-"` // ,omitempty means ignore this field if it is empty
}

type Products []*Product

//Abstract the logic of writing to JSON
func (p *Products) ToJSON(w http.ResponseWriter) error {
	//this directly encodes the data to the writer in JSON format without having to create a buffer which is memory efficient
	e := json.NewEncoder(w)
	return e.Encode(p)
}

//Abstract the logic of reading from Database
func GetProducts() Products {
	return productsList
}

var productsList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
