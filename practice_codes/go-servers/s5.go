package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
)

func loadFile(fileName string) (string, error) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(file), nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	file, err := loadFile(path.Base(r.URL.EscapedPath()))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Page not Found: 404")
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, file)

	//fmt.Fprintf(w, "Hello %s, welcome to your http server with HostName=%s", r.URL.Path[1:], r.Host)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}
