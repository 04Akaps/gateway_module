package router

import (
	"github.com/04Akaps/gateway_module/config"
	"github.com/gofiber/fiber/v2"
)

type post struct {
	engine *fiber.App
	cfg    config.HttpCfg

	handler func(c *fiber.Ctx) error
}

func NewPost(cfg config.HttpCfg, engine *fiber.App) post {
	g := post{
		cfg:    cfg,
		engine: engine,
	}

	g.handler = func(c *fiber.Ctx) error {
		// TODO Handler

		return nil
	}

	engine.Post(g.cfg.Path, g.handler)

	return g
}
