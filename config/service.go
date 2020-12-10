package config

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
)

func Check() error {
	if _, err := os.Stat(FileName()); err != nil {
		if os.IsNotExist(err) {
			return errors.New(fmt.Sprintf("config %s file not exist", FileName()))
		}
		return err
	}
	f, err := ioutil.ReadFile(FileName())
	if err != nil {
		return err
	}
	var conf Config
	if err := toml.Unmarshal(f, &conf); err != nil {
		return err
	}

	fmt.Println(conf)

	return nil
}

func FileName() string {
	return "C:\\Program Files\\iikoReporter\\config.toml"
}
