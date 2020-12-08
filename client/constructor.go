package client

import (
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"mail/config"
	"mail/logger"
	"os"
)

type Service struct {
	Configfile string
	config     config.Client
	logger     logger.Logger
}

func New(configfile string) Service {
	return Service{
		Configfile: configfile,
		logger:     logger.New("client", logrus.DebugLevel),
	}
}

func (s *Service) readConfig() error {
	var c config.Config
	if _, err := os.Stat(s.Configfile); err != nil {
		return err
	}
	v, err := ioutil.ReadFile(s.Configfile)
	if err != nil {
		return err
	}
	if err := toml.Unmarshal(v, &c); err != nil {
		return err
	}
	s.config = c.Client
	return nil
}
