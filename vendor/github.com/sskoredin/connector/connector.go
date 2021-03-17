package connector

import (
	ftp_client "github.com/sskoredin/connector/ftp"
	"github.com/sskoredin/connector/iiko"
	period "github.com/sskoredin/connector/period"
)

type IikoClient interface {
	CollectGoodsForFTP(period period.ReportPeriod) (*iiko_client.ResponseGoodsToFTPData, error)
	CollectIikoOLAPReport(period period.ReportPeriod) (*iiko_client.ResponseIikoOlAPReportData, error)
	CollectIikoAmountReport(period period.ReportPeriod) (*iiko_client.ResponseIikoAmountReportData, error)
}

func NewIikoClient() (IikoClient, error) {
	s, err := iiko_client.New()
	if err != nil {
		return nil, err
	}
	return s, nil
}

type FTPClient interface {
	SendGoods(filename string) error
}

func NewFTPClient() (FTPClient, error) {
	s, err := ftp_client.New()
	if err != nil {
		return nil, err
	}
	return s, nil
}
