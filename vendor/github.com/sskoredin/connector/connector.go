package connector

import (
	"github.com/jlaffaye/ftp"
	"github.com/sskoredin/connector/domain/entity"
	ftpclient "github.com/sskoredin/connector/ftp"
	"github.com/sskoredin/connector/iiko"
	iikobizclient "github.com/sskoredin/connector/iiko_biz"
	"time"
)

type IikoClient interface {
	CollectGoodsForFTP(period entity.ReportPeriod) (*iiko_client.ResponseGoodsToFTPData, error)
	CollectIikoOLAPReport(period entity.ReportPeriod) (*iiko_client.ResponseIikoOlAPReportData, error)
	CollectIikoAmountReport(period entity.ReportPeriod) (*iiko_client.ResponseIikoAmountReportData, error)
}

func NewIikoClient() (IikoClient, error) {
	s, err := iiko_client.New()
	if err != nil {
		return nil, err
	}
	return s, nil
}

type IikoBizClient interface {
	Auth() error
	CheckToken() error

	GetOrganizations() ([]entity.Organization, error)
	GetTransactionsByOrganizationAndPeriod(organizationID, customerID string, start, end time.Time) ([]entity.Transaction, error)
	GetPaymentsTypes(organizationID string) ([]entity.Payment, error)

	GetCustomersByOrganizationAndPeriod(organizationID string, start, end time.Time) ([]entity.Customer, error)
	GetCustomerDetailedInfo(organizationID, customerID string) (entity.CustomerDetailed, error)
	RemoveCategoryFromCustomer(customerID, organizationId, categoryID string) error
	AddCategoryToCustomer(customerID, organizationId, categoryID string) error
	WriteOff(organizationID, customerID, walletID string, sum float64) error
	CreateOrReplaceCustomer(customer entity.Customer) (string, error)

	GetDiscounts(organization string) ([]entity.IikoDiscount, error)
	GetNomenclature(organizationID string) (*entity.Nomenclature, error)
	AddOrder(order entity.OrderQuery) error

	GetMenu(organizationID string) (*entity.Nomenclature, error)
}

func NewIikoBizClient() (IikoBizClient, error) {
	s, err := iikobizclient.New()
	if err != nil {
		return nil, err
	}
	return s, nil
}

type FTPClient interface {
	SendFile(filename string) error
	GetFile(filename string) error
	Walk() ([]*ftp.Entry, error)
	DeleteFiles(files []entity.OrderFile) error
}

func NewFTPClient() (FTPClient, error) {
	s, err := ftpclient.New()
	if err != nil {
		return nil, err
	}
	return s, nil
}
