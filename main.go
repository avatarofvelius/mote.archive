package main

import (
	"log"

	"net/http"
)

func main() {
	greeting()

	http.HandleFunc("/", root)
	http.HandleFunc("/content", contentHandler)
	http.HandleFunc("/yaml", testyaml)

	http.Handle("/data/", http.StripPrefix("/data/", http.FileServer(http.Dir("_data"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("_static"))))

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}

// Unused
/*

	http.HandleFunc("/json", getJSON)
	http.HandleFunc("/template", templateHandler)

	http.HandleFunc("/writer", writerHandler)
	http.HandleFunc("/book", bookHandler)
	http.HandleFunc("/character", characterHandler)
*/
