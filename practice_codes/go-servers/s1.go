package main

import (
	"fmt"
	"io"
	"net/http"
)

type HttpHandler struct{}

func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	//data := []byte("Hello World!")

	//res.Write(data)

	io.WriteString(res, "Hello")
	fmt.Fprint(res, " World!")
	res.Write([]byte("."))
}

func main_1() {
	handler := HttpHandler{}
	http.ListenAndServe(":9000", handler)
}
