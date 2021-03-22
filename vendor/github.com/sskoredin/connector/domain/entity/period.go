package entity

import (
	"github.com/jinzhu/now"
	"time"
)

const DateLayout = "2006-01-02"

type ReportPeriod struct {
	Start string
	End   string
}

func (p *ReportPeriod) Parse() error {
	st, err := start(p.Start)
	if err != nil {
		return err
	}
	e, err := end(p.End)
	if err != nil {
		return err
	}

	p.Start = st
	p.End = e

	return nil
}

func start(st string) (string, error) {
	start := now.BeginningOfDay()
	start = start.AddDate(0, 0, -1)
	if len(st) > 0 {
		v, err := time.Parse("02.01.2006", st)
		if err != nil {
			return "", err
		}
		start = v
	}
	return start.Format("02.01.2006"), nil
}

func end(e string) (string, error) {
	end := now.BeginningOfDay()
	if len(e) > 0 {
		v, err := time.Parse("02.01.2006", e)
		if err != nil {
			return "", err
		}
		end = v
	}

	return end.Format("02.01.2006"), nil
}
