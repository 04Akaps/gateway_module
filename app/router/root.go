package router

import (
	"fmt"
	"github.com/04Akaps/gateway_module/app/client"
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

	client client.HttpClient
}

func NewRouter(cfg config.Config, client client.HttpClient) Router {
	r := Router{
		cfg:    cfg,
		port:   fmt.Sprintf(":%s", cfg.App.Port),
		client: client,
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

	for _, v := range cfg.Http.Router {
		r.registerRouter(v)
	}

	return r
}

func (r Router) registerRouter(v config.Router) {
	method := v.Method

	switch method {
	case http.GET:
		handler := NewGet(v, r.client)
		r.engine.Get(v.Path, handler)
	case http.POST:
		handler := NewPost(v, r.client)
		r.engine.Post(v.Path, handler)
	case http.DELETE:
		handler := NewDelete(v, r.client)
		r.engine.Delete(v.Path, handler)
	case http.PUT:
		handler := NewPut(v, r.client)
		r.engine.Put(v.Path, handler)
	default:
		log.Log.Panic("Failed to find method", zap.String("method", method.ToString()))
	}

}

func (r Router) Run() error {
	return r.engine.Listen(r.port)
}
