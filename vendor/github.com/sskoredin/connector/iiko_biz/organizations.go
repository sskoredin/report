package iiko_biz_client

import (
	"github.com/sskoredin/connector/domain/entity"
	"github.com/sskoredin/connector/query"
)

func (s *Service) GetOrganizations() ([]entity.Organization, error) {
	q, err := query.New(s.link(organizationsList), s.token).Timeout("00:02:00").
		Build()
	if err != nil {
		return nil, err
	}
	var organizations []entity.Organization
	if err := q.GetWithAttemptsInStruct(&organizations); err != nil {
		return nil, err
	}

	return organizations, nil
}
