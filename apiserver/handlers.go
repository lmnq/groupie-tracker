package apiserver

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

// IndexHandler ..
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	index, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println("gg")
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	locations := UniqLocations()
	err = index.Execute(w, struct {
		Artists       []Artist
		UniqLocations []string
	}{
		Artists:       Artists,
		UniqLocations: locations,
	})
	if err != nil {
		log.Println("ff")
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}

// DetailsHandler ..
func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, "/artists/") {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	if r.Method != http.MethodGet {
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.URL.Path[9:])
	if err != nil || id < 1 {
		ErrorHandler(w, http.StatusBadRequest)
		return
	}
	if id > len(Artists) {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	temp, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	err = temp.Execute(w, Artists[id-1])
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}

// SearchHandler ..
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/search" {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodPost {
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	input := r.FormValue("search")
	if input == "" {
		ErrorHandler(w, http.StatusBadRequest)
		return
	}
	category := r.FormValue("categories")
	answer, err := Search(input, category)
	if err != nil {
		ErrorHandler(w, http.StatusBadRequest)
		return
	}
	temp, err := template.ParseFiles("templates/search.html")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	temp.Execute(w, answer)
}

// FilterHandler ..
func FilterHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/filter" {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodPost {
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	r.ParseForm()
	// fmt.Println(r.FormValue("location"))
	answer, err := Filter(r)
	if err != nil {
		ErrorHandler(w, http.StatusBadRequest)
		return
	}
	temp, err := template.ParseFiles("templates/search.html")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	temp.Execute(w, answer)
}

// ErrorHandler ..
func ErrorHandler(w http.ResponseWriter, status int) {
	switch status {
	case 400:
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	case 404:
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	case 405:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	default:
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
