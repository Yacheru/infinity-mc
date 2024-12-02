package elastic

import (
	"bytes"
	"io"
	"net/http"
	"news-service/init/logger"
	"news-service/pkg/constants"
)

type NewsElastic struct {
	es *Client
}

func NewNewsElastic(es *Client) *NewsElastic {
	return &NewsElastic{
		es: es,
	}
}

func (e *NewsElastic) GetNews() error {
	body := []byte(`
		{
			
		}
	`)

	req, err := http.NewRequest(http.MethodGet, e.es.client+e.es.index+constants.Search, io.NopCloser(bytes.NewReader(body)))
	if err != nil {
		logger.Error(err.Error(), constants.LoggerElasticsearch)
	}

	resp, err := e.es.http.Do(req)
	if err != nil {
		logger.Error(err.Error(), constants.LoggerElasticsearch)

		return err
	}
	defer resp.Body.Close()

	return nil
}
