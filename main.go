package main

import (
	"github.com/04Akaps/gateway_module/app"
	"github.com/04Akaps/gateway_module/app/dependency"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		dependency.Cfg,
		dependency.HttpClient,
		dependency.Producer,
		dependency.Router,
		fx.Provide(app.NewApp),
		fx.Invoke(func(app.App) {}),
	).Run()
}
