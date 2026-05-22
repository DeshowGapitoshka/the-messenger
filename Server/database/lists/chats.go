package lists

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Chat struct {
	Id               int        `json:"id"`
	First_person_id  int        `json:"first_person_id"`
	Second_person_id int        `json:"second_person_id"`
	Name             string     `json:"name"`
	LastMessage      *string    `json:"last_message"`
	AvatarLink       *string    `json:"avatar_link"`
	UpdatedAt        *time.Time `json:"updated_at"`
	UnreadCount      *int       `json:"unread_count"`
	Created_at       *time.Time `json:"created_at"`
}

func errorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func times() time.Time {
	return time.Now()
}

var database *sql.DB

func init() {
	var err error
	err = godotenv.Load()
	errorHandler(err)
	connToPG := os.Getenv("DB_CONNECT_TO_BASEDATA")
	if connToPG == "" {
		log.Fatal("connToPG is empty")
	}

	database, err = sql.Open("postgres", connToPG)
	errorHandler(err)
}

func AddNewChat(name string, first_person_id int, second_person_id int) int {
	dbQuery := os.Getenv("DB_CHAT_CREATE")
	if dbQuery == "" {
		log.Println("DB_CHAT_CREATE is empty")
		return 500
	}
	_, err := database.Exec(dbQuery, name, first_person_id, second_person_id)
	if err != nil {
		errorHandler(err)
		return 500
	}
	return 200
}

func GetUserChatsById(userId int) (int, []Chat) {

	dbQuery := os.Getenv("DB_GET_USER_CHATS")
	if dbQuery == "" {
		log.Println("DB_GET_USER_CHATS is empty")
		return 500, nil
	}

	rows, err := database.Query(dbQuery, userId)
	if err != nil {
		log.Printf("Error getting chats: %v", err)
		return 500, nil
	}
	defer rows.Close()

	var chats []Chat

	for rows.Next() {
		var chat Chat
		err := rows.Scan(
			&chat.Id,
			&chat.First_person_id,
			&chat.Second_person_id,
			&chat.Name,
			&chat.LastMessage,
			&chat.AvatarLink,
			&chat.UpdatedAt,
			&chat.UnreadCount,
			&chat.Created_at,
		)
		if err != nil {
			log.Printf("Error scanning chat: %v", err)
			continue
		}
		chats = append(chats, chat)
	}

	return 200, chats
}
