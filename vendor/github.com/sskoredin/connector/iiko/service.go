package iiko_client

import (
	"encoding/xml"
	_ "encoding/xml"
	"fmt"
	"github.com/sskoredin/connector/domain/entity"
	"github.com/sskoredin/connector/query"
)

const (
	apiUrl        = "/resto/api"
	olapReportUrl = apiUrl + "/reports/olap"
	authUrl       = apiUrl + "/auth"

	reportTypeSales        = "SALES"
	reportTypeTransactions = "TRANSACTIONS"
)

func (s Service) CollectGoodsForFTP(period entity.ReportPeriod) (*ResponseGoodsToFTPData, error) {
	if err := s.config.Read(); err != nil {
		return nil, err
	}
	if err := s.auth(); err != nil {
		return nil, err
	}
	if err := period.Parse(); err != nil {
		return nil, err
	}

	return s.getGoodsToFTP(period)
}

func (s Service) CollectIikoOLAPReport(period entity.ReportPeriod) (*ResponseIikoOlAPReportData, error) {
	if err := s.config.Read(); err != nil {
		return nil, err
	}
	if err := s.auth(); err != nil {
		return nil, err
	}
	if err := period.Parse(); err != nil {
		return nil, err
	}

	return s.iikoOLAPReport(period)
}

func (s Service) getGoodsToFTP(period entity.ReportPeriod) (*ResponseGoodsToFTPData, error) {
	q, err := query.New(s.link(olapReportUrl), s.token).
		ReportType(reportTypeTransactions).
		Period(period).
		GroupRows("Product.Num").
		Args("FinalBalance.Amount").
		Build()
	if err != nil {
		return nil, err
	}

	d, err := q.Get()
	if err != nil {
		return nil, err
	}

	result := new(ResponseGoodsToFTPData)

	if err := xml.Unmarshal(d, result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s Service) CollectIikoAmountReport(period entity.ReportPeriod) (*ResponseIikoAmountReportData, error) {
	if err := s.config.Read(); err != nil {
		return nil, err
	}

	if err := s.auth(); err != nil {
		return nil, err
	}

	if err := period.Parse(); err != nil {
		return nil, err
	}

	return s.iikoAmountReport(period)
}

func (s Service) iikoOLAPReport(period entity.ReportPeriod) (*ResponseIikoOlAPReportData, error) {
	q, err := query.New(s.link(olapReportUrl), s.token).
		ReportType(reportTypeSales).
		Period(period).
		GroupRows("JurName", "DishGroup.SecondParent", "DishCode", "DishName", "PayTypes", "OrderDiscount.Type", "CloseTime", "OrderNum").
		Args("DishAmountInt", "DishSumInt", "DiscountSum", "DishDiscountSumInt", "ProductCostBase.ProductCost").
		Build()

	if err != nil {
		return nil, err
	}

	s.logger.Info(q)
	d, err := q.Get()
	if err != nil {
		return nil, err
	}

	result := new(ResponseIikoOlAPReportData)

	if err := xml.Unmarshal(d, result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s Service) iikoAmountReport(period entity.ReportPeriod) (*ResponseIikoAmountReportData, error) {
	q, err := query.New(s.link(olapReportUrl), s.token).
		ReportType(reportTypeTransactions).
		Period(period).
		GroupRows("Product.SecondParent", "Product.Store", "Product.Num", "Product.Name").
		Args("StartBalance.Amount", "Amount.In", "Amount.Out", "FinalBalance.Amount").
		Build()
	if err != nil {
		return nil, err
	}

	d, err := q.Get()
	if err != nil {
		return nil, err
	}

	result := new(ResponseIikoAmountReportData)

	if err := xml.Unmarshal(d, result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s Service) link(link string) string {
	return fmt.Sprintf("%s%s", s.config.API, link)
}
