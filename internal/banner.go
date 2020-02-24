package internal

import (
	"fmt"
	"net/http"
	"text/template"
)

func Router(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		IndexPage(w, r)
	} else if r.Method == "POST" {
		// SearchPage(w, r)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func IndexPage(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s", err404)
		return
	}
	artists, err := GetAllData()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	t, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if err := t.Execute(w, artists); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
