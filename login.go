package main

import (
	"encoding/json"
	"fmt"
)

type LoginResponse struct {
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

type CheckLoginResponse struct {
	Result          string `json:"result"`
	IsAuthenticated bool   `json:"isAuthenticated"`
}

func login(user string, email string, passw string) {

	url := root_url() + "/auth/login"

	var body LoginRequest
	body.Username = user
	body.Email = email
	body.Password = passw

	json_data, err := json.Marshal(&body)
	handleError(err)

	response := httpRequest("POST", url, json_data, "")
	var responseObject LoginResponse
	json.Unmarshal(response, &responseObject)
	fmt.Println(string(responseObject.Result))

	// ------------------------------------------------------------------

	bearer_token := "Bearer " + string(responseObject.Token.Session)
	check_url := root_url() + "/auth/check"
	check_response := httpRequest("GET", check_url, nil, bearer_token)
	var checkResponseObject CheckLoginResponse
	json.Unmarshal(check_response, &checkResponseObject)
	fmt.Println(string(checkResponseObject.Result))
	fmt.Println(checkResponseObject.IsAuthenticated)
}
