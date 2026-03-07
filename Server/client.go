package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"io"
)

type Message struct {
	id      int
	User_id int
	Data    string
}

func main() {
	var version int
	fmt.Println("1 - POST       2 - GET")
	fmt.Scan(&version)
	if version == 1{
		var postData Message
		fmt.Print("Введите сообщение: ")
		fmt.Scan(&postData.Data)
		fmt.Print("Введите ваш айди")
		fmt.Scan(&postData.User_id)

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

	if version == 2{
		eqq := "http://localhost:8050/getmessage?id=" + strconv.Itoa(1)
		req,err := http.Get(eqq)
		defer req.Body.Close()
		if err != nil{
			panic(err)
		}
		body, err := io.ReadAll(req.Body)
		if err != nil{
			panic(err)
		}
		fmt.Println(string(body))
	}
}
