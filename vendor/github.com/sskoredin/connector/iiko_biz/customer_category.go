package iiko_biz_client

import (
	"fmt"
	"github.com/sskoredin/connector/query"
	"github.com/sskoredin/go_iiko/domain/entity"
)

func (s *Service) GetCategoriesByOrganization(organizationID string) ([]entity.GuestCategory, error) {
	s.logger.Info(fmt.Sprintf("get category_id for %s", s.config.CategoryName))
	q, err := query.New(s.linkWithValue(organizationGuestCategory, organizationID), s.token).
		Build()
	if err != nil {
		return nil, err
	}

	var categories []entity.GuestCategory
	if err := q.GetWithAttemptsInStruct(&categories); err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *Service) AddCategoryToCustomer(customerID, organizationId, categoryID string) error {
	q, err := query.New(s.linkWithValue(customerAddCategory, customerID), s.token).
		Organization(organizationId).CategoryId(categoryID).
		Build()
	if err != nil {
		return err
	}

	_, err = q.Post("", nil)
	return err

}

func (s *Service) RemoveCategoryFromCustomer(customerID, organizationId, categoryID string) error {
	q, err := query.New(s.linkWithValue(customerRemoveCategory, customerID), s.token).
		Organization(organizationId).CategoryId(categoryID).
		Build()
	if err != nil {
		return err
	}

	_, err = q.Post("", nil)
	return err
}
