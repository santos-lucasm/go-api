package auth

import (
	"encoding/json"
	"fmt"
	http "mangadex/http"
	utils "mangadex/utils"
)

type refreshRequest struct {
	Token string `json:"token"`
}

type refreshResponse struct {
	Result  string       `json:"result"`
	Token   tokenRefresh `json:"token"`
	Message string       `json:"message"`
}

type tokenRefresh struct {
	Session string `json:"session"`
	Refresh string `json:"refresh"`
}

func RefreshToken(token string) (string, string) {

	url := utils.RootUrl() + "/auth/refresh"

	var body refreshRequest
	body.Token = token

	json_data, err := json.Marshal(&body)
	utils.HandleError(err)

	response := http.Request("POST", url, json_data, "")

	var responseObject refreshResponse
	json.Unmarshal(response, &responseObject)
	fmt.Printf("{REFRESH-TOKEN} result: %s\n", responseObject.Result)
	fmt.Printf("{REFRESH-TOKEN} result: %s\n", responseObject.Message)

	return responseObject.Token.Session, responseObject.Token.Refresh
}
