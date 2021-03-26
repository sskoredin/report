package report

import (
	logger "github.com/sskoredin/telegram_client"
)

type Report struct {
	logger logger.Logger
}

func New() Report {
	return Report{
		logger: logger.New("report"),
	}
}
