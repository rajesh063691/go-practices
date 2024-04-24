package handlers

import (
	"errors"
	"log"
	"net/http"
	"product-api/data"
	"regexp"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) getPathID(r *http.Request) (int, error) {
	p.l.Printf("Inside getPathID")
	reg := regexp.MustCompile(`/([0-9]+)`)
	g := reg.FindAllStringSubmatch(r.URL.Path, -1)
	if len(g) != 1 {
		return -1, errors.New("invalid URI")
	}

	if len(g[0]) != 2 {
		return -1, errors.New("invalid URI")
	}

	intStr := g[0][1]
	id, err := strconv.Atoi(intStr)
	if err != nil {
		return -1, errors.New("string to int conversion error")
	}

	return id, nil
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// handle GET
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}

	//handle POST
	if r.Method == http.MethodPost {
		p.postProduct(w, r)
		return
	}

	//handle UPDATE
	if r.Method == http.MethodPut {
		// need to take id
		p.l.Printf("PUT")
		id, err := p.getPathID(r)
		if err != nil {
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}
		p.l.Printf("Got ID =%d", id)
		p.updateProduct(id, w, r)
		return
	}

	//handle DELETE
	if r.Method == http.MethodDelete {
		id, err := p.getPathID(r)
		if err != nil {
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}
		p.deleteProduct(id, w, r)
		return
	}

	//catch all
	w.WriteHeader(http.StatusMethodNotAllowed)
}

//gets all products
func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Printf("Handle GET Handler")

	lp := data.GetProducts()
	err := lp.ToJSON(w)

	if err != nil {
		p.l.Printf("unable to marshal json", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//post the product
func (p *Products) postProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Printf("Handle POST Handler")
	prod := data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		p.l.Printf("unable to unmarshal post request, error=%v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// update list
	data.UpdateList(&prod)

	p.l.Printf("Product Posted Successfully!!!")
}

//post the product
func (p *Products) updateProduct(id int, w http.ResponseWriter, r *http.Request) {
	p.l.Printf("Handle PUT Handler")

	prod := data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		p.l.Printf("unable to unmarshal put request, error=%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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

//delete the product
func (p *Products) deleteProduct(id int, w http.ResponseWriter, r *http.Request) {
	p.l.Printf("Handle DELETE Handler")

	prod := data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		p.l.Printf("unable to unmarshal put request, error=%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
