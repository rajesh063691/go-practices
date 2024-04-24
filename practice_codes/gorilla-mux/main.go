package main

import (
	"context"
	"gorilla-mux/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-openapi/runtime/middleware"

	"github.com/gorilla/mux"
)

//var bindAddress = env.String("BIND_ADDRESS", false, ":9000", "Bind address fro the Server")

func main() {

	//	env.Parse()

	ll := log.New(os.Stdout, "product-api ", log.LstdFlags)
	ph := handlers.NewProducts(ll)

	sm := mux.NewRouter()
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.PostProducts)
	postRouter.Use(ph.MiddleWareValidateProduct)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct)
	putRouter.Use(ph.MiddleWareValidateProduct)

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id:[0-9]+}", ph.DeleteProduct)
	deleteRouter.Use(ph.MiddleWareValidateProduct)

	// middleware for swagger
	ops := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(ops, nil)
	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

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
