package app

import (
	"github.com/sirupsen/logrus"
	"mail/config"
	"mail/daemon"
	"mail/logger"
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
	d := daemon.New(config.FileName())
	return d.Run()
}
