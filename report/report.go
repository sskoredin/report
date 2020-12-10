package report

import (
	"log"
	"mail/client"
	"mail/converter"
	"mail/mail"
	"time"
)

func MakeReportWithAttempts(start, end string, attempt int) error {
	log.Println("start process", attempt)
	return makeReport(start, end, attempt)
}

func makeReport(start, end string, attempt int) error {
	cli := client.New()
	resp, err := cli.Collect(start, end)
	if err != nil {
		if attempt >= 5 {
			log.Println("max attempts")
			return err
		}
		time.Sleep(10 * time.Minute)
		attempt++
		if err := MakeReportWithAttempts(start, end, attempt); err != nil {
			return err
		}
	}
	if resp == nil {
		log.Println("Empty response")
		return nil
	}
	report := converter.Convert(resp)
	log.Println("response converted ")
	if err := report.ToXlsx(start, end); err != nil {
		return err
	}
	log.Println("report prepared")

	return sendReport(start, end)
}
func sendReport(start, end string) error {
	return mail.New().Send(start, end)
}
