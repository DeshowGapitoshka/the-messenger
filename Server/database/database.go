package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Message struct {
	Id      int
	User_id int
	Data    string
}

var messages []Message

const (
	connToPG = "user=Deshow password=12345 dbname=messanger sslmode=disable"
)

var database *sql.DB

func init() {
	var err error
	database, err = sql.Open("postgres", connToPG)
	if err != nil {
		panic(err)
	}
}

func errorHandler(err error) {
	if err != nil {
		panic(err)
	}
}

func StartServer() {
	_, err := database.Exec("CREATE TABLE IF NOT EXISTS messager (id SERIAL PRIMARY KEY, user_id INTEGER NOT NULL, message TEXT)")
	errorHandler(err)
}

func InputInBase(user_id int, message string) {
	_, err := database.Exec("INSERT INTO messager (user_id, message) VALUES ($1, $2)", user_id, message)
	errorHandler(err)
}

func OutputFromBase() []Message {
	rows, err := database.Query("SELECT id, user_id, message FROM messager")
	errorHandler(err)
	defer rows.Close()
	var messages []Message
	for rows.Next() {
		var message Message
		err = rows.Scan(&message.Id, &message.User_id, &message.Data)
		errorHandler(err)
		messages = append(messages, message)
	}
	return messages
}
