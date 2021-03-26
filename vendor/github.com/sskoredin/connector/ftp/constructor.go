package ftp_client

import (
	"github.com/sskoredin/config/configs"
	logger "github.com/sskoredin/telegram_client"
)

type Service struct {
	config configs.FTP
	logger logger.Logger
}

func New() (*Service, error) {
	s := &Service{
		logger: logger.New("ftp"),
	}
	if err := s.config.Read(); err != nil {
		return nil, err
	}

	return s, nil
}
