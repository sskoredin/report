package config

import (
	consul "github.com/sskoredin/config"
)

type Config struct {
	client consul.ConsulClient
	Conf
}
type Conf interface {
	Read() error
}
