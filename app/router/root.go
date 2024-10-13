package router

import (
	"fmt"
	"github.com/04Akaps/gateway_module/config"
	"github.com/04Akaps/gateway_module/log"
	"github.com/04Akaps/gateway_module/types/http"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"
	"strings"
	"time"
)

type Router struct {
	port   string
	cfg    config.Config
	engine *fiber.App

	getMapper    []get
	postMapper   []post
	deleteMapper []delete
	putMapper    []put
}

func NewRouter(cfg config.Config) Router {
	r := Router{
		cfg:  cfg,
		port: fmt.Sprintf(":%s", cfg.App.Port),
	}

	r.engine = fiber.New()
	r.engine.Use(recover.New())
	r.engine.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: strings.Join([]string{"GET", "POST", "PUT", "DELETE"}, ", "),
		//AllowHeaders:     strings.Join(AllowHeaders, ", "),
		//ExposeHeaders:    strings.Join(AllowHeaders, ", "),
		AllowCredentials: false,
		AllowOriginsFunc: func(origin string) bool { return true },
		MaxAge:           12 * int(time.Hour.Seconds()),
	}))

	r.engine.Get("/healthCheck", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "ok"})
	})

	// TODO HttpCfg의 설정에 따라, HttpClient를 하나씩 배정받아서, Router를 등록해 줘야 한다.

	for _, v := range cfg.Http {
		r.registerRouter(v)
	}

	return r
}

func (r Router) registerRouter(v config.HttpCfg) {
	method := v.Router.Method

	switch method {
	case http.GET:
		r.getMapper = append(r.getMapper, NewGet(v, r.engine))
	case http.POST:
		r.postMapper = append(r.postMapper, NewPost(v, r.engine))
	case http.DELETE:
		r.deleteMapper = append(r.deleteMapper, NewDelete(v, r.engine))
	case http.PUT:
		r.putMapper = append(r.putMapper, NewPut(v, r.engine))
	default:
		log.Log.Panic("Failed to find method", zap.String("method", method.ToString()))
	}

}

// Run
