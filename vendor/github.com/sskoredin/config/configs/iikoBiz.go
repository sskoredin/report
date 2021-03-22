package configs

import (
	consul "github.com/sskoredin/config"
)

type IikoBiz struct {
	User         string `toml:"user" json:"user"`
	Password     string `toml:"password" json:"password"`
	API          string `toml:"api" json:"api"`
	CategoryName string `toml:"category_name" json:"category_name"`
}

func (c *IikoBiz) Read() error {
	client, err := consul.NewClient()
	if err != nil {
		return err
	}

	c.User, err = client.Get(consul.IikoBizLogin)
	if err != nil {
		return err
	}
	c.Password, err = client.Get(consul.IikoBizPassword)
	if err != nil {
		return err
	}
	c.API, err = client.Get(consul.IikoBizAPIURl)
	if err != nil {
		return err
	}
	c.CategoryName, err = client.Get(consul.IikoBizCategoryName)
	if err != nil {
		return err
	}
	return nil
}
