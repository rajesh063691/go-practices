package main

import (
	"fmt"
	"net/http"
)

func handler_4(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s, welcome to your http server with HostName=%s", r.URL.Path[1:], r.Host)
}

func main_4() {
	http.HandleFunc("/", handler_4)
	http.ListenAndServe(":9000", nil)
}
