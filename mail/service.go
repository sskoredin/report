package mail

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	mv2 "gopkg.in/mail.v2"
	"io/ioutil"
	"mail/config"
	"mail/converter"
	"os"
	"time"
)

type Service struct {
	Configfile string
	config     config.Mail
}

func New(configFile string) Service {
	return Service{
		Configfile: configFile,
	}
}

func (s Service) Send() error {
	err := s.readConfig()
	if err != nil {
		return err
	}
	return s.send()
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
	s.config = c.Mail
	return nil
}

func (s Service) send() error {
	filename := converter.ReportName()
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		return errors.New("Have no attachment ")
	}
	d := mv2.NewDialer(s.config.Host, s.config.Port, s.config.User, s.config.Password)
	d.StartTLSPolicy = mv2.MandatoryStartTLS
	sm, err := d.Dial()
	if err != nil {
		return err
	}
	m := mv2.NewMessage()
	for _, recipient := range s.config.Recipients {
		m.SetAddressHeader("From", s.config.User, "robot")
		m.SetHeader("To", recipient)
		m.SetHeader("Subject", s.config.Subject)
		m.Attach(filename)

		m.SetBody("text/html", fmt.Sprintf("OLAP report %s", time.Now().Format("02.01.2006 15:04:05")))
		if err := mv2.Send(sm, m); err != nil {
			return err
		}
		logrus.Infof("send to %s successfully", recipient)
	}
	return nil
}
