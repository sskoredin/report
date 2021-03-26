package app

import (
	"github.com/sskoredin/iiko_report/daemon"
	"github.com/sskoredin/iiko_report/rest"
	logger "github.com/sskoredin/telegram_client"
	"golang.org/x/sync/errgroup"
)

type App struct {
	logger logger.Logger
}

func New() App {
	return App{
		logger: logger.New("app"),
	}
}

func (a App) Run() error {

	r := rest.New()
	d := daemon.New()
	var g errgroup.Group
	g.Go(r.Run)
	g.Go(d.Run)
	return g.Wait()
}
