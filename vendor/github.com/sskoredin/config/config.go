package config

import (
	"github.com/hashicorp/consul/api"
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
	client *api.Client
}

func NewClient() (*cClient, error) {
	// Get a new client
	config := api.DefaultConfig()
	config.Address = "consul:8600"
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &cClient{
		client,
	}, err
}

func (c cClient) Get(key string) (string, error) {
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
			}
		}
	} else {
		value = string(pair.Value)
	}

	return value, nil
}

func (c cClient) GetInt(key string) (int, error) {
	v, err := c.Get(key)
	if err != nil {
		return 0, err
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

func (c cClient) Set(key string, value string) error {
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
