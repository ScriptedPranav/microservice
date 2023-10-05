package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
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