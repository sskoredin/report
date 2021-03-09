package daemon

import (
	"github.com/sskoredin/iiko_report/client"
	"github.com/sskoredin/iiko_report/config"
	"github.com/sskoredin/iiko_report/logger"
	"github.com/sskoredin/iiko_report/mail"
)

type Daemon struct {
	config        config.Daemon
	logger        *logger.Logger
	clientService client.Service
	mailService   mail.Service
}

func New() Daemon {
	return Daemon{
		logger:        logger.New("daemon"),
		clientService: client.New(),
		mailService:   mail.New(),
	}
}
