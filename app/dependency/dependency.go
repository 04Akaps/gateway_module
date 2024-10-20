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
	fx.Provide(func(cfg config.Config) map[string]kafka.Producer {
		clients := make(map[string]kafka.Producer, len(cfg.App))

		for _, app := range cfg.App {
			clients[app.App.Name] = kafka.NewProducer(*app.Producer)
		}

		return clients
	}),
)

var HttpClient = fx.Module(
	"http_client",
	fx.Provide(func(cfg config.Config, producer map[string]kafka.Producer) map[string]client.HttpClient {
		clients := make(map[string]client.HttpClient, len(cfg.App))

		for _, app := range cfg.App {
			clients[app.App.Name] = client.NewHttpClient(app, producer)
		}

		return clients
	}),
)

var Router = fx.Module(
	"router",
	fx.Provide(func(cfg config.Config, client map[string]client.HttpClient) map[string]router.Router {
		clients := make(map[string]router.Router, len(cfg.App))

		for _, app := range cfg.App {
			clients[app.App.Name] = router.NewRouter(app, client)
		}

		return clients
	}),
)
