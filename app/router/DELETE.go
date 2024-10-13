package router

import (
	"github.com/04Akaps/gateway_module/config"
	"github.com/gofiber/fiber/v2"
)

type delete struct {
	engine *fiber.App
	cfg    config.HttpCfg

	handler func(c *fiber.Ctx) error
}

func NewDelete(cfg config.HttpCfg, engine *fiber.App) delete {
	h := delete{
		cfg:    cfg,
		engine: engine,
	}

	h.handler = func(c *fiber.Ctx) error {
		// TODO Handler

		return nil
	}

	engine.Post(h.cfg.Path, h.handler)

	return h
}
