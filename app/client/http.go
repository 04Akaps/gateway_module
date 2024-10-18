package client

import (
	"github.com/04Akaps/gateway_module/common"
	"github.com/04Akaps/gateway_module/config"
	"github.com/04Akaps/gateway_module/kafka"
	"github.com/04Akaps/gateway_module/log"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
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
	cfg      config.Config
	producer kafka.Producer

	batchTime float64

	mapper       []ApiRequestTopic
	fetchLock    sync.Mutex
	fetchChannel chan ApiRequestTopic
}

func NewHttpClient(
	cfg config.Config,
	producer kafka.Producer,
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
		producer:     producer,
		batchTime:    batchTime,
		mapper:       make([]ApiRequestTopic, 0), // 빈 슬라이스로 초기화
		fetchChannel: make(chan ApiRequestTopic),
	}
}

func (h *HttpClient) GET(c *fiber.Ctx, url string, router config.Router) error {
	var buffer interface{}

	req := h.getRequest(router).SetResult(&buffer)
	resp, err := req.Get(url)
	defer h.handleRequest(resp, req)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(NewCallError(url, err, resp))
	}

	return c.Status(resp.StatusCode()).JSON(buffer)
}

func (h *HttpClient) POST(c *fiber.Ctx, url string, requestBody interface{}, router config.Router) error {
	var buffer interface{}

	req := h.getRequest(router).SetResult(&buffer).SetBody(requestBody)
	resp, err := req.Post(url)
	defer h.handleRequest(resp, req)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(NewCallError(url, err, resp))
	}

	return c.Status(resp.StatusCode()).JSON(buffer)
}

func (h *HttpClient) PUT(c *fiber.Ctx, url string, requestBody interface{}, router config.Router) error {
	var buffer interface{}

	req := h.getRequest(router).SetResult(&buffer).SetBody(requestBody)
	resp, err := req.Put(url)
	defer h.handleRequest(resp, req)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(NewCallError(url, err, resp))
	}

	return c.Status(resp.StatusCode()).JSON(buffer)
}

func (h *HttpClient) DELETE(c *fiber.Ctx, url string, requestBody interface{}, router config.Router) error {
	var buffer interface{}

	req := h.getRequest(router).SetResult(&buffer).SetBody(requestBody)
	resp, err := req.Delete(url)
	defer h.handleRequest(resp, req)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(NewCallError(url, err, resp))
	}

	return c.Status(resp.StatusCode()).JSON(buffer)
}

func (h *HttpClient) getRequest(router config.Router) *resty.Request {
	request := h.client.R().EnableTrace()
	setRequest(request, router)
	return request
}

func setRequest(req *resty.Request, router config.Router) {
	if router.Auth != nil {
		if len(router.Auth.Schema) != 0 {
			req.SetAuthScheme(router.Auth.Schema)
		}

		req.SetAuthToken(router.Auth.Token)
	}

	if router.Header != nil {
		req.SetHeaders(router.Header)
	}
}
