package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func setProgramPath() {
	currPath := os.Getenv("GOPATH")
	currPathSize := len(currPath) - 1
	currPath = currPath[:currPathSize]
	os.Setenv("PROG_PATH", currPath+"/src/mangadex-api")
}

func loadEnvVariables() {
	err := godotenv.Load(os.Getenv("PROG_PATH") + "/utils/credentials.env")
	HandleError(err)
}

func CredentialsInit() (string, string, string) {
	setProgramPath()
	loadEnvVariables()
	user := os.Getenv("USER")
	email := os.Getenv("EMAIL")
	passw := os.Getenv("PASSWORD")
	return user, email, passw
}
