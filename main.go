package main

import (
	"fmt"
	"log"
	"os"

	// html/template
	// database/sql

	"encoding/json"
	"net/http"
)

func greeting() {
	fmt.Println("http://Mote.Archive/Dae")
	fmt.Println("Welcome to the Archivist Mote.")
	fmt.Println("-----------------------------------------------------------")
	fmt.Println("options: /home :: /motes :: /static/json :: ")
	fmt.Println()
	fmt.Println("[:Dae-W:] :: Listening on Port 3000...")
}

func getJSON(w http.ResponseWriter, r *http.Request) {
	file, err := os.ReadFile("./7_book-data.json")

	var payload map[string]interface{}
	_ = json.Unmarshal(file, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(file))

	fmt.Println("[:Dae-W:] :: 200 OK, Data Returned As JSON.")
}

func main() {
	greeting()

	http.HandleFunc("/Dae", getJSON)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}

/*
	log.Printf("origin: %s\n", payload["origin"])
	log.Printf("user: %s\n", payload["user"])
	log.Printf("status: %t\n", payload["active"])
*/
