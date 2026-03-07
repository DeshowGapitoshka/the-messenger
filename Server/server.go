package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	main2 "server/database"
)

type Message struct {
	Id      int
	User_id int
	Data    string
}

// Connection to BD and starting server
func main() {
	main2.StartServer()
	http.HandleFunc("/messages", messageHandler)
	http.HandleFunc("/getmessage", getMessage)
	http.HandleFunc("/health", checkHealth)
	log.Print("Listening on port 8050")
	err := http.ListenAndServe("localhost:8050", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Distribution of requests
func messageHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getMessages(w, r)
	case http.MethodPost:
		postMessage(w, r)
	default:
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}

// Output message from DB to localhost:8050/messages
func getMessages(w http.ResponseWriter, r *http.Request) {
	message := main2.OutputFromBase()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

func getMessage(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	message := main2.OutputFromBaseID(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

// Input message to DB from http.Request
func postMessage(w http.ResponseWriter, r *http.Request) {
	var message Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	main2.InputInBase(message.User_id, message.Data)
	fmt.Fprintf(w, "post new message '%v'", message)
}

// Check the server status
func checkHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}