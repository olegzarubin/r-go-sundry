package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Main page")
}

func main() {
	http.HandleFunc("/", handler)

	http.HandleFunc("/page", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Single page:", r.URL.String())
	})

	http.HandleFunc("/pages/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Multiple pages:", r.URL.String())
	})

	fmt.Println("starting server at :8080")
	log.Println("starting server at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
