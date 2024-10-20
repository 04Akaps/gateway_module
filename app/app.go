package app

import (
	"context"
	"github.com/04Akaps/gateway_module/app/router"
	"github.com/04Akaps/gateway_module/log"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type App struct {
	router map[string]router.Router
}

func NewApp(
	lc fx.Lifecycle,
	router map[string]router.Router,
) App {
	a := App{router: router}

	lc.Append(fx.Hook{
		OnStart: a.onStart,
		OnStop:  a.stop,
	})

	return a
}

func (a App) onStart(c context.Context) error {

	for key, r := range a.router {
		if err := r.Run(); err != nil {
			log.Log.Panic("Failed start server", zap.String("key", key))
		}
	}

	return nil
}

func (a App) stop(c context.Context) error {
	log.Log.Info("lifeCycle ended", zap.Error(c.Err()))
	return nil
}
