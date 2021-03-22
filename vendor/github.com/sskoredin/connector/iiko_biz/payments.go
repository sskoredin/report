package iiko_biz_client

import (
	"github.com/sskoredin/connector/domain/entity"
	"github.com/sskoredin/connector/query"
)

func (s *Service) GetPaymentsTypes(organizationID string) ([]entity.Payment, error) {
	q, err := query.New(getPaymentsTypes, s.token).OrganizationID(organizationID).
		Build()
	if err != nil {
		return nil, err
	}
	var payments []entity.Payment
	if err := q.GetWithAttemptsInStruct(payments); err != nil {
		return nil, err
	}

	return payments, nil
}
