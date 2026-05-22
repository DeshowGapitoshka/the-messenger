package list

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	lists "server/database/lists"
	tokens "server/database/tokens"
	"time"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Body    interface{} `json:"body"`
}

type Chat struct {
	Id               int       `json:"id"`
	First_person_id  int       `json:"first_person_id"`
	Second_person_id int       `json:"second_person_id"`
	Name             string    `json:"name"`
	LastMessage      string    `json:"last_message"`
	AvatarLink       string    `json:"avatar_link"`
	UpdatedAt        time.Time `json:"updated_at"`
	UnreadCount      int       `json:"unread_count"`
}

func CreateNewChat(w http.ResponseWriter, r *http.Request) {
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

	chat := make(map[string]interface{})
	err = json.NewDecoder(r.Body).Decode(&chat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	name := chat["name"].(string)
	second_person_id := chat["second_person_id"].(float64)

	result := lists.AddNewChat(name, user_id, int(second_person_id))

	switch result {
	case 200:
		response := Response{
			Status:  200,
			Message: "Чат успешно создан",
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

func GetUserChatsById(w http.ResponseWriter, r *http.Request) {
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

	status, chats := lists.GetUserChatsById(user_id)

	if status == 200 {
		response := Response{
			Status:  200,
			Message: "Успешно получены чаты конкретного пользователя",
			Body:    chats,
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
