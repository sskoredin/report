package ftp_client

import (
	"bytes"
	"fmt"
	"github.com/jlaffaye/ftp"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func (s Service) SendFile(filename string) error {
	s.logger.Infof("Sending file to ftp path %s", s.config.Path)

	c, err := s.dial()
	if err != nil {
		return err
	}

	defer c.Quit()

	err = c.Login(s.config.User, s.config.Password)
	if err != nil {
		return err
	}

	if err := c.ChangeDir(s.config.Path); err != nil {
		return err
	}

	d, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	defer os.Remove(filename)

	data := bytes.NewBuffer(d)
	return c.Stor(s.config.File, data)
}

func (s Service) GetFile(filename string) error {
	s.logger.Infof("Getting file from ftp path %s", s.config.OrdersPath)

	c, err := s.dial()
	if err != nil {
		return err
	}

	defer c.Quit()

	err = c.Login(s.config.User, s.config.Password)
	if err != nil {
		return err
	}

	if err := c.ChangeDir(s.config.OrdersPath); err != nil {
		return err
	}

	resp, err := c.Retr(filename)
	if err != nil {
		return err
	}
	defer resp.Close()

	outFile, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer outFile.Close()

	_, err = io.Copy(outFile, resp)
	return err
}

func (s Service) Walk() ([]*ftp.Entry, error) {
	s.logger.Infof("Walk in ftp path %s", s.config.OrdersPath)

	c, err := s.dial()
	if err != nil {
		return nil, err
	}

	defer c.Quit()

	err = c.Login(s.config.User, s.config.Password)
	if err != nil {
		return nil, err
	}

	return c.List(s.config.OrdersPath)
}

func (s Service) dial() (*ftp.ServerConn, error) {
	c, err := ftp.Dial(fmt.Sprintf("%s:%d", s.config.Host, s.config.Port), ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		return nil, err
	}

	return c, nil
}
