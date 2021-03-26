package mail

import (
	"github.com/sskoredin/config/configs"
	logger "github.com/sskoredin/telegram_client"
)

type Service struct {
	Configfile string
	config     configs.MailOlap
	logger     logger.Logger
}

func New() Service {
	return Service{
		logger: logger.New("mail"),
	}
}
