package handlers

import (
	"gorilla-mux/data"
	"net/http"
)

// PostProducts post the product
func (p *Products) PostProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Printf("Handle POST Handler")
	// getting prod value from request context
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	// update list
	data.UpdateList(&prod)

	p.l.Printf("Product Posted Successfully!!!")
}
