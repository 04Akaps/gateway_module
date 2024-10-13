package router

import (
	"github.com/04Akaps/gateway_module/config"
	"github.com/gofiber/fiber/v2"
)

type put struct {
	engine *fiber.App
	cfg    config.HttpCfg

	handler func(c *fiber.Ctx) error
}

func NewPut(cfg config.HttpCfg, engine *fiber.App) put {
	h := put{
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
