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

type Person struct {
	Id       int
	Login    string
	Password string
}

// Connection to BD and starting server
func main() {
	main2.StartServer()
	http.HandleFunc("/messages", messageHandler)
	http.HandleFunc("/getmessage", getMessage)
	http.HandleFunc("/api/health", checkHealth)
	http.HandleFunc("/api/register", register)
	http.HandleFunc("/api/login", login)
	http.HandleFunc("/getaccount", getPerson)
	log.Print("Listening on port 8080")
	err := http.ListenAndServe("localhost:8080", nil)
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

func register(w http.ResponseWriter, r *http.Request) {
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := main2.InputInBasePerson(person.Login, person.Password)
	if result == true {
		fmt.Fprintf(w, "Registred new account")
	} else {
		fmt.Fprintf(w, "This account already exists")
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := main2.CheckLogin(person.Login, person.Password)
	if result == true {
		fmt.Fprintf(w, "Вход выполнен успешно")
	} else {
		fmt.Fprintf(w, "Неправильный логин или пароль")
	}
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	person := main2.OutputFromBaseIdPerson(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

// Check the server status
func checkHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}
