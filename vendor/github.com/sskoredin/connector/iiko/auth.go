package iiko_client

import (
	_ "encoding/xml"
	"errors"
	"fmt"
	"github.com/sskoredin/connector/domain/entity"
	query "github.com/sskoredin/connector/query"
)

func (s *Service) auth() error {
	q, err := query.New(s.link(authUrl), "").
		Login(s.config.User).
		Pass(s.config.Password).
		Build()
	if err != nil {
		return err
	}
	s.logger.Info(q)
	resp, err := q.Get()
	if err != nil {
		return err
	}

	result := fmt.Sprintf(string(resp))
	if len(result) > 45 {
		return errors.New(entity.ErrNotAuth)
	}

	s.token = fmt.Sprintf(string(resp))
	s.logger.Infof("token %s", s.token)
	return nil
}

func (s Service) Token() string {
	return s.token
}
