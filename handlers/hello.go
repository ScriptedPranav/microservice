package handlers

import (
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")

	_, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w,"Oops",http.StatusBadRequest)
		return
	}
}
