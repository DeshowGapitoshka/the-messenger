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

type Person struct {
	Login string
	Password string
}

func main() {
	var version int
	fmt.Println("1 - POST_Message       2 - GET_Message")
	fmt.Println("3 - POST_Account       4 - GET_Account")
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

	if version == 3{
		var postData Person
		fmt.Print("Введите логин: ")
		fmt.Scan(&postData.Login)
		fmt.Print("Введите пароль(цифры)")
		fmt.Scan(&postData.Password)

		// Convert to JSON
		jsonData, err := json.Marshal(postData)
		if err != nil {
			panic(err)
		}

		resp, err := http.Post(
			"http://localhost:8050/register",
			"application/json",
			bytes.NewBuffer(jsonData),
		)

		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
	}

	if version == 4{
		eqq := "http://localhost:8050/getaccount?id=" + strconv.Itoa(4)
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
