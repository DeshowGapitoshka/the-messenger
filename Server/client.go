package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	id      int
	User_id int
	Data    string
}

func main() {
	var postData Message
	fmt.Print("Введите сообщение: ")
	fmt.Scan(&postData.Data)
	postData.User_id = 666

	// Convert to JSON
	jsonData, err := json.Marshal(postData)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(
		"http://localhost:8050/messages",
		"application/json",
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
