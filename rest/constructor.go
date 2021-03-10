package rest

import (
	"github.com/sskoredin/config/configs"
	logger "github.com/sskoredin/telegram_client"
)

type Rest struct {
	logger *logger.Logger
	config configs.Rest
}

func New() Rest {
	return Rest{
		logger: logger.New("api"),
	}
}
