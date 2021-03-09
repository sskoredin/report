package telegram_client

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sskoredin/config"
	"log"
	"net/http"
	"strings"
)

type Client struct {
	token string
	url   string
}

func NewLogger() (*Client, error) {
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

type Logger struct {
	*Client
	name string
}

func New(name string) *Logger {
	c, err := NewLogger()
	if err != nil {
		log.Println(err)
	}

	return &Logger{
		Client: c,
		name:   name,
	}
}
func (l Logger) Info(msg ...interface{}) error {
	log.Print()
	return l.send(l.convert("Info", msg))
}

func (l Logger) Infof(format string, v ...interface{}) error {
	return l.send(l.convert("Info", fmt.Sprintf(format, v...)))
}

func (l Logger) Warn(msg ...interface{}) error {
	return l.send(l.convert("Warn", msg))
}

func (l Logger) Warnf(format string, v ...interface{}) error {
	return l.send(l.convert("Debug", fmt.Sprintf(format, v...)))
}

func (l Logger) Debug(msg ...interface{}) error {
	return l.send(l.convert("Debug", msg))
}

func (l Logger) Debugf(format string, v ...interface{}) error {
	return l.send(l.convert("Debug", fmt.Sprintf(format, v...)))
}

func (l Logger) Fatal(msg ...interface{}) error {
	return l.send(l.convert("Fatal", msg))
}

func (l Logger) Fatalf(format string, v ...interface{}) error {
	return l.send(l.convert("Fatal", fmt.Sprintf(format, v...)))
}

func (l Logger) Error(msg ...interface{}) error {
	return l.send(l.convert("Error", msg))
}

func (l Logger) convert(level string, msg ...interface{}) string {
	return fmt.Sprintf("iiko_olap_report\n[%s][%s]\n%v", l.name, level, msg)
}

func (c Client) send(msg string) error {
	payload := strings.NewReader(msg)
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, c.url, payload)
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
