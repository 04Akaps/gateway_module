package app

import (
	"context"
	"github.com/04Akaps/gateway_module/app/router"
	"github.com/04Akaps/gateway_module/log"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type App struct {
	router router.Router
}

func NewApp(
	lc fx.Lifecycle,
	router router.Router,
) App {
	a := App{router: router}

	lc.Append(fx.Hook{
		OnStart: a.onStart,
		OnStop:  a.stop,
	})

	return a
}

func (a App) onStart(c context.Context) error {

	go func() {
		if err := a.router.Run(); err != nil {
			log.Log.Panic("Failed start server")
		}
	}()

	return nil
}

func (a App) stop(c context.Context) error {
	log.Log.Info("lifeCycle ended", zap.Error(c.Err()))
	return nil
}
