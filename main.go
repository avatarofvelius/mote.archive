package main

import (
	"fmt"
	"log"

	"net/http"
)

func main() {
	http.HandleFunc("/", root)

	http.HandleFunc("/content", contentHandler)
	http.HandleFunc("/writer", writerHandler)
	http.HandleFunc("/book", bookHandler)
	http.HandleFunc("/character", characterHandler)

	http.Handle("/data/", http.StripPrefix("/data/", http.FileServer(http.Dir("_data"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("_static"))))

	fmt.Println("[:Dae-W:] :: Listening on Port 3000...")
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}

// Unused
/*
	greeting()
	http.HandleFunc("/json", getJSON)
	http.HandleFunc("/template", templateHandler)
*/
