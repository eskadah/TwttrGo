package main

import (
	"encoding/json"
	"github.com/eskadah/TwttrGo/lib"
	"net/http"
)

func main() {
	http.HandleFunc("/search", searchResults)
	http.Handle("/", http.FileServer(http.Dir("../public")))
	http.ListenAndServe(":5000", nil)
}

func searchResults(w http.ResponseWriter, r *http.Request) {
	searchResults := lib.BuildSearchResults(r.FormValue("query"))
	js, err := json.Marshal(searchResults)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
