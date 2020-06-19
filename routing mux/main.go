package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// ShowDocument returns document title and page number from request
func ShowDocument(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]

	fmt.Fprintf(w, "You've requested the document: %s on page %s\n", title, page)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/documents/{title}/page/{page}", ShowDocument).Methods("GET")

	http.ListenAndServe(":9000", r)
}
