package daemon

import (
	"github.com/BurntSushi/toml"
	"github.com/robfig/cron"
	"io/ioutil"
	"mail/config"
	"mail/converter"
	"os"
	"os/signal"
	"time"
)

func (d Daemon) Run() error {
	c := cron.New()
	schedule, err := cron.ParseStandard(d.config.Scheduler)
	if err != nil {
		return err

	}
	c.Schedule(
		schedule,
		cron.FuncJob(d.process),
	)

	d.logger.Infof("Daemon was initialized, timeout %v, starts settings: %v",
		d.config.Timeout, d.config.Scheduler)

	go c.Start()
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig

	return nil
}

func (d *Daemon) readConfig() error {
	var c config.Config
	if _, err := os.Stat(d.Configfile); err != nil {
		return err
	}
	v, err := ioutil.ReadFile(d.Configfile)
	if err != nil {
		return err
	}
	if err := toml.Unmarshal(v, &c); err != nil {
		return err
	}
	d.config = c.Daemon
	return nil
}

func (d Daemon) process() {
	d.logger.Debugf("Start process daemon...")
	err := d.readConfig()
	if err != nil {
		d.logger.Error(err)
	}

	if err := d.doWithAttempts(0); err != nil {
		d.logger.Error(err)
		return
	}
}

func (d Daemon) doWithAttempts(attempt int) error {

	resp, err := d.clientService.Collect()
	if err != nil {
		if attempt >= 5 {
			return err
		}
		time.Sleep(10 * time.Minute)
		attempt++
		if err := d.doWithAttempts(attempt); err != nil {
			return err
		}
	}
	if resp == nil {
		return nil
	}
	report := converter.Convert(resp)

	if err := report.ToXlsx(); err != nil {
		return err
	}

	return d.sendReport()
}

func (d Daemon) sendReport() error {
	return d.mailService.Send()
}
