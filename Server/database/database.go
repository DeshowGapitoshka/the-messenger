package database

import (
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
	"log"
)

type Message struct {
	Id      int
	User_id int
	Data    string
}

var messages []Message

var database *sql.DB

func init() {
	var err error
	err = godotenv.Load()
	errorHandler(err)
	connToPG := os.Getenv("DB_CONNTOPG")
	if connToPG == ""{
		log.Fatal("connToPG is empty")
	}

	database, err = sql.Open("postgres", connToPG)
	errorHandler(err)
}

func errorHandler(err error) {
	if err != nil {
		panic(err)
	}
}

func StartServer() {
	var err error
	dbConf := os.Getenv("DB_TABLECONF")
	if dbConf == ""{
		log.Fatal("dbConf is Empty")
	}
	_, err = database.Exec(dbConf)
	errorHandler(err)
}

func InputInBase(user_id int, message string) {
	dbInsert := os.Getenv("DB_INSERTMES")
	if dbInsert == ""{
		log.Fatal("dbInsert is empty")
	}
	var err error
	_, err = database.Exec(dbInsert, user_id, message)
	errorHandler(err)
}

func OutputFromBase() []Message {
	dbOutput := os.Getenv("DB_OUTPUTMES")
	if dbOutput == ""{
		log.Fatal("dbOutput is empty")
	}
	rows, err := database.Query(dbOutput)
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
