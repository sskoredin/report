package daemon

import (
	"github.com/sirupsen/logrus"
	"github.com/sskoredin/iiko_report/client"
	"github.com/sskoredin/iiko_report/logger"
	"github.com/sskoredin/iiko_report/mail"
)

type Daemon struct {
	Configfile    string
	config        config.Daemon
	logger        logger.Logger
	clientService client.Service
	mailService   mail.Service
}

func New(configfile string) Daemon {
	return Daemon{
		Configfile:    configfile,
		logger:        logger.New("daemon", logrus.DebugLevel),
		clientService: client.New(),
		mailService:   mail.New(),
	}
}
