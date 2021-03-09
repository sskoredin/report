package config

import (
	"github.com/hashicorp/consul/api"
)

type ConsulClient interface {
	Get(key string) (string, error)
	Set(key string, value string) error
}

type cClient struct {
	client *api.Client
}

func NewClient() (*cClient, error) {
	// Get a new client
	client, err := api.NewClient(api.DefaultConfig())
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
