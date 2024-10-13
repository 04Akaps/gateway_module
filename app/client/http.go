package client

import (
	"github.com/04Akaps/gateway_module/config"
	"github.com/bytedance/sonic"
	"github.com/go-resty/resty/v2"
)

type httpClient struct {
	*resty.Client
}

type HttpClientImpl interface {
	QUERY_GET(url string, buffer interface{}, key, value []string) error
	URL_GET(url string, buffer interface{}) error

	POST(url string, req, buffer interface{}) error
	PUT(url string, req, buffer interface{}) error
	DELETE(url string, req, buffer interface{}) error
}

/*
Simple Http Client
Using Router For 1:1 Matching
*/
func HttpClient(cfg config.HttpCfg) HttpClientImpl {
	client := resty.New().
		SetJSONMarshaler(sonic.Marshal).
		SetJSONUnmarshaler(sonic.Unmarshal)

	if cfg.Auth != nil {
		if cfg.Auth.Key != "" {
			client.SetAuthScheme(cfg.Auth.Key)
		}
		client.SetAuthToken(cfg.Auth.Value)
	}

	if cfg.Header != nil {
		client.SetHeaders(cfg.Header)
	}

	// TODO Apply Kafka For Api Trace

	return &httpClient{client}
}

func (h httpClient) QUERY_GET(url string, buffer interface{}, key, value []string) error {
	return nil
}

func (h httpClient) URL_GET(url string, buffer interface{}) error {
	return nil
}

func (h httpClient) POST(url string, req, buffer interface{}) error {
	return nil
}

func (h httpClient) PUT(url string, req, buffer interface{}) error {
	return nil
}

func (h httpClient) DELETE(url string, req, buffer interface{}) error {
	return nil
}
