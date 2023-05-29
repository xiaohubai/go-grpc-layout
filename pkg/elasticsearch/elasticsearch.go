package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/xiaohubai/go-grpc-layout/configs/conf"
)

type ElasticSearch struct {
	client *elasticsearch.Client
}

func NewClient(c *conf.Data_ES) (*elasticsearch.Client, error) {
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: c.Address,
		Username:  c.Username,
		Password:  c.Password,
	})
	return client, err
}

func New(client *elasticsearch.Client) *ElasticSearch {
	return &ElasticSearch{client: client}
}

func (es *ElasticSearch) InsertDoc(ctx context.Context, indexName string, b []byte) (err error) {
	var data map[string]interface{}
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(data); err != nil {
		return err
	}
	res, err := es.client.Index(indexName, &buf)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
