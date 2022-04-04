package main

import (
	"fmt"

	auth "mangadex/auth"
	scanlation "mangadex/scanlation"
	utils "mangadex/utils"
)

// Allowed operations
const (
	LOGIN           string = "LOGIN"
	CHECK_LOGIN            = "CHECK_LOGIN"
	LOGOUT                 = "LOGOUT"
	REFRESH_TOKEN          = "REFRESH_TOKEN"
	SCANLATION_LIST        = "SCANLATION_LIST"
)

func main() {

	user, email, passw := utils.CredentialsInit()
	var bearer_token string
	var refresh_token string

	utils.ShowHeader()
	input := utils.InputHandle("> ")

	for input != "" {

		utils.ClearTerminal()
		fmt.Printf("Operation: %s", input)
		switch input {

		case LOGIN:
			bearer_token, refresh_token = auth.Login(user, email, passw)
		case CHECK_LOGIN:
			auth.CheckToken(bearer_token)
		case REFRESH_TOKEN:
			bearer_token, refresh_token = auth.RefreshToken(refresh_token)
		case LOGOUT:
			auth.Logout(bearer_token)
		case SCANLATION_LIST:
			scanlation.ScanlationList(5, 0, nil, "", "pt-br", nil)
		}

		utils.InputHandle("\nPress ENTER to continue...")
		utils.ShowHeader()
		input = utils.InputHandle("> ")
	}
}
