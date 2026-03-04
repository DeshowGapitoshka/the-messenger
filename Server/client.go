package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Data string
	Id   int
}

func main() {
	var postData Message
	fmt.Print("Введите сообщение: ")
	fmt.Scan(&postData.Data)
	postData.Id = 0

	// Convert to JSON
	jsonData, err := json.Marshal(postData)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(
		"http://localhost:8080/messages",
		"application/json",
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
