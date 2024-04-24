package handlers

import (
	"gorilla-mux/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// swagger:route DELETE /products/{id} deleteProduct
// Returns a list of products
// Responses:
//	201: noContent
//		description: The validation message
//
// DeleteProduct deletes a product from database
func (p *Products) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Printf("Handle DELETE Handler")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}

	// getting prod value from request context
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	// update list
	err = data.DeleteProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	p.l.Printf("Product Deleted Successfully!!!")
}
