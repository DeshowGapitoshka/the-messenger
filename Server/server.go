package main

import (
	"log"
	"net/http"
	lists "server/lists"
	messages "server/messages"
	users "server/users"
	"time"
)

type Pole struct {
	p1 int
	p2 int
	p3 int
	p4 int
	p5 string
	p6 string
	p7 time.Time
	p8 string
	p9 string
}

func main() {
	mux := http.NewServeMux()
	//	mux.HandleFunc("/api/health", checkHealth)
	mux.HandleFunc("POST /registration", users.Register)
	mux.HandleFunc("/auth", users.DistributionMethod)
	mux.HandleFunc("/messages", messages.PostMessage)
	mux.HandleFunc("GET /messages/{message_id}", messages.GetMessage)
	mux.HandleFunc("PUT /messages/{message_id}", messages.PutMessage)

	mux.HandleFunc("POST /chats", lists.CreateNewChat)
	mux.HandleFunc("POST /groups", lists.CreateNewGroup)
	mux.HandleFunc("POST /contacts", lists.AddNewContact)

	mux.HandleFunc("GET /userdata", users.GetUserDataById)
	mux.HandleFunc("GET /chats", lists.GetUserChatsById)
	mux.HandleFunc("GET /groups", lists.GetUserGroupsById)
	mux.HandleFunc("GET /contacts", lists.GetUserContactsById)

	// Оборачиваем весь маршрутизатор в CORS
	handler := enableCORSForMux(mux) //BE2681D73EB0903771C9B2DCC76FCFC04768FF7F
	handler = LoggingMiddleware(handler)

	log.Print("Listening on port 8080")
	err := http.ListenAndServeTLS("0.0.0.0:8080", "cert.pem", "key.pem", handler)

	if err != nil {
		log.Fatal(err)
	}
}

func enableCORSForMux(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "https://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Expose-Headers", "Set-Cookie")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("📥 %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
