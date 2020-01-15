package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func runServer(addr string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Addr:", addr, "rootHandler URL:", r.URL.String())
	})
	mux.HandleFunc("/test/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Addr:", addr, "testHandler URL:", r.URL.String())
	})

	server := http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("starting server at", addr)
	log.Println("starting server at", addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func main() {
	go runServer(":8081")
	runServer(":8080")
}
