package config

import consul "github.com/sskoredin/config"

type Client struct {
	User     string `toml:"user" json:"user"`
	Password string `toml:"password" json:"password"`
	API      string `toml:"api" json:"api"`
}

func (c Client) Read() error {
	client, err := consul.NewClient()
	if err != nil {
		return err
	}

	c.User, err = client.Get(consul.IikoLogin)
	if err != nil {
		return err
	}
	c.Password, err = client.Get(consul.IikoPassword)
	if err != nil {
		return err
	}
	c.API, err = client.Get(consul.IikoAPIURl)
	if err != nil {
		return err
	}
	return nil
}
