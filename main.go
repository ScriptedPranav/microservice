package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ScriptedPranav/microservice/handlers"
)

func main() {
	l := log.New(os.Stdout,"product-api",log.LstdFlags)

	helloHandler := handlers.NewHello(l)
	goodbyeHandler := handlers.NewGoodbye(l)
	sm :=  http.NewServeMux()
	sm.Handle("/hello",helloHandler)
	sm.Handle("/goodbye",goodbyeHandler)

	s := &http.Server{
		Addr : ":8080",
		Handler: sm,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}
	s.ListenAndServe()
}