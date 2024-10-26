package client

import (
	"github.com/04Akaps/gateway_module/common"
	"github.com/04Akaps/gateway_module/config"
	"github.com/04Akaps/gateway_module/kafka"
	"github.com/04Akaps/gateway_module/log"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"sync"
)

const (
	_baseUrlSetError = "Base Url is not existed"

	_defaultBatchSize = 10
	_defaultBatchTime = 2
)

type HttpClient struct {
	client   *resty.Client
	cfg      config.App
	producer kafka.Producer

	batchTime float64

	mapper       []ApiRequestTopic
	fetchLock    sync.Mutex
	fetchChannel chan ApiRequestTopic
}

func NewHttpClient(
	cfg config.App,
	producer map[string]kafka.Producer,
) HttpClient {
	batchTime := cfg.Producer.BatchTime

	if batchTime == 0 {
		batchTime = _defaultBatchTime
	}

	client := resty.New().
		SetJSONMarshaler(common.JsonHandler.Marshal).
		SetJSONUnmarshaler(common.JsonHandler.Unmarshal)

	if cfg.Http.BaseURL == "" {
		log.Log.Panic(_baseUrlSetError, zap.String("baseUrl", cfg.Http.BaseURL))
	}

	client.SetBaseURL(cfg.Http.BaseURL)

	return HttpClient{
		client:       client,
		cfg:          cfg,
		producer:     producer[cfg.App.Name],
		batchTime:    batchTime,
		mapper:       make([]ApiRequestTopic, 0), // 빈 슬라이스로 초기화
		fetchChannel: make(chan ApiRequestTopic),
	}
}

func (h *HttpClient) GET(url string, router config.Router) (*resty.Response, error) {
	var err error
	var req *resty.Request
	var resp *resty.Response

	req = h.getRequest(router)
	resp, err = req.Get(url)

	_, err = common.CB.Execute(func() ([]byte, error) {
		req = h.getRequest(router)
		resp, err = req.Get(url)

		if err != nil {
			return nil, err
		}

		return nil, nil
	})

	defer h.handleRequest(resp, req.Body)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *HttpClient) POST(url string, requestBody interface{}, router config.Router) (*resty.Response, error) {
	var err error
	var req *resty.Request
	var resp *resty.Response

	_, err = common.CB.Execute(func() ([]byte, error) {
		req = h.getRequest(router).SetBody(requestBody)
		resp, err = req.Post(url)

		if err != nil {
			return nil, err
		}

		return nil, nil
	})

	defer h.handleRequest(resp, req.Body)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *HttpClient) PUT(url string, requestBody interface{}, router config.Router) (*resty.Response, error) {
	var err error
	var req *resty.Request
	var resp *resty.Response

	_, err = common.CB.Execute(func() ([]byte, error) {
		req = h.getRequest(router).SetBody(requestBody)
		resp, err = req.Put(url)

		if err != nil {
			return nil, err
		}

		return nil, nil
	})

	defer h.handleRequest(resp, req.Body)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *HttpClient) DELETE(url string, requestBody interface{}, router config.Router) (*resty.Response, error) {
	var err error
	var req *resty.Request
	var resp *resty.Response

	_, err = common.CB.Execute(func() ([]byte, error) {
		req = h.getRequest(router).SetBody(requestBody)
		resp, err = req.Delete(url)

		if err != nil {
			return nil, err
		}

		return nil, nil
	})

	defer h.handleRequest(resp, req.Body)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *HttpClient) getRequest(router config.Router) *resty.Request {
	req := h.client.R().EnableTrace()

	if router.Auth != nil {
		if len(router.Auth.Schema) != 0 {
			req.SetAuthScheme(router.Auth.Schema)
		}

		req.SetAuthToken(router.Auth.Token)
	}

	if router.Header != nil {
		req.SetHeaders(router.Header)
	}

	return req
}
