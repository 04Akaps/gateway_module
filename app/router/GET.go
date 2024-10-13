package router

import (
	"github.com/04Akaps/gateway_module/app/client"
	"github.com/04Akaps/gateway_module/config"
	"github.com/04Akaps/gateway_module/log"
	"github.com/04Akaps/gateway_module/types/http"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"go.uber.org/zap"
	"strings"
)

type get struct {
	engine *fiber.App
	cfg    config.HttpCfg

	handler func(c *fiber.Ctx) error
	client  client.HttpClientImpl
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
		var key []string
		var value []string

		switch g.cfg.Router.GetType {
		case http.QUERY:
			// test?key_one=1&key_two=23
			for _, v := range cfg.Router.Variable {
				key = append(key, v)
				value = append(value, utils.CopyString(c.Params(v)))
			}
		case http.URL:
			// path에서 ?이 붙인 부분을 해당 key로 대체하여 전송

			parts := strings.Split(g.cfg.Path, "?")

			var result strings.Builder
			for i, part := range parts {
				result.WriteString(part)
				if i < len(cfg.Router.Variable) {
					result.WriteString(cfg.Router.Variable[i])
				}
			}

			fullUrl := result.String()

			// test/?/next_/?
			//c.
		}
		// 1. url path
		// 2. query path

		// TODO Handler

		return nil
	}

	engine.Get(g.cfg.Path, g.handler)

	return g
}
