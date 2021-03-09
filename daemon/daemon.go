package daemon

import (
	"github.com/BurntSushi/toml"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	"github.com/sskoredin/iiko_report/report"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
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

	if err := report.MakeReportWithAttempts("", "", 0); err != nil {
		d.logger.Error(err)
		return
	}
}
