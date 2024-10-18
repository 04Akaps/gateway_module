package router

import (
	"github.com/04Akaps/gateway_module/app/client"
	"github.com/04Akaps/gateway_module/config"
	"github.com/04Akaps/gateway_module/log"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type put struct {
	cfg    config.Router
	client client.HttpClient
}

func NewPut(
	cfg config.Router,
	client client.HttpClient,
) func(c *fiber.Ctx) error {
	if len(cfg.Variable) != 0 {
		log.Log.Panic("we don't support variable in put request", zap.String("path", cfg.Path))
	}

	p := put{cfg: cfg, client: client}

	handler := func(c *fiber.Ctx) error {
		return p.handleRequest(c)
	}

	return handler
}

func (p put) handleRequest(c *fiber.Ctx) error {
	path := p.cfg.Path
	apiResult := p.client.PUT(c, path, c.Request().Body(), p.cfg)
	return apiResult
}
