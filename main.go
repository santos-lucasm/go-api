package main

import (
	"bufio"
	"fmt"

	auth "mangadex/auth"
	utils "mangadex/utils"
	"os"
)

const (
	LOGIN         string = "LOGIN"
	CHECK_LOGIN          = "CHECK_LOGIN"
	LOGOUT               = "LOGOUT"
	REFRESH_TOKEN        = "REFRESH_TOKEN"
)

func inputHandle() string {
	fmt.Print("\tEnter operation: ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	return input.Text()
}

func main() {

	user, email, passw := utils.CredentialsInit()
	var bearer_token string
	var refresh_token string

	input := inputHandle()
	for input != "" {

		switch input {

		case LOGIN:
			bearer_token, refresh_token = auth.Login(user, email, passw)
		case CHECK_LOGIN:
			auth.CheckToken(bearer_token)
		case REFRESH_TOKEN:
			bearer_token, refresh_token = auth.RefreshToken(refresh_token)
		case LOGOUT:
			auth.Logout(bearer_token)

		}
		input = inputHandle()
	}
}
