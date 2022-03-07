package main

import (
	"fmt"
	"os"

	"rsol.com/go-mux/app"
)

func main() {
	fmt.Println("main program")
	os.Setenv("APP_DB_USERNAME", "postgres")
	os.Setenv("APP_DB_PASSWORD", "postgres")
	os.Setenv("APP_DB_NAME", "rsol")

	a := app.App{}

	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	a.Run(":8010")
}
