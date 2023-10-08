package handlers

import (
	"log"
	"net/http"

	products "github.com/ScriptedPranav/microservice/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(w,r)
		return
	}
	if r.Method == http.MethodPost {
		p.addProduct(w,r)
		return
	}
	//catch all
	w.WriteHeader(http.StatusMethodNotAllowed)
}
func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	lp := products.GetProducts()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w,"Unable to marshal json",http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(w http.ResponseWriter, r*http.Request) {
	p.l.Println("Handle POST Product")
	prod := &products.Product{}
	err := prod.FromJSON(r)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
	}
	products.AddProduct(prod)
	//# in the Printf means print the struct in a pretty format
	p.l.Printf("Prod: %#v", prod)
}

func (p *Products) updateProduct(id int, w http.ResponseWriter, r*http.Request) {
	p.l.Println("Handle PUT request")
	prod := &products.Product{}
	err := prod.FromJSON(r)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
	}
	err = products.UpdateProduct(id,prod)
	if err == products.ErrProductNotFound {
		http.Error(w,"Product not found",http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w,"Product not found",http.StatusInternalServerError)
		return
	}

}

