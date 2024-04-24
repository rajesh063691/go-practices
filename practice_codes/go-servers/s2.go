package main

import (
	"fmt"
	"net/http"
)

func main_2() {

	// 1st way to create instance of ServerMux and call HandleFunc to register patterns & handlers
	// mux := http.NewServeMux()

	// mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(rw, "Hello World!")
	// })

	// mux.HandleFunc("/hello/golang", func(rw http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(rw, "Hello Golang!")
	// })

	// http.ListenAndServe(":9000", nil)

	// 2nd way to use http.HandleFunc to register patterns & handlers without explicitly creating instance of DerverMux, bcoz under the hood http package creats DefaultServerMux to handle this.So, one can pass Handler as nil and even though all the patterns will get register

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello World!")
	})

	http.HandleFunc("/hello/golang", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello Golang!")
	})

	// here passing handler as nil, and it will register all the patterns and handlers by the help of  DefaultServerMux
	http.ListenAndServe(":9000", nil)

}
