package mail

import (
	"errors"
	"fmt"
	"github.com/sskoredin/iiko_report/converter"
	mv2 "gopkg.in/mail.v2"
	"os"
)

func (s Service) Send(start, end string) error {
	err := s.config.Read()
	if err != nil {
		return err
	}
	s.logger.Debugf("config:%+v", s.config)
	return s.send(start, end)
}

func (s Service) send(start, end string) error {
	s.logger.Info("sending mail")

	filename := converter.ReportName(start, end)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return errors.New("Have no attachment ")
	}

	d := mv2.NewDialer(s.config.Host, s.config.Port, s.config.User, s.config.Password)
	d.StartTLSPolicy = mv2.MandatoryStartTLS
	sm, err := d.Dial()
	if err != nil {
		return err
	}

	for _, recipient := range s.config.Recipients {
		m := mv2.NewMessage()
		m.SetAddressHeader("From", s.config.User, "robot")
		m.SetHeader("To", recipient)
		m.SetHeader("Subject", fmt.Sprintf("%s %s", s.config.Subject, end))
		m.Attach(filename)

		m.SetBody("text/html", fmt.Sprintf("OLAP report %s", end))
		if err := mv2.Send(sm, m); err != nil {
			return err
		}
		s.logger.Infof("send to %s successfully", recipient)
	}

	return nil
}
