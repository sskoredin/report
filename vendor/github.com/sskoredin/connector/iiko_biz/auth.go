package iiko_biz_client

import (
	"github.com/sskoredin/connector/query"
	"strings"
	"time"
)

func (s *Service) Auth() error {
	if err := s.config.Read(); err != nil {
		return err
	}

	now := time.Now()
	if s.tokenCreatedAt != nil && now.Sub(*s.tokenCreatedAt) < 2*time.Minute {
		return nil
	}
	q, err := query.New(s.link(authUrl), "").
		UserID(s.config.User).
		UserSecret(s.config.Password).
		Build()
	if err != nil {
		return err
	}

	d, err := q.Get()
	if err != nil {
		return err
	}

	s.token = strings.ReplaceAll(string(d), "\"", "")
	s.tokenCreatedAt = &now
	return nil
}

func (s *Service) AuthV2() error {
	if err := s.config.Read(); err != nil {
		return err
	}

	now := time.Now()
	if s.tokenCreatedAt != nil && now.Sub(*s.tokenCreatedAt) < 2*time.Minute {
		return nil
	}
	q, err := query.New(s.link(authUrl), "").
		UserID(s.config.User).
		Pass(s.config.Password).
		Build()
	if err != nil {
		return err
	}

	d, err := q.Get()
	if err != nil {
		return err
	}

	s.token = strings.ReplaceAll(string(d), "\"", "")
	s.tokenCreatedAt = &now
	return nil
}

func (s *Service) CheckToken() error {
	if err := s.config.Read(); err != nil {
		return err
	}
	_, err := query.New(s.link(authUrl), "").
		UserID(s.config.User).
		UserSecret(s.config.Password).
		Build()

	return err
}

func (s Service) Token() string {
	return s.token
}
