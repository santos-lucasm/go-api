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

	// fmt.Println(user)
	login(user, email, passw)
}
