package main

import (
	"log"
	"mail/app"
)

func main() {
	runApp()
}

func runApp() {
	a := app.New()
	if err := a.Run(); err != nil {
		log.Println(err)
		return
	}
}
