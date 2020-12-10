package config

import "fmt"

type Database struct {
	File string `toml:"file"`
}

type Rest struct {
	Host string `toml:"host" json:"host"`
	Port int    `toml:"port" json:"port"`
}

func (r Rest) ListenAddr() string {
	return fmt.Sprintf("%s:%d", r.Host, r.Port)
}

type Config struct {
	Mail   Mail   `toml:"mail" json:"mail" `
	Client Client `toml:"client" json:"client" `
	Rest   Rest   `toml:"rest" json:"rest" `
	Daemon Daemon `toml:"daemon" json:"daemon"`
}

type Client struct {
	User     string `toml:"user" json:"user"`
	Password string `toml:"password" json:"password"`
	API      string `toml:"api" json:"api"`
}

type Mail struct {
	Host       string   `toml:"host"`
	Port       int      `toml:"port"`
	User       string   `toml:"user"`
	Addressee  string   `toml:"addressee"`
	Password   string   `toml:"password"`
	Subject    string   `toml:"subject"`
	Recipients []string `toml:"recipients"`
}

type Daemon struct {
	Timeout   int    `toml:"timeout" json:"timeout"`
	Scheduler string `toml:"scheduler" json:"scheduler"`
}
