package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Desciption string  `json:"description"`
	Price      float32 `json:"price"`
	SKU        string  `json:"sku"`
	CreatedOn  string  `json:"-"`
	UpdatedOn  string  `json:"-"`
	DeletedOn  string  `json:"-"`
}

var ErrProductNotFound = fmt.Errorf("product not found")

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(&p)
}

func UpdateList(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[pos] = p
	return nil
}

func DeleteProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}
	// appending in same slice from 0 to ele, and ele+1 to last element. last ele is ignored internally which slicing
	productList = append(productList[:pos], productList[pos+1:]...)
	return nil
}

func findProduct(id int) (p *Product, index int, err error) {
	for i, prod := range productList {
		if id == prod.ID {
			return prod, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

func getNextID() int {
	curID := productList[len(productList)-1]
	return curID.ID + 1
}

func GetProducts() Products {
	return productList
}

var productList = []*Product{
	{
		ID:         1,
		Name:       "Latte",
		Desciption: "mikly coffee",
		Price:      2.45,
		SKU:        "latte123",
		CreatedOn:  time.Now().UTC().String(),
		UpdatedOn:  time.Now().UTC().String(),
		DeletedOn:  "",
	},
	{
		ID:         2,
		Name:       "Espresso",
		Desciption: "short and strong without milk",
		Price:      1.99,
		SKU:        "espresso123",
		CreatedOn:  time.Now().UTC().String(),
		UpdatedOn:  time.Now().UTC().String(),
		DeletedOn:  "",
	},
}
