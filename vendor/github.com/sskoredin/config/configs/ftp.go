package configs

import (
	consul "github.com/sskoredin/config"
)

type FTP struct {
	Host     string `toml:"host" json:"host"`
	Port     int    `toml:"port" json:"port"`
	User     string `toml:"user" json:"user"`
	Password string `toml:"password" json:"password"`
	Path     string `toml:"path" json:"path"`
	File     string `toml:"file" json:"file"`
}

func (c *FTP) Read() error {
	client, err := consul.NewClient()
	if err != nil {
		return err
	}

	c.Host, err = client.Get(consul.FTPHost)
	if err != nil {
		return err
	}
	c.Port, err = client.GetInt(consul.FTPPort)
	if err != nil {
		return err
	}

	c.User, err = client.Get(consul.FTPUser)
	if err != nil {
		return err
	}
	c.Password, err = client.Get(consul.FTPPassword)
	if err != nil {
		return err
	}
	c.Path, err = client.Get(consul.FTPPath)
	if err != nil {
		return err
	}
	c.File, err = client.Get(consul.FTPFile)
	if err != nil {
		return err
	}
	return nil
}
