package main

import (
	"github.com/sskoredin/iiko_report/app"
	logger "github.com/sskoredin/telegram_client"
)

func main() {
	runApp()
}

func runApp() {
	a := app.New()
	l := logger.New("main")
	if err := a.Run(); err != nil {
		l.Error(err.Error())
		return
	}
}
