package router

import (
	"github.com/04Akaps/gateway_module/app/client"
	"github.com/04Akaps/gateway_module/config"
	"github.com/04Akaps/gateway_module/log"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type delete struct {
	cfg    config.Router
	client client.HttpClient
}

func NewDelete(
	cfg config.Router,
	client client.HttpClient,
) func(c *fiber.Ctx) error {
	if len(cfg.Variable) != 0 {
		log.Log.Panic("we don't support variable in delete request", zap.String("path", cfg.Path))
	}
	d := delete{cfg: cfg, client: client}

	handler := func(c *fiber.Ctx) error {
		return d.handleRequest(c)
	}

	return handler
}

func (d delete) handleRequest(c *fiber.Ctx) error {
	path := d.cfg.Path
	apiResult := d.client.DELETE(c, path, c.Request().Body(), d.cfg)
	return apiResult
}
