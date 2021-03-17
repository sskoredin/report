package iiko_client

import (
	"encoding/xml"
	_ "encoding/xml"
	"fmt"
	"github.com/sskoredin/connector/iiko/query"
	iikoclient "github.com/sskoredin/connector/period"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	olapReportUrl          = "/resto/api/reports/olap"
	authUrl                = "/resto/api/auth"
	errNotAuth             = "not authorized"
	reportTypeSales        = "SALES"
	reportTypeTransactions = "TRANSACTIONS"
)

func (s Service) CollectGoodsForFTP(period iikoclient.ReportPeriod) (*ResponseGoodsToFTPData, error) {
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

func (s Service) CollectIikoOLAPReport(period iikoclient.ReportPeriod) (*ResponseIikoOlAPReportData, error) {
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

func (s Service) getGoodsToFTP(period iikoclient.ReportPeriod) (*ResponseGoodsToFTPData, error) {
	params := query.New(s.token).
		ReportType(reportTypeTransactions).
		Period(period).
		GroupRows("Product.Num").
		Args("FinalBalance.Amount")

	q, err := s.queryBuilder(olapReportUrl, params)
	if err != nil {
		return nil, err
	}
	d, err := get(q)
	if err != nil {
		return nil, err
	}

	result := new(ResponseGoodsToFTPData)

	if err := xml.Unmarshal(d, result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s Service) CollectIikoAmountReport(period iikoclient.ReportPeriod) (*ResponseIikoAmountReportData, error) {
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

func (s Service) iikoOLAPReport(period iikoclient.ReportPeriod) (*ResponseIikoOlAPReportData, error) {
	params := query.New(s.token).
		ReportType(reportTypeSales).
		Period(period).
		GroupRows("JurName", "DishGroup.SecondParent", "DishCode", "DishName", "PayTypes", "OrderDiscount.Type").
		Args("DishSumInt", "DiscountSum", "DishDiscountSumInt", "ProductCostBase.ProductCost")

	q, err := s.queryBuilder(olapReportUrl, params)
	if err != nil {
		return nil, err
	}
	d, err := get(q)
	if err != nil {
		return nil, err
	}

	result := new(ResponseIikoOlAPReportData)

	if err := xml.Unmarshal(d, result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s Service) iikoAmountReport(period iikoclient.ReportPeriod) (*ResponseIikoAmountReportData, error) {
	params := query.New(s.token).
		ReportType(reportTypeTransactions).
		Period(period).
		GroupRows("Product.SecondParent", "Product.Store", "Product.Num", "Product.Name").
		Args("StartBalance.Amount", "Amount.In", "Amount.Out", "FinalBalance.Amount")

	q, err := s.queryBuilder(olapReportUrl, params)
	if err != nil {
		return nil, err
	}
	d, err := get(q)
	if err != nil {
		return nil, err
	}

	result := new(ResponseIikoAmountReportData)

	if err := xml.Unmarshal(d, result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s Service) queryBuilder(link string, params interface{}) (string, error) {
	host := fmt.Sprintf("%s%s", s.config.API, link)
	u, err := url.Parse(host)
	if err != nil {
		return "", err
	}
	u.Scheme = "http"
	q := u.Query()
	if pp, ok := params.([]string); ok {
		for _, s := range pp {
			v := strings.Split(s, "=")
			q.Add(v[0], v[1])
		}
	}
	if pp, ok := params.(map[string]string); ok {
		for k, v := range pp {
			q.Set(k, v)
		}
	}

	u.RawQuery = q.Encode()
	res := u.String()
	s.logger.Debug(res)

	return res, nil
}

func get(q string) ([]byte, error) {
	resp, err := http.Get(q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
