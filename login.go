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

type LogoutResponse struct {
	Result string `json:"result"`
}

type RefreshRequest struct {
	Token string `json:"token"`
}

type RefreshResponse struct {
	Result  string     `json:"result"`
	Token   TokenLogin `json:"token"`
	Message string     `json:"message"`
}

func login(user string, email string, passw string) (string, string) {

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

	return "Bearer " + responseObject.Token.Session, responseObject.Token.Refresh

}

func check_login(token string) {

	url := root_url() + "/auth/check"

	response := httpRequest("GET", url, nil, token)

	var responseObject CheckLoginResponse
	json.Unmarshal(response, &responseObject)

	fmt.Println(string(responseObject.Result))
	fmt.Println(responseObject.IsAuthenticated)
}

func logout(token string) {

	url := root_url() + "/auth/logout"

	response := httpRequest("POST", url, nil, token)

	var responseObject LogoutResponse
	json.Unmarshal(response, &responseObject)

	fmt.Println(string(responseObject.Result))
}

func refresh_token(token string) (string, string) {

	url := root_url() + "/auth/refresh"

	var body RefreshRequest
	body.Token = token

	json_data, err := json.Marshal(&body)
	handleError(err)

	response := httpRequest("POST", url, json_data, "")

	var responseObject RefreshResponse
	json.Unmarshal(response, &responseObject)
	fmt.Println(string(responseObject.Result))

	return responseObject.Token.Session, responseObject.Token.Refresh
}
