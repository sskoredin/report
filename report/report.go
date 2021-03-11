package report

import (
	"github.com/sskoredin/iiko_report/client"
	"github.com/sskoredin/iiko_report/converter"
	"github.com/sskoredin/iiko_report/mail"
	"time"
)

func MakeReportWithAttempts(start, end string, attempt int) error {
	r := New()
	r.logger.Info("start process", attempt)
	return r.makeReport(start, end, attempt)
}

func (r Report) makeReport(start, end string, attempt int) error {
	cli := client.New()

	resp, err := cli.Collect(start, end)
	if err != nil {
		if attempt >= 5 {
			r.logger.Info("max attempts")
			return err
		}
		time.Sleep(10 * time.Minute)
		attempt++
		if err := MakeReportWithAttempts(start, end, attempt); err != nil {
			r.logger.Error(err)
			return err
		}
	}

	if resp == nil {
		r.logger.Warn("Empty response")
		return nil
	}

	report := converter.Convert(resp)

	r.logger.Debug("response converted ")

	if err := report.ToXlsx(start, end); err != nil {
		r.logger.Error(err)
		return err
	}
	r.logger.Debug("report prepared")

	return sendReport(start, end)
}
func sendReport(start, end string) error {
	return mail.New().Send(start, end)
}
