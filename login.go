package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type loginResponse struct {
	Result string     `json:"result"`
	Token  TokenLogin `json:"token"`
}

type TokenLogin struct {
	Session string `json:"session"`
	Refresh string `json:"refresh"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func login() {

	url := root_url() + "/auth/login"

	// data := map[string]string{"username": "a", "email": "a", "password": "a"}
	var body LoginRequest
	body.Username = "lucas"
	body.Email = "santos.lucasmmatheus@gmail.com"
	body.Password = "Ftypw"

	json_data, err := json.Marshal(&body)
	if err != nil {
		log.Fatal(err)
	}

	response := httpPost(url, json_data)
	responseData := processResponse(response)

	var responseObject loginResponse
	json.Unmarshal(responseData, &responseObject)
	fmt.Println(string(responseObject.Result))
	fmt.Println(string(responseObject.Token.Session))
	fmt.Println(string(responseObject.Token.Refresh))

	// fmt.Println(loginProcess())
}
