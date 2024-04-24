// Package classification Product API
//
// Documentation for Product API123
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package handlers

import (
	"context"
	"gorilla-mux/data"
	"log"
	"net/http"
)

// A list of products returns in the response
// swagger:response productResponse
type productResponse struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// swagger:response noContent
type noContentWrapper struct {
	// All products in the system
	// in: body
	Body struct{}
}

// swagger:parameters deleteProduct
type productIDParameterWrapper struct {
	// The id of the product to delete from the database
	// in: path
	// required: true
	ID int `json:"id"`
}

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

type KeyProduct struct{}

func (p *Products) MiddleWareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Printf("[ERROR] unable to unmarshal request, error=%v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//Validate the prod struct
		err = prod.Validator()
		if err != nil {
			p.l.Printf("[ERROR] unable to validate struct, error=%v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
