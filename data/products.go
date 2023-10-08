package products

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"regexp"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"` // - means ignore this field
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"` // ,omitempty means ignore this field if it is empty
}

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}
func validateSKU(fl validator.FieldLevel) bool {
	reg := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := reg.FindAllString(fl.Field().String(), -1)
	if len(matches) != 1 {
		return false
	}
	return true
}
func (p *Product) FromJSON(r *http.Request) error {
	e := json.NewDecoder(r.Body)
	return e.Decode(p)
}

type Products []*Product

// Abstract the logic of writing to JSON
func (p *Products) ToJSON(w http.ResponseWriter) error {
	//this directly encodes the data to the writer in JSON format without having to create a buffer which is memory efficient
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// Abstract the logic of reading from Database
func GetProducts() Products {
	return productsList
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productsList = append(productsList, p)
}

func getNextID() int {
	lp := productsList[len(productsList)-1]
	return lp.ID + 1
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productsList[pos] = p
	return nil
}

func findProduct(id int) (*Product, int, error) {
	for i, p := range productsList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
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
