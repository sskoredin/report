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

var params = map[string]string{}

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
