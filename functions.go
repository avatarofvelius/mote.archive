package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func root(w http.ResponseWriter, req *http.Request) {
	fmt.Println("[:Dae-W:] :: 200 OK, Root Loaded.")
}

func writerHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(
		"./_static/template/base.html",
		"./_static/template/writer.html",
	)
	tmpl.Execute(w, nil)

	fmt.Println("[:Dae-W:] :: 200 OK, Writer Loaded.")
}

func bookHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(
		"./_static/template/base.html",
		"./_static/template/book.html",
	)
	tmpl.Execute(w, nil)

	fmt.Println("[:Dae-W:] :: 200 OK, Writer Loaded.")
}

func characterHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(
		"./_static/template/base.html",
		"./_static/template/character.html",
	)
	tmpl.Execute(w, nil)

	fmt.Println("[:Dae-W:] :: 200 OK, Writer Loaded.")
}

// Unused Functaions

/*
func templateHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(
		"./_static/template/base.html",
		"./_static/template/content.html",
	)
	tmpl.Execute(w, nil)

	fmt.Println("[:Dae-W:] :: 200 OK, Template Loaded.")
}
*/

/*
func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("Index Page <hr><br><a href='/template'>Template</a><br><br><a href='data/'>Data</a><br><a href='/static'>Static Files</a>"))
	fmt.Println("[:Dae-W:] :: 200 OK, Index Loaded.")
}
*/

/*
func getJSON(w http.ResponseWriter, r *http.Request) {
	file, err := os.ReadFile("./data/character/cvb-564-86-Hybrid_character-data.json")

	var payload map[string]interface{}
	_ = json.Unmarshal(file, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(file))

	fmt.Println("[:Dae-W:] :: 200 OK, Data Returned As JSON.")
}
*/

/*
func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("Index Page <hr><br><a href='/template'>Template</a><br><br><a href='data/'>Data</a><br><a href='/static'>Static Files</a>"))
	fmt.Println("[:Dae-W:] :: 200 OK, Index Loaded.")
}
*/

/*
func greeting() {
	fmt.Println("-----------------------------------------------------------")
	fmt.Println("Welcome to the Archivist Mote.")
	fmt.Println()
	fmt.Println("[:http://Mote.Archive/:]")
	fmt.Println("[:options:]   --   /home :: /motes :: /books :: /characters")
	fmt.Println()
	fmt.Println("[:Dae-W:] :: Okaeri Inaeri,")
	fmt.Println("[:Dae-W:] :: Listening on Port 3000...")
}
*/
