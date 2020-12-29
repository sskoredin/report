package app

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"mail/config"
	"mail/daemon"
	"mail/logger"
	"mail/report"
	"mail/rest"
)

type App struct {
	logger logger.Logger
}

func New() App {
	return App{
		logger: logger.New("app", logrus.DebugLevel),
	}
}

func (a App) Run() error {
	if err := config.Check(); err != nil {
		a.logger.Error(err)
		return err
	}
	r := rest.New()
	d := daemon.New(config.FileName())
	var g errgroup.Group
	g.Go(r.Run)
	g.Go(d.Run)
	return g.Wait()
}

func (a App) Send() error {
	if err := config.Check(); err != nil {
		a.logger.Error(err)
		return err
	}
	if err := report.MakeReportWithAttempts("", "", 0); err != nil {
		a.logger.Error(err)
		return err
	}
	return nil
}
