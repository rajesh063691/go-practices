package handlers

import (
	"gorilla-mux/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// UpdateProduct updates the product
func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Printf("Handle PUT Handler")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}

	// getting prod value from request context
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	// update list
	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	p.l.Printf("Product Updated Successfully!!!")
}
