package config

import (
	consul "github.com/sskoredin/config"
	"strconv"
	"strings"
)

type Mail struct {
	Host       string   `toml:"host"`
	Port       int      `toml:"port"`
	User       string   `toml:"user"`
	Addressee  string   `toml:"addressee"`
	Password   string   `toml:"password"`
	Subject    string   `toml:"subject"`
	Recipients []string `toml:"recipients"`
}

func (c *Mail) Read() error {
	client, err := consul.NewClient()
	if err != nil {
		return err
	}

	c.Host, err = client.Get(consul.MailHost)
	if err != nil {
		return err
	}
	port, err := client.Get(consul.MailPort)
	if err != nil {
		return err
	}
	c.Port, _ = strconv.Atoi(port)

	c.User, err = client.Get(consul.MailLogin)
	if err != nil {
		return err
	}
	c.Addressee, err = client.Get(consul.MailAddressee)
	if err != nil {
		return err
	}
	c.Password, err = client.Get(consul.MailPassword)
	if err != nil {
		return err
	}
	c.Subject, err = client.Get(consul.MailOLAPSubject)
	if err != nil {
		return err
	}
	recipients, err := client.Get(consul.MailOLAPRecipients)
	if err != nil {
		return err
	}
	recipients = strings.Trim(recipients, "[")
	recipients = strings.Trim(recipients, "]")
	c.Recipients = strings.Split(recipients, ",")

	return nil
}
