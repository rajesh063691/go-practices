package main

import (
	"fmt"
	"testing"

	"swagger-docs/client/client"
	"swagger-docs/client/client/products"
)

func TestProductApi(t *testing.T) {
	// c := client.Default
	// param := products.NewListProductsParams()
	// prod, err := c.Products.ListProducts(param)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// fmt.Printf("prod=%#v", prod)
	//var ft *strfmt.Registry

	cfg := client.DefaultTransportConfig().WithHost("localhost:9000")
	// var tconfig = client.TransportConfig{}
	// tconfig.BasePath = "/"
	// tconfig.Host = "localhost:9000"

	c := client.NewHTTPClientWithConfig(nil, cfg)

	param := products.NewListProductsParams()
	prod, err := c.Products.ListProducts(param)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("prod=%#v", prod)
}
