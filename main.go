package main

import (
	"GoBlogEntry/app"
	"os"
)

func main() {
	a := app.App{}
	a.Initialize(
		os.Getenv("DB_CONNECTION"),
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	a.Run(":8080")
}
