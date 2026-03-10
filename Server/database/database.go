package database

import (
	"database/sql"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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

func errorHandler(err error) {
	if err != nil {
		panic(err)
	}
}

func StartServer() {
	var err error
	dbConf := os.Getenv("DB_TABLE_CONFIG")
	dbPersConf := os.Getenv("DB_ACCOUNTS_TABLE_CONFIG")
	if dbConf == "" {
		log.Fatal("dbConf is Empty")
	}
	if dbPersConf == "" {
		log.Fatal("dbPersConf is Empty")
	}
	_, err = database.Exec(dbConf)
	errorHandler(err)
	_, err = database.Exec(dbPersConf)
	errorHandler(err)
}

func InputInBase(user_id int, message string) {
	dbInsert := os.Getenv("DB_INSERT_MESSAGE")
	if dbInsert == "" {
		log.Fatal("dbInsert is empty")
	}
	var err error
	_, err = database.Exec(dbInsert, user_id, message)
	errorHandler(err)
}

func InputInBasePerson(login string, password string) bool {
	var existingLogin string
	dbAccPars := os.Getenv("DB_ACCOUNTS_TABLE_PARSER")
	dbInsertPars := os.Getenv("DB_INSERT_ACCOUNT_TABLE")
	if dbAccPars == "" {
		log.Fatal("dbAccPars is empty")
	}
	if dbInsertPars == "" {
		log.Fatal("dbInsertPers is empty")
	}
	err := database.QueryRow(dbAccPars, login).Scan(&existingLogin)
	if err == sql.ErrNoRows {
		_, err := database.Exec(dbInsertPars, login, password)
		errorHandler(err)
		return true
	}
	return false
}

func CheckLogin(login string, password string) bool {
	var existingLogin Person
	dbAccCheckPars := os.Getenv("DB_ACCOUNT_PASSWORD_PARSER")
	err := database.QueryRow(dbAccCheckPars, login).Scan(&existingLogin.Login, &existingLogin.Password)
	if err == sql.ErrNoRows {
		return false
	}
	if login == existingLogin.Login && password == existingLogin.Password {
		return true
	} else {
		return false
	}
}

func OutputFromBase() []Message {
	dbOutput := os.Getenv("DB_OUTPUT_MESSAGES")
	if dbOutput == "" {
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

func OutputFromBaseID(number string) Message {
	var message Message
	dbOutputID := os.Getenv("DB_OUTPUT_ID_MESSAGE")
	if dbOutputID == "" {
		log.Fatal("dbOutputID is empty")
	}
	id, _ := strconv.Atoi(number)
	row := database.QueryRow(dbOutputID, id)
	row.Scan(&message.Id, &message.User_id, &message.Data)
	return message
}

func OutputFromBaseIdPerson(number string) Person {
	var person Person
	dbOutputIdPerson := os.Getenv("DB_OUTPUT_ID_PERSON")
	if dbOutputIdPerson == "" {
		log.Fatal("dbOutputIdPerson is empty")
	}
	id, _ := strconv.Atoi(number)
	row := database.QueryRow(dbOutputIdPerson, id)
	row.Scan(&person.Id, &person.Login, &person.Password)
	return person
}
