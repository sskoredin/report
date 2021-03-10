package configs

import (
	consul "github.com/sskoredin/config"
)

type MailOlap struct {
	Host       string   `toml:"host"`
	Port       int      `toml:"port"`
	User       string   `toml:"user"`
	Addressee  string   `toml:"addressee"`
	Password   string   `toml:"password"`
	Subject    string   `toml:"subject"`
	Recipients []string `toml:"recipients"`
}

func (c *MailOlap) Read() error {
	client, err := consul.NewClient()
	if err != nil {
		return err
	}

	c.Host, err = client.Get(consul.MailHost)
	if err != nil {
		return err
	}
	c.Port, err = client.GetInt(consul.MailPort)
	if err != nil {
		return err
	}
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
	c.Subject, err = client.Get(consul.OLAPMailSubject)
	if err != nil {
		return err
	}
	c.Recipients, err = client.GetArray(consul.OLAPMailRecipients)
	if err != nil {
		return err
	}

	return nil
}
