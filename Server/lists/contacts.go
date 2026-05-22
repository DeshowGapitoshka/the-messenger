package list

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	lists "server/database/lists"
	tokens "server/database/tokens"
)

type Contact struct {
	Id         int    `json:"id"`
	User_id    int    `json:"user_id"`
	Contact_id int    `json:"contact_id"`
	Created_at string `json:"created_at"`
}

func AddNewContact(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	fmt.Println(cookie)
	if err != nil {
		w.WriteHeader(429)
		return
	}
	token := fmt.Sprint(cookie)
	if token == "nil" {
		w.WriteHeader(500)
		return
	}
	token = token[6:]
	user_id := tokens.CheckToken(token)

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	contact := make(map[string]interface{})
	err = json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	contact_id := contact["contact_id"].(float64)

	result := lists.AddNewContact(user_id, int(contact_id))

	switch result {
	case 200:
		response := Response{
			Status:  200,
			Message: "Контакт успешно добавлен",
			Body:    nil,
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "%s", jsonData)
		return
	case 409:
		w.WriteHeader(409)
		return

	default:
		w.WriteHeader(500)
		return
	}
}

func GetUserContactsById(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	fmt.Println(cookie)
	if err != nil {
		w.WriteHeader(429)
		return
	}
	token := fmt.Sprint(cookie)
	if token == "nil" {
		w.WriteHeader(500)
		return
	}
	token = token[6:]
	user_id := tokens.CheckToken(token)

	status, contacts := lists.GetUserContactsById(user_id)

	if status == 200 {
		response := Response{
			Status:  200,
			Message: "Успешно получены контакты конкретного пользователя",
			Body:    contacts,
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Fprintf(w, "%s", jsonData)
		return
	}
	if status == 500 {
		w.WriteHeader(500)
		return
	}
}
