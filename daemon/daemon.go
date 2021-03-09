package daemon

import (
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	"github.com/sskoredin/iiko_report/report"
	"os"
	"os/signal"
)

func (d Daemon) Run() error {
	c := cron.New()
	if err := d.config.Read(); err != nil {
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

func (d Daemon) process() {
	d.logger.Debug("Start process daemon...")

	if err := report.MakeReportWithAttempts("", "", 0); err != nil {
		d.logger.Error(err)
		return
	}
}
