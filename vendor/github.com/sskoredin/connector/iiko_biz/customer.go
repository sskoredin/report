package iiko_biz_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	entity "github.com/sskoredin/connector/domain/entity"
	"github.com/sskoredin/connector/query"
	"time"
)

func (s *Service) GetCustomersByOrganizationAndPeriod(organizationID string, start, end time.Time) ([]entity.Customer, error) {
	period := entity.ReportPeriod{
		Start: start.Format(entity.DateLayout),
		End:   end.Format(entity.DateLayout),
	}
	s.logger.Info(fmt.Sprintf("Collect customers for organization id: %s and period %s - %s", organizationID, period.Start, period.End))

	q, err := query.New(s.link(customersGetCustomersByOrganizationAndPeriod), s.token).
		Organization(organizationID).
		IikoBizPeriod(period).
		Build()
	if err != nil {
		return nil, err
	}

	var customersIiko []entity.CustomerIiko
	if err := q.GetWithAttemptsInStruct(&customersIiko); err != nil {
		return nil, err
	}

	customers := make([]entity.Customer, len(customersIiko))
	for i := range customersIiko {
		customers[i] = customersIiko[i].Convert()
	}

	return customers, nil
}

func (s *Service) GetCustomerDetailedInfo(organizationID, customerID string) (entity.CustomerDetailed, error) {
	q, err := query.New(s.link(customersGetCustomerByID), s.token).
		Organization(organizationID).
		ID(customersGetCustomerByID).
		Build()
	if err != nil {
		return entity.CustomerDetailed{}, err
	}

	var detailed entity.CustomerDetailed
	if err := q.GetWithAttemptsInStruct(&detailed); err != nil {
		return entity.CustomerDetailed{}, err
	}

	return detailed, nil
}

func (s *Service) WriteOff(organizationID, customerID, walletID string, sum float64) error {
	q, err := query.New(s.link(customersWithDrawBalance), s.token).OrganizationID(organizationID).
		CustomerId(customerID).
		WalletId(walletID).
		Sum(fmt.Sprintf("%.2f", sum)).
		Build()
	if err != nil {
		return err
	}
	_, err = q.Post("", nil)
	return err
}

func (s *Service) CreateOrReplaceCustomer(customer entity.Customer) error {
	q, err := query.New(s.link(customerCreateOrReplace), s.token).OrganizationID(customer.Organization).
		Build()
	d, err := json.Marshal(customer)
	if err != nil {
		return err
	}

	_, err = q.Post("", bytes.NewReader(d))
	return err
}
