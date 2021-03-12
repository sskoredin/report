package main

import (
	"github.com/sskoredin/iiko_report/app"
	logger "github.com/sskoredin/telegram_client"
	"log"
	"time"
	_ "time/tzdata"
)

func init() {
	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		log.Fatal(err)
	}
	time.Local = loc // -> this is setting the global timezone
}

func main() {
	log := logger.New("main")
	log.Infof("Start iiko report at %v", time.Now().String())

	if err := runApp(); err != nil {
		log.Fatal(err)
	}
}

func runApp() error {
	a := app.New()
	return a.Run()
}
