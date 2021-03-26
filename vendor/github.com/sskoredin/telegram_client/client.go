package telegram_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/sskoredin/config"
	"log"
	"net/http"
	"os"
	"path"
)

type Logger interface {
	Info(msg ...interface{}) error
	Infof(format string, v ...interface{}) error
	Debug(msg ...interface{}) error
	Debugf(format string, v ...interface{}) error
	Warn(msg ...interface{}) error
	Warnf(format string, v ...interface{}) error
	Error(msg ...interface{}) error
	Errorf(format string, v ...interface{}) error
	Fatal(msg ...interface{}) error
	Fatalf(format string, v ...interface{}) error
}

type Message struct {
	Msg string `json:"msg"`
}

func NewLogger() (*Client, error) {
	if !isProd() {
		return &Client{}, nil
	}

	client, err := config.NewClient()
	if err != nil {
		return nil, err
	}
	token, err := client.Get(config.TelegramXAuthToken)
	if err != nil {
		return nil, err
	}
	url, err := client.Get(config.TelegramURL)
	if err != nil {
		return nil, err
	}

	return &Client{
		token: token,
		url:   url,
	}, nil
}

type logger struct {
	*Client
	name string
}

func New(name string) Logger {
	c, err := NewLogger()
	if err != nil {
		log.Println(err)
	}

	return &logger{
		Client: c,
		name:   name,
	}
}

func (l logger) Debug(msg ...interface{}) error {
	return l.send(logrus.DebugLevel, l.convert(msg))
}

func (l logger) Debugf(format string, v ...interface{}) error {
	return l.send(logrus.DebugLevel, l.convert(fmt.Sprintf(format, v...)))
}

func (l logger) Info(msg ...interface{}) error {
	return l.send(logrus.InfoLevel, l.convert(msg))
}

func (l logger) Infof(format string, v ...interface{}) error {
	return l.send(logrus.InfoLevel, l.convert(fmt.Sprintf(format, v...)))
}

func (l logger) Warn(msg ...interface{}) error {
	return l.send(logrus.WarnLevel, l.convert(msg))
}

func (l logger) Warnf(format string, v ...interface{}) error {
	return l.send(logrus.WarnLevel, l.convert(fmt.Sprintf(format, v...)))
}
func (l logger) Error(msg ...interface{}) error {
	return l.send(logrus.ErrorLevel, l.convert(msg))
}

func (l logger) Errorf(format string, v ...interface{}) error {
	return l.send(logrus.ErrorLevel, l.convert(fmt.Sprintf(format, v...)))
}

func (l logger) Fatal(msg ...interface{}) error {
	return l.send(logrus.FatalLevel, l.convert(msg))
}

func (l logger) Fatalf(format string, v ...interface{}) error {
	return l.send(logrus.FatalLevel, l.convert(fmt.Sprintf(format, v...)))
}

func (l logger) convert(msg ...interface{}) string {
	return fmt.Sprintf("%s\n[%s]\n%v", programname(), l.name, msg)
}

func programname() string {
	dir, _ := os.Getwd()
	return path.Base(dir)
}

type Client struct {
	token string
	url   string
}

func (c Client) send(level logrus.Level, msg string) error {
	if !isProd() {
		print(level, msg)
		return nil
	}

	client := &http.Client{}
	mes := Message{
		Msg: msg,
	}
	data, err := json.Marshal(mes)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, "http://"+c.url, bytes.NewReader(data))
	if err != nil {
		return err
	}

	req.Header.Add("X-Auth-Token", c.token)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusCreated {
		return errors.New("failed to send message")
	}

	return nil
}

func print(level logrus.Level, msg string) {
	l := logrus.New()
	l.Level = level
	l.Println(msg)
}

func isProd() bool {
	v := os.Getenv("APP_PROD")
	return v != ""
}
