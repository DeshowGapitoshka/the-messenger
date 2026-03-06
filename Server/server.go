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

func main() {
	main2.StartServer()
	http.HandleFunc("/messages", messageHandler)
	http.HandleFunc("/health", checkHealth)
	log.Print("Listening on port 8050")
	err := http.ListenAndServe("localhost:8050", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getMessage(w, r)
	case http.MethodPost:
		postMessage(w, r)
	default:
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}

func getMessage(w http.ResponseWriter, r *http.Request) {
	message := main2.OutputFromBase()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

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

func checkHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}
