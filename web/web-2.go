package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "rootHandler URL:", r.URL.String())
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "testHandler URL:", r.URL.String())
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/test/", testHandler)

	server := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("starting server at :8080")
	log.Println("starting server at :8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
