package cmd

import (
	"io/ioutil"
	"log"
	"net/http"

	// entities "github.com/4ubak/Groupie-tracker/internal/entities"
	internal "github.com/4ubak/Groupie-tracker/internal"
)

func Execute() {
	// _, _ = ioutil.ReadFile("./front/err.html")
	// changeDir := http.FileServer(http.Dir("style"))
	// http.Handle("/style/", http.StripPrefix("/style/", changeDir))
	http.HandleFunc("/", internal.Router)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
