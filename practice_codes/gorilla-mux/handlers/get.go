package handlers

import (
	"gorilla-mux/data"
	"net/http"
)

// swagger:route GET /products listProducts
// Returns a list of products
// Responses:
//	200: productResponse

// GetProducts gets all products
func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Printf("Handle GET Handler")

	lp := data.GetProducts()
	err := lp.ToJSON(w)

	if err != nil {
		p.l.Printf("unable to marshal json", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
