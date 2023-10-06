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

