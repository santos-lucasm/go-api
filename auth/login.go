package auth

import (
	"encoding/json"
	"fmt"
	http "mangadex/http"
	utils "mangadex/utils"
)

type loginRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	Result string      `json:"result"`
	Token  tokenLogin  `json:"token"`
	Errors errorsLogin `json:"errors"`
}

type tokenLogin struct {
	Session string `json:"session"`
	Refresh string `json:"refresh"`
}

type errorsLogin struct {
	Id     string `json:"id"`
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func Login(user string, email string, passw string) (string, string) {

	url := utils.RootUrl() + "/auth/login"

	// create POST body
	var body loginRequest
	body.Username = user
	body.Email = email
	body.Password = passw
	json_data, err := json.Marshal(&body)
	utils.HandleError(err)

	response := http.Request("POST", url, json_data, "")

	//fill struct using received bytes
	var responseObject loginResponse
	json.Unmarshal(response, &responseObject)
	fmt.Printf("{LOGIN} result: %s\n", responseObject.Result)
	return "Bearer " + responseObject.Token.Session, responseObject.Token.Refresh
}
