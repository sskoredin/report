package config

import (
	"github.com/hashicorp/consul/api"
	"log"
	"os"
	"strconv"
	"strings"
)

type ConsulClient interface {
	Get(key string) (string, error)
	GetInt(key string) (int, error)
	GetArray(key string) ([]string, error)
	Set(key string, value string) error
}

type cClient struct {
	client    *api.Client
	devParams map[string]string
}

var params = map[string]string{
	"TELEGRAM_CHAT_IDS":    "1539614935_492307185",
	"TELEGRAM_XAUTH-TOKEN": "111111-111",
	"TELEGRAM_URL":         "logger:3250",

	"FTP_HOST":       "ftp.davido01.nichost.ru",
	"FTP_PORT":       "21",
	"FTP_USER":       "davido01_skor",
	"FTP_PASSWORD":   "uLigERaOUpEliSChI1",
	"FTP_PATH":       "davidoff-bronnaya.ru/docs/auto",
	"FTP_ORDER_PATH": "davidoff-bronnaya.ru/docs/auto",
	"FTP_FILE":       "import.csv",
	"REST_HOST":      "0.0.0.0",
	"REST_PORT":      "8230",

	"API_IIKO_URL":  "http://94.127.179.181:9081",
	"IIKO_LOGIN":    "Yakupov",
	"IIKO_PASSWORD": "1Qz2876",

	"API_BIZ_IIKO_URL":             "http://94.127.179.181:9900",
	"API_BIZ_CATEGORY_NAME":        "Уведомлен о сгорании бонусов",
	"IIKO_BIZ_LOGIN":               "sergey@skoredin.pro",
	"IIKO_BIZ_PASSWORD":            "B@rkas1257",
	"MAIL_HOST":                    "mail.nic.ru",
	"MAIL_PORT":                    "587",
	"MAIL_LOGIN":                   "robot@davidoffclub.ru",
	"MAIL_PASSWORD":                "SwITERyCHEwConoT1",
	"MAIL_ADDRESSEE":               "robot",
	"MAIL_OLAP_SUBJECT":            "OLAP report",
	"MAIL_OLAP_RECIPIENTS":         "[sskoredin@gmail.com]",
	"MAIL_AMOUNT_SUBJECT":          "OLAP Ostatki v minuse report iiko",
	"MAIL_AMOUNT_RECIPIENTS":       "[sskoredin@gmail.com]",
	"DAEMON_OLAP_SCHEDULER":        "37 12 * * *",
	"DAEMON_AMOUNT_SCHEDULER":      "37 12 * * *",
	"DAEMON_INTEGRATION_SCHEDULER": "37 12 * * *",
	"POSTGRES_HOST":                "0.0.0.0",
	"POSTGRES_PORT":                "5432",
	"POSTGRES_DB":                  "iiko",
	"POSTGRES_USER":                "support",
	"POSTGRES_PASSWORD":            "qwerty",
}

func NewClient() (*cClient, error) {
	// Get a new client
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return nil, err
	}
	return &cClient{
		client:    client,
		devParams: params,
	}, err
}

func (c cClient) Get(key string) (string, error) {
	if !isProd() {
		v := c.getDev(key)
		return v, nil
	}

	kv := c.client.KV()
	// Lookup the pair
	pair, _, err := kv.Get(key, nil)
	if err != nil {
		return "", err
	}

	var value string
	if pair == nil || len(pair.Value) == 0 {
		if v, ok := defaultValues[key]; ok {
			value = v
			if err := c.Set(key, value); err != nil {
				return "", err
			} else {
				log.Printf("Not found value for  %s", key)
			}
		}
	} else {
		value = string(pair.Value)
	}

	return value, nil
}

func (c cClient) getDev(key string) string {
	return c.devParams[key]
}

func (c cClient) GetInt(key string) (int, error) {
	var v string
	var err error
	if isProd() {
		v, err = c.Get(key)
		if err != nil {
			return 0, err
		}
	} else {
		v = c.getDev(key)
	}

	i, err := strconv.Atoi(v)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func (c cClient) GetArray(key string) ([]string, error) {
	array, err := c.Get(key)
	if err != nil {
		return nil, err
	}
	array = strings.Trim(array, "[")
	array = strings.Trim(array, "]")
	return strings.Split(array, ","), nil

}
func (c cClient) setDev(key string, value string) {
	c.devParams[key] = value
}
func (c cClient) Set(key string, value string) error {
	if !isProd() {
		c.setDev(key, value)
		return nil
	}

	if err := c.updateEnv(key, value); err != nil {
		return err
	}

	// Get a handle to the KV API
	kv := c.client.KV()

	// PUT a new KV pair
	p := &api.KVPair{Key: key, Value: []byte(value)}
	_, err := kv.Put(p, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c cClient) updateEnv(key, value string) error {
	if value == "" {
		return nil
	}
	if !isProd() {
		c.Set(key, value)
		return nil
	}

	if os.Getenv(key) == value {
		return nil
	}

	return os.Setenv(key, value)
}

func isProd() bool {
	v := os.Getenv("APP_PROD")
	return v != ""
}
