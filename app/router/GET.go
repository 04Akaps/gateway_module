package router

import (
	"fmt"
	"github.com/04Akaps/gateway_module/app/client"
	"github.com/04Akaps/gateway_module/config"
	"github.com/04Akaps/gateway_module/log"
	"github.com/04Akaps/gateway_module/types/http"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"go.uber.org/zap"
	"strings"
)

const (
	_not_supported_type = "Failed to find api type"
)

/*
	전송하고자 하는 url이 다음과 같다고 가정하자.

	1. https://test/v1/user/test/123

	이런 url 형태에서 대해서는 yaml에 다음과 같은 형식으로 받게 규격을 정해놓음

	base_url - https://test
	router.path - /v1/user/:id/:number
	get_type: "url",

	해당 gateway에서는 동일하게 /v1/user/?의 형태로 api 생성
	- localhost/v1/user/:?

	해당 값이 들어오면, 값을 파싱해서, 전송


	2. https://test/v1/auth?name=123123&password=1930

	이런 url 형태에서 대해서는 yaml에 다음과 같은 형식으로 받게 규격을 정해놓음

	base_url - https://test
	router.path - /v1/auth
	varibale - [name, password]
	get_type: "query",
*/

type get struct {
	cfg    config.Router
	client client.HttpClient
}

func NewGet(
	cfg config.Router,
	client client.HttpClient,
) func(c *fiber.Ctx) error {
	if err := cfg.GetType.CheckType(); err != nil {
		log.Log.Panic(err.Error(), zap.String("type", cfg.GetType.ToString()))
	}

	g := get{cfg: cfg, client: client}

	handler := func(c *fiber.Ctx) error {
		switch g.cfg.GetType {
		case http.QUERY:
			return g.queryType(c)
		case http.URL:
			return g.urlType(c)
		default:
			log.Log.Error(
				_not_supported_type,
				zap.String("type", g.cfg.GetType.ToString()),
			)
			return c.Status(fiber.StatusNotFound).JSON(_not_supported_type)
		}
	}

	return handler
}

func (g get) queryType(c *fiber.Ctx) error {

	var urlBuilder strings.Builder

	path := g.cfg.Path

	for i, v := range g.cfg.Variable {
		key := v
		value := utils.CopyString(c.Query(key))

		if i == 0 {
			path += fmt.Sprintf("?%s=%s", key, value)
		} else {
			path += fmt.Sprintf("&%s=%s", key, value)
		}
	}

	urlBuilder.WriteString(path)
	fullUrl := urlBuilder.String()

	apiResult, err := g.client.GET(fullUrl, g.cfg)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(client.NewCallError(fullUrl, err, apiResult))
	}

	return c.Status(apiResult.StatusCode()).JSON(string(apiResult.Body()))
}

func (g get) urlType(c *fiber.Ctx) error {
	var urlBuilder strings.Builder
	urlBuilder.WriteString(string(c.Request().URI().Path()))
	fullUrl := urlBuilder.String()

	apiResult, err := g.client.GET(fullUrl, g.cfg)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(client.NewCallError(fullUrl, err, apiResult))
	}

	return c.Status(apiResult.StatusCode()).JSON(string(apiResult.Body()))
}
