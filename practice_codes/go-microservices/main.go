package main

import (
	"context"
	"go-microservices/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	ll := log.New(os.Stdout, "product-api ", log.LstdFlags)
	hh := handlers.NewHello(ll)
	gh := handlers.NewGoodBye(ll)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	s := &http.Server{
		Addr:         ":9000",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			ll.Fatal(err)
		}
	}()

	sigChal := make(chan os.Signal, 1)
	signal.Notify(sigChal, os.Interrupt, os.Kill)

	sig := <-sigChal

	ll.Println("Received terminate signal. Gracefully Shutdown!", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
