package configs

import (
	consul "github.com/sskoredin/config"
)

type DaemonAmount struct {
	Timeout   int    `toml:"timeout" json:"timeout"`
	Scheduler string `toml:"scheduler" json:"scheduler"`
}

func (c *DaemonAmount) Read() error {
	client, err := consul.NewClient()
	if err != nil {
		return err
	}

	c.Timeout = 3600
	c.Scheduler, err = client.Get(consul.AmountDaemonScheduler)
	if err != nil {
		return err
	}
	return nil
}
