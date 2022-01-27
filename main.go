package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	fileServer := http.FileServer(http.Dir("./html")) // New code
	http.Handle("/", fileServer)                      // New code
	http.HandleFunc("/message", handlerMessage)
	http.HandleFunc("/example-data", returnJSONExample)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handlerMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func returnJSONExample(w http.ResponseWriter, r *http.Request) {
	listJSONobj := []JsonExample{
		{
			Key:   "first",
			Value: "any value",
		},
		{
			Key:   "second",
			Value: "can be",
		},
		{
			Key:   "third",
			Value: "entered",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	// get a payload p := Payload{d}
	json.NewEncoder(w).Encode(listJSONobj)
	return
}
