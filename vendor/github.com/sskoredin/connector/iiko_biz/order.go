package iiko_biz_client

import (
	"bytes"
	"encoding/json"
	"github.com/sskoredin/connector/domain/entity"
	"github.com/sskoredin/connector/query"
)

func (s Service) GetDiscounts(organization string) ([]entity.IikoDiscount, error) {
	var err error
	q, err := query.New(s.link(deliveryDiscounts), s.token).
		Organization(organization).
		Build()
	if err != nil {
		return nil, err
	}
	var discounts []entity.IikoDiscount
	if err := q.GetWithAttemptsInStruct(discounts); err != nil {
		return nil, err
	}

	return discounts, nil
}

func (s Service) GetNomenclature(organizationID string) (*entity.Nomenclature, error) {
	q, err := query.New(s.linkWithValue(getNomenclature, organizationID), s.token).
		Build()
	if err != nil {
		return nil, err
	}
	var nomenclature entity.Nomenclature
	if err := q.GetWithAttemptsInStruct(&nomenclature); err != nil {
		return nil, err
	}

	return &nomenclature, nil
}

func (s Service) AddOrder(order entity.OrderQuery) error {
	if err := s.AuthV2(); err != nil {
		return err
	}
	q, err := query.New(s.link1(addOrder), s.token).
		RequestTimeout("10000").
		Build()
	if err != nil {
		return err
	}
	d, err := json.Marshal(order)
	if err != nil {
		return err
	}
	_, err = q.Post("application/json", bytes.NewBuffer(d))
	return err
}
