package iiko_client

import (
	"github.com/sskoredin/config/configs"
	logger "github.com/sskoredin/telegram_client"
)

type Service struct {
	token  string
	config configs.Iiko
	logger *logger.Logger
}

func New() (*Service, error) {
	s := &Service{logger: logger.New("iiko_connector")}

	if err := s.config.Read(); err != nil {
		return nil, err
	}
	return s, nil
}
