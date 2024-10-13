package router

import (
	"github.com/04Akaps/gateway_module/config"
	"github.com/04Akaps/gateway_module/log"
	"github.com/04Akaps/gateway_module/types/http"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type get struct {
	engine *fiber.App
	cfg    config.HttpCfg

	handler func(c *fiber.Ctx) error
}

func NewGet(cfg config.HttpCfg, engine *fiber.App) get {
	if err := cfg.Router.GetType.CheckType(); err != nil {
		log.Log.Panic(err.Error(), zap.String("type", cfg.Router.GetType.ToString()))
	}

	g := get{
		cfg:    cfg,
		engine: engine,
	}

	g.handler = func(c *fiber.Ctx) error {
		request := make(map[string]interface{}, 10)

		switch g.cfg.Router.GetType {
		case http.QUERY:

			//utils.co
			//c.Params()
		case http.URL:

		}
		// 1. url path
		// 2. query path

		// TODO Handler

		return nil
	}

	engine.Get(g.cfg.Path, g.handler)

	return g
}
