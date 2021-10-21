package main

import (
	a "apiserver/apiserver"
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", a.IndexHandler)
	http.HandleFunc("/artists/", a.DetailsHandler)
	http.HandleFunc("/search", a.SearchHandler)
	http.HandleFunc("/filter", a.FilterHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	fmt.Printf("The server is running on this address: http://localhost:8081\n")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
