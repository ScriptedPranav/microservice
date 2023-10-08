package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"github.com/ScriptedPranav/microservice/handlers"
	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout,"product-api",log.LstdFlags)


	productHandler := handlers.NewProducts(l)
	sm :=  mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products",productHandler.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products/{id:[0-9]+}",productHandler.UpdateProduct)
	putRouter.Use(productHandler.MiddlewareProductValidation)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products",productHandler.AddProduct)
	postRouter.Use(productHandler.MiddlewareProductValidation)

	s := &http.Server{
		Addr : ":8080",
		Handler: sm,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan,os.Interrupt)
	signal.Notify(signalChan,os.Kill)

	sig := <-signalChan
	l.Println(sig)
	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(tc)
}