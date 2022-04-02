package main

import (
	auth "mangadex/auth"
	utils "mangadex/utils"
)

func main() {

	user, email, passw := utils.CredentialsInit()

	bearer_token, ref := auth.Login(user, email, passw)

	auth.CheckToken(bearer_token)

	bearer_token, ref = auth.RefreshToken(ref)

	auth.Logout(bearer_token)
}
