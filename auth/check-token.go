package auth

import (
	"encoding/json"
	"fmt"
	http "mangadex/http"
	utils "mangadex/utils"
)

type checkTokenResponse struct {
	Result          string `json:"result"`
	IsAuthenticated bool   `json:"isAuthenticated"`
}

func CheckToken(token string) {

	url := utils.RootUrl() + "/auth/check"

	response := http.Request("GET", url, nil, token)

	// fill struct using received bytes
	var responseObject checkTokenResponse
	json.Unmarshal(response, &responseObject)
	fmt.Printf("{CHECK-TOKEN} result: %s\n", responseObject.Result)
	fmt.Printf("{CHECK-TOKEN} isAuthenticated: %t\n", responseObject.IsAuthenticated)
}
