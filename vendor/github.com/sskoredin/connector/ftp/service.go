package ftp_client

import (
	"bytes"
	"fmt"
	"github.com/jlaffaye/ftp"
	"io/ioutil"
	"os"
	"time"
)

func (s Service) SendGoods(filename string) error {
	s.logger.Infof("Sending file to ftp path %s", s.config.Path)

	c, err := ftp.Dial(fmt.Sprintf("%s:%d", s.config.Host, s.config.Port), ftp.DialWithTimeout(5*time.Second))
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
