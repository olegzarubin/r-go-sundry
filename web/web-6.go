package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("RequestID", "d41d8cd98f00b204")

	fmt.Fprintln(w, "You browser is", r.UserAgent())
	fmt.Fprintln(w, "You accept", r.Header.Get("Accept"))
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("starting server at :8080")
	log.Println("starting server at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
