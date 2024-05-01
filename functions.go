package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"gopkg.in/yaml.v3"
)

// Functions

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

// Routes

func root(w http.ResponseWriter, req *http.Request) {
	fmt.Println("[:Dae-W:] :: 200 OK, Root Loaded.")
}

func contentHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(
		"./_static/template/base.html",
		"./_static/template/content.html",
	)
	tmpl.Execute(w, r)

	fmt.Println("[:Dae-W:] :: 200 OK, Content Loaded.")
}

func testyaml(w http.ResponseWriter, req *http.Request) {

	/* lists in alpabecial order in map.
	possible solution :
	1_character-id 2_first_name
	a_character-id b_first_name
	*/

	var data = `
character-id: cvb-564-86-Hybrid
first-name: jhkl
last-name: lhj
gender: Female
ethnicity: Danish
portrait-path: b.jpg
notes: notes.md	
`
	type T struct {
		id        string `yaml:"character-id"`
		first     string `yaml:"first-name"`
		last      string `yaml:"last-name"`
		gender    string `yaml:"gender"`
		ethnicity string `yaml:"ethnicity"`
		portrait  string `yaml:"portrait-path"`
		notes     string `yaml:"notes"`
	}

	t := T{}

	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println(string(d))

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Println(m)

	tmpl, _ := template.ParseFiles(
		"./_static/template/base.html",
		"./_static/template/yaml.html",
	)
	tmpl.Execute(w, &m)

	fmt.Println("[:Dae-W:] :: 200 OK, YAML Loaded.")
}

// Unused Functaions

/*
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

	fmt.Println("[:Dae-W:] :: 200 OK, Book Loaded.")
}

func characterHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(
		"./_static/template/base.html",
		"./_static/template/character.html",
	)
	tmpl.Execute(w, nil)

	fmt.Println("[:Dae-W:] :: 200 OK, Character Loaded.")
}

*/

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

 */
