package iiko_biz_client

import (
	"fmt"
	"github.com/sskoredin/config/configs"
	logger "github.com/sskoredin/telegram_client"
	"time"
)

type Service struct {
	token          string
	config         configs.IikoBiz
	tokenCreatedAt *time.Time
	logger         *logger.Logger
}

func New() (*Service, error) {
	s := &Service{
		logger: logger.New("client"),
	}
	if err := s.config.Read(); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Service) link(link string) string {
	return fmt.Sprintf("%s%s", s.config.API, link)
}
func (s *Service) linkWithValue(link, value string) string {
	link = fmt.Sprintf(link, value)
	return s.link(link)
}
