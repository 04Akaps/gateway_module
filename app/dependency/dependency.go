package dependency

import (
	"flag"
	"github.com/04Akaps/gateway_module/app/client"
	"github.com/04Akaps/gateway_module/app/router"
	"github.com/04Akaps/gateway_module/config"
	"github.com/04Akaps/gateway_module/kafka"
	"go.uber.org/fx"
)

var (
	yamlPath = flag.String("yamlPath", "./deploy_sample.yaml", "config path")
)

func init() {
	flag.Parse()
}

var Cfg = fx.Module(
	"config",
	fx.Provide(func() config.Config {
		return config.NewCfg(*yamlPath)
	}),
)

var Producer = fx.Module(
	"kafka_producer",
	fx.Provide(func(cfg config.Config) kafka.Producer {
		return kafka.NewProducer(cfg.Producer)
	}),
)

var HttpClient = fx.Module(
	"http_client",
	fx.Provide(func(cfg config.Config, producer kafka.Producer) client.HttpClient {
		return client.NewHttpClient(cfg, producer)
	}),
)

var Router = fx.Module(
	"router",
	fx.Provide(func(cfg config.Config, client client.HttpClient) router.Router {
		return router.NewRouter(cfg, client)
	}),
)
