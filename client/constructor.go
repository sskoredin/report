package client

import (
	"github.com/sskoredin/config/configs"
	logger "github.com/sskoredin/telegram_client"
)

type Service struct {
	config configs.Client
	logger *logger.Logger
}

func New() Service {
	return Service{
		logger: logger.New("client"),
	}
}
