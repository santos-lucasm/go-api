package main

import (
	"os"

	"github.com/joho/godotenv"
)

func root_url() string {
	return "https://api.mangadex.org"
}

func load_env_variables() {
	err := godotenv.Load("credentials.env")
	handleError(err)
}

func main() {

	load_env_variables()

	user := os.Getenv("USER")
	email := os.Getenv("EMAIL")
	passw := os.Getenv("PASSWORD")

	//TODO: rework this into modules
	bearer_token, ref := login(user, email, passw)

	check_login(bearer_token)

	bearer_token, ref = refresh_token(ref)

	logout(bearer_token)
}
