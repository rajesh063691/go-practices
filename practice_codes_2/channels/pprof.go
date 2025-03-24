package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main_pprof() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	time.Sleep(30 * time.Second)
}
