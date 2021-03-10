package configs

import (
	consul "github.com/sskoredin/config"
)

type DaemonOlap struct {
	Timeout   int    `toml:"timeout" json:"timeout"`
	Scheduler string `toml:"scheduler" json:"scheduler"`
}

func (c *DaemonOlap) Read() error {
	client, err := consul.NewClient()
	if err != nil {
		return err
	}

	c.Timeout = 3600
	c.Scheduler, err = client.Get(consul.OlAPDaemonScheduler)
	if err != nil {
		return err
	}
	return nil
}
