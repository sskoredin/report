package iiko_client

import (
	"crypto/sha1"
	_ "encoding/xml"
	"errors"
	"fmt"
)

func (s Service) auth() error {
	params := make(map[string]string, 2)
	params["login"] = s.config.User
	params["pass"] = passHash(s.config.Password)
	query, err := s.queryBuilder(authUrl, params)
	if err != nil {
		return err
	}

	resp, err := get(query)
	if err != nil {
		return err
	}

	result := fmt.Sprintf(string(resp))
	if len(result) > 45 {
		return errors.New(errNotAuth)
	}

	s.token = fmt.Sprintf(string(resp))
	return nil
}

func passHash(pass string) string {
	h := sha1.New()
	h.Write([]byte(pass))
	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs)
}

func (s Service) Token() string {
	return s.token
}

func (s Service) isAuthorized() error {
	if len(s.token) == 0 {
		return errors.New(errNotAuth)
	}
	return nil
}
