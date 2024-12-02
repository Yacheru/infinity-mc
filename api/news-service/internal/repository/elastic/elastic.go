package elastic

import (
	"bytes"
	"io"
	"net/http"
	"news-service/init/config"
	"news-service/init/logger"
	"news-service/pkg/constants"
)

type Client struct {
	http   *http.Client
	index  string
	client string
}

func NewElasticClient(cfg *config.Config) (*Client, error) {
	client := &Client{
		http:   new(http.Client),
		index:  cfg.ElasticIndex,
		client: cfg.ElasticClient,
	}

	logger.Debug("pinging elasticsearch...", constants.LoggerElasticsearch)

	if err := client.Ping(); err != nil {
		logger.Error(err.Error(), constants.LoggerElasticsearch)

		return nil, err
	}

	logger.DebugF("creating elastic index (%s)...", constants.LoggerElasticsearch, cfg.ElasticIndex)

	if err := client.CreateIndex(); err != nil {
		logger.Error(err.Error(), constants.LoggerElasticsearch)

		return nil, err
	}

	logger.Info("elastic client is working", constants.LoggerElasticsearch)

	return client, nil
}

func (c *Client) Ping() error {
	resp, err := c.http.Get(c.client)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return constants.ErrorPingElastic
	}

	return nil
}

func (c *Client) CreateIndex() error {
	body := []byte(`{
		"settings": {
        	"number_of_shards" : 1
    	},
		"mappings": {
			"properties": {
				"news": {
					"title": { "type": "string" }
					"description": { "type": "string" }
				}
			}
		}
	}`)

	req, err := http.NewRequest(http.MethodPut, c.client+constants.CreateIndex+c.index, io.NopCloser(bytes.NewBuffer(body)))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}
