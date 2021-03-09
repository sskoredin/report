package config

import (
	"fmt"
	consul "github.com/sskoredin/config"
	"strconv"
)

type Rest struct {
	Host string `toml:"host" json:"host"`
	Port int    `toml:"port" json:"port"`
}

func (r Rest) ListenAddr() string {
	return fmt.Sprintf("%s:%d", r.Host, r.Port)
}
func (r Rest) Read() error {
	client, err := consul.NewClient()
	if err != nil {
		return err
	}
	r.Host, err = client.Get(consul.OLAPRestHost)
	if err != nil {
		return err
	}
	port, err := client.Get(consul.OLAPRestPort)
	p, err := strconv.Atoi(port)
	if err != nil {
		return err
	}
	r.Port = p
	return nil
}
