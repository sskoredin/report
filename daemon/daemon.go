package daemon

import (
	"github.com/BurntSushi/toml"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"mail/config"
	"mail/converter"
	"os"
	"os/signal"
	"time"
)

func (d Daemon) Run() error {
	c := cron.New()
	if err := d.readConfig(); err != nil {
		return err
	}
	logrus.Println(d.config)
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
	log.Println("process")
	d.logger.Debugf("Start process daemon...")

	if err := d.doWithAttempts(0); err != nil {
		d.logger.Error(err)
		return
	}
}

func (d Daemon) doWithAttempts(attempt int) error {
	log.Println("start process", attempt)
	resp, err := d.clientService.Collect()
	if err != nil {
		if attempt >= 5 {
			log.Println("max attempts")
			return err
		}
		time.Sleep(10 * time.Minute)
		attempt++
		if err := d.doWithAttempts(attempt); err != nil {
			return err
		}
	}
	if resp == nil {
		d.logger.Warn("Empty response")
		log.Println("Empty response")
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
