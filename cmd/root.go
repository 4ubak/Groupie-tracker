package cmd

import (
	"io/ioutil"
	"log"
	"net/http"
	internal "github.com/4ubak/groupie-tracker/internal"
)

func Execute() {
	err404, _ = ioutil.ReadFile("./front/err.html")
	http.HandleFunc("/", internal.Router)
	// changeDir := http.FileServer(http.Dir("style"))
	// http.Handle("/style/", http.StripPrefix("/style/", changeDir))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
