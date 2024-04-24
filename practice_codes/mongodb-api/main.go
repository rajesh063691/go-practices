package main

import (
	"fmt"
	"log"
	"mongoapi/router"
	"net/http"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Server is getting started ....")
	fmt.Println("Listning at port 4000 ...")
	log.Fatal(http.ListenAndServe(":4000", r))
}
