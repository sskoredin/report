package client

import (
	"github.com/sskoredin/iiko_report/config"
	logger "github.com/sskoredin/telegram_client"
)

type Service struct {
	config config.Client
	logger *logger.Logger
}

func New() Service {
	return Service{
		logger: logger.New("client"),
	}
}
