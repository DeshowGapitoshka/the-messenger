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

type Group struct {
	Id          int       `json:"id"`
	Owner_id    string    `json:"owner_id"`
	Name        string    `json:"name"`
	Info        string    `json:"info"`
	Users_id    []int     `json:"users_id"`
	Admins_id   []int     `json:"admins_id"`
	Enemies_id  []int     `json:"enemies_id"`
	Created_at  string    `json:"created_at"`
	LastMessage string    `json:"last_message"`
	AvatarLink  string    `json:"avatar_link"`
	UpdatedAt   time.Time `json:"updated_at"`
	UnreadCount int       `json:"unread_count"`
}

func CreateNewGroup(w http.ResponseWriter, r *http.Request) {
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

	name := contact["name"].(string)
	info := contact["info"].(string)
	users_id_interface := contact["users_id"].([]interface{})
	users_id := make([]int, len(users_id_interface))
	for i, v := range users_id_interface {
		users_id[i] = int(v.(float64))
	}
	admins_id_interface := contact["admins_id"].([]interface{})
	admins_id := make([]int, len(users_id_interface))
	for i, v := range admins_id_interface {
		admins_id[i] = int(v.(float64))
	}

	result := lists.AddNewGroup(user_id, name, info, users_id, admins_id)

	switch result {
	case 200:
		response := Response{
			Status:  200,
			Message: "Группа успешно создана",
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

func GetUserGroupsById(w http.ResponseWriter, r *http.Request) {
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

	status, groups := lists.GetUserGroupsById(user_id)

	if status == 200 {
		response := Response{
			Status:  200,
			Message: "Успешно получены группы конкретного пользователя",
			Body:    groups,
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
