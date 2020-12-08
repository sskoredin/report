package config

import (
	"errors"
	"fmt"
	"os"
	"path"
)

const file = "config.toml"

func Check() error {
	p, err := os.Getwd()
	if err != nil {
		return err
	}
	configFile := path.Join(p, "config.toml")
	if _, err := os.Stat(configFile); err != nil {
		if os.IsNotExist(err) {
			return errors.New(fmt.Sprintf("config %s file not exist", configFile))
		}
		return err
	}
	return nil
}

func FileName() string {
	p, _ := os.Getwd()

	return path.Join(p, file)
}
