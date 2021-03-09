package client

import (
	"github.com/sirupsen/logrus"
	"github.com/sskoredin/iiko_report/config"
	"github.com/sskoredin/iiko_report/logger"
)

type Service struct {
	config config.Client
	logger logger.Logger
}

func New() Service {
	return Service{
		logger: logger.New("client", logrus.DebugLevel),
	}
}
