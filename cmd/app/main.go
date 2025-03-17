package main

import "social/internal/app"

func main() {
	application, err := app.NewApp()
	if err != nil {
		panic(err)
	}

	defer application.Close()

	application.Server.Run()
}
