package lists

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Contact struct {
	Id         int    `json:"id"`
	User_id    int    `json:"user_id"`
	Contact_id int    `json:"contact_id"`
	Created_at string `json:"created_at"`
	Status     string `json:"status"`
}

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

func AddNewContact(userId int, contactId int) int {
	dbQuery := os.Getenv("DB_ADD_CONTACT")
	if dbQuery == "" {
		log.Println("DB_ADD_CONTACT is empty")
		return 500
	}
	_, err := database.Exec(dbQuery, userId, contactId)
	if err != nil {
		errorHandler(err)
		return 500
	}
	return 200
}

func GetUserContactsById(userId int) (int, []Contact) {

	dbQuery := os.Getenv("DB_GET_USER_CONTACTS")
	if dbQuery == "" {
		log.Println("DB_GET_USER_CONTACTS is empty")
		return 500, nil
	}

	rows, err := database.Query(dbQuery, userId)
	if err != nil {
		log.Printf("Error getting chats: %v", err)
		return 500, nil
	}
	defer rows.Close()

	var contacts []Contact

	for rows.Next() {
		var contact Contact
		err := rows.Scan(
			&contact.Id,
			&contact.User_id,
			&contact.Contact_id,
			&contact.Created_at,
			&contact.Status,
		)
		if err != nil {
			log.Printf("Error scanning chat: %v", err)
			continue
		}
		contacts = append(contacts, contact)
	}

	return 200, contacts
}
