package main

import (
	"github.com/sskoredin/iiko_report/app"
	"log"
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
