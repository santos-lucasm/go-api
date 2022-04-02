package auth

import (
	"encoding/json"
	"fmt"
	http "mangadex/http"
	utils "mangadex/utils"
)

type logoutResponse struct {
	Result string `json:"result"`
}

func Logout(token string) {

	url := utils.RootUrl() + "/auth/logout"

	response := http.Request("POST", url, nil, token)

	var responseObject logoutResponse
	json.Unmarshal(response, &responseObject)
	fmt.Printf("{LOGOUT} result: %s\n", responseObject.Result)
}
