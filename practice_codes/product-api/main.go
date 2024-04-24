package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"product-api/handlers"
	"time"
)

func main() {
	ll := log.New(os.Stdout, "product-api ", log.LstdFlags)
	ph := handlers.NewProducts(ll)

	sm := http.NewServeMux()
	sm.Handle("/", ph)

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

	// this will be the holding point untill returns from signal
	sig := <-sigChal

	ll.Println("Received terminate signal. Gracefully Shutting down!", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
