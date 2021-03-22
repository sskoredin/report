package iiko_biz_client

import (
	"github.com/sskoredin/connector/domain/entity"
	"github.com/sskoredin/connector/query"
)

func (s *Service) GetMenu(organizationID string) ([]entity.Nomenclature, error) {
	q, err := query.New(s.linkWithValue(nomenclature, organizationID), s.token).
		Build()
	if err != nil {
		return nil, err
	}
	var nomenclature []entity.Nomenclature
	if err := q.GetWithAttemptsInStruct(&nomenclature); err != nil {
		return nil, err
	}

	return nomenclature, nil
}
