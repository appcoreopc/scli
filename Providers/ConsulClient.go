package Providers

import (
	"encoding/json"

	"github.com/appcoreopc/scli/Model"
	"github.com/hashicorp/consul/api"
)

type ConsulClient struct {
}

func (c *ConsulClient) WriteKey(key, value string) {

	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	// Get a handle to the KV API
	kv := client.KV()

	// PUT a new KV pair

	p := &api.KVPair{Key: "REDIS_MAXCLIENTS", Value: []byte(`{"page": 1, "fruits": ["apple", "peach"]}`)}
	_, err = kv.Put(p, nil)
	if err != nil {
		panic(err)
	}
}

func (c *ConsulClient) GetKeyJson(key string) interface{} {

	jsonBytes := c.GetKey(key)

	var model = Model.CommandCliModel{}
	json.Unmarshal(jsonBytes, &model)
	return model
}

func (c *ConsulClient) GetKey(key string) []byte {

	client, err := api.NewClient(api.DefaultConfig())

	if err != nil {
		panic(err)
	}

	kv := client.KV()

	pair, _, err := kv.Get(key, nil)
	if err != nil {
		panic(err)
	}

	return pair.Value
}
