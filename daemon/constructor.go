package daemon

import (
	"github.com/sskoredin/config/configs"
	"github.com/sskoredin/iiko_report/mail"
	logger "github.com/sskoredin/telegram_client"
)

type Daemon struct {
	config      configs.DaemonOlap
	logger      logger.Logger
	mailService mail.Service
}

func New() Daemon {
	return Daemon{
		logger:      logger.New("daemon"),
		mailService: mail.New(),
	}
}
