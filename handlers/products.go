package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"
	products "github.com/ScriptedPranav/microservice/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}


func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	lp := products.GetProducts()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w,"Unable to marshal json",http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(w http.ResponseWriter, r*http.Request) {
	p.l.Println("Handle POST Product")
	prod := r.Context().Value(KeyProduct{}).(*products.Product)
	products.AddProduct(prod)
	//# in the Printf means print the struct in a pretty format
	p.l.Printf("Prod: %#v", prod)
}

func (p *Products) UpdateProduct(w http.ResponseWriter, r*http.Request) {
	vars := mux.Vars(r)
	id,err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w,"Unable to convert id",http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT request")
	//get the product from the context using the key and cast it to the *products.Product type
	prod := r.Context().Value(KeyProduct{}).(*products.Product)
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

type KeyProduct struct{}

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := &products.Product{}
		err := prod.FromJSON(r)
		if err != nil {
			http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(),KeyProduct{},prod)
		req := r.WithContext(ctx)
		next.ServeHTTP(w,req)
	})
}

