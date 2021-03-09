package config

import (
	"fmt"
	consul "github.com/sskoredin/config"
	"strconv"
	"strings"
)

type Config struct {
	client config.ConsulClient
}

type Rest struct {
	Host string `toml:"host" json:"host"`
	Port int    `toml:"port" json:"port"`
}

func (r Rest) ListenAddr() string {
	return fmt.Sprintf("%s:%d", r.Host, r.Port)
}
func GetRest() (*Rest, error) {
	client, err := consul.NewClient()
	if err != nil {
		return nil, err
	}
	var c Rest

}

type Client struct {
	User     string `toml:"user" json:"user"`
	Password string `toml:"password" json:"password"`
	API      string `toml:"api" json:"api"`
}

type MailConfig struct {
	Host       string   `toml:"host"`
	Port       int      `toml:"port"`
	User       string   `toml:"user"`
	Addressee  string   `toml:"addressee"`
	Password   string   `toml:"password"`
	Subject    string   `toml:"subject"`
	Recipients []string `toml:"recipients"`
}

func GetMail() (*MailConfig, error) {
	client, err := consul.NewClient()
	if err != nil {
		return nil, err
	}

	var c MailConfig
	c.Host, err = client.Get(consul.MailHost)
	if err != nil {
		return nil, err
	}
	port, err := client.Get(consul.MailPort)
	if err != nil {
		return nil, err
	}
	c.Port, _ = strconv.Atoi(port)

	c.User, err = client.Get(consul.MailLogin)
	if err != nil {
		return nil, err
	}
	c.Addressee, err = client.Get(consul.MailAddressee)
	if err != nil {
		return nil, err
	}
	c.Password, err = client.Get(consul.MailPassword)
	if err != nil {
		return nil, err
	}
	c.Subject, err = client.Get(consul.MailOLAPSubject)
	if err != nil {
		return nil, err
	}
	recipients, err := client.Get(consul.MailOLAPRecipients)
	if err != nil {
		return nil, err
	}
	recipients = strings.Trim(recipients, "[")
	recipients = strings.Trim(recipients, "]")
	c.Recipients = strings.Split(recipients, ",")

	return &c, nil
}

type Daemon struct {
	Timeout   int    `toml:"timeout" json:"timeout"`
	Scheduler string `toml:"scheduler" json:"scheduler"`
}
