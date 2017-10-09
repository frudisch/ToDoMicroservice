package main

import (
	"GoBlogEntry/app"
	"os"
)

func main() {
	os.Setenv("APP_DB_USERNAME", "go_user")
	os.Setenv("APP_DB_PASSWORD", "go_user_passwd")
	os.Setenv("APP_DB_NAME", "todo")

	a := app.App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	a.Run(":8080")
}
