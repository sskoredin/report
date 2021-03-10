package daemon

import (
	"github.com/sskoredin/config/configs"
	"github.com/sskoredin/iiko_report/client"
	"github.com/sskoredin/iiko_report/mail"
	logger "github.com/sskoredin/telegram_client"
)

type Daemon struct {
	config        configs.DaemonOlap
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
