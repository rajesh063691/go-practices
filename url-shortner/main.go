package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/renstrom/shortuuid"
)

// mapper to hold the mapping of the short url to the long url
type Mapper struct {
	Mapping map[string]string
	Lock    sync.Mutex
}

var urlMapper Mapper

func init() {
	urlMapper = Mapper{
		Mapping: make(map[string]string, 0),
	}
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/health", HealthHandler)
	r.HandleFunc("/short-it", CreateShortURLHandler).Methods("POST")
	r.HandleFunc("/short/{key}", GetShortURLHandler).Methods("GET")

	err := http.ListenAndServe(":4444", r)
	if err != nil {
		log.Fatalf("error starting server: %v", err)
	}

}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("server is running..."))
}

func CreateShortURLHandler(w http.ResponseWriter, r *http.Request) {
	urlMapper.Lock.Lock()
	defer urlMapper.Lock.Unlock()

	// read the long url from the request
	// generate a short url
	// store the mapping in the mapper
	// return the short url
	longUrl := r.URL.Query().Get("longUrl")
	if longUrl == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("long url is missing"))
		return
	}

	encURL := shortuuid.New()
	urlMapper.Mapping[encURL] = longUrl

	shorURL := fmt.Sprintf("http://localhost:4444/short/%s", encURL)

	log.Println("short url created successfully: ", shorURL)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("short url: %s", shorURL)))

}

func GetShortURLHandler(w http.ResponseWriter, r *http.Request) {
	urlMapper.Lock.Lock()
	defer urlMapper.Lock.Unlock()

	// read the short url from the request
	// look up the mapping in the mapper
	// redirect to the long url
	key := mux.Vars(r)["key"]
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("short url is missing"))
		return
	}

	if _, ok := urlMapper.Mapping[key]; !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("short url not found"))
		return
	}

	http.Redirect(w, r, urlMapper.Mapping[key], http.StatusFound)
}
