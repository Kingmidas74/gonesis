package app

import (
	"context"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/kingmidas74/gonesis-engine/internal/app/container/env"
	"github.com/kingmidas74/gonesis-engine/internal/app/container/handler"
	"github.com/kingmidas74/gonesis-engine/internal/config"
	"github.com/kingmidas74/gonesis-engine/internal/handler/http"
)

type App struct {
	fxOptions  fx.Option
	httpServer *http.Server
	fxApp      *fx.App
}

func New(cfg config.Config, log *zap.Logger) (*App, error) {
	var app = new(App)

	app.FxProvides(
		func() fx.Option {
			return fx.Provide(func() config.Config {
				return cfg
			})
		},

		func() fx.Option {
			return fx.Provide(func() *zap.Logger {
				return log
			})
		},

		handler.New,
		env.New,
	)

	return app, nil
}

func (app *App) FxProvides(ff ...func() fx.Option) {
	options := make([]fx.Option, len(ff))
	for i, f := range ff {
		options[i] = f()
	}
	app.fxOptions = fx.Options(options...)
}

func (app *App) Init() error {
	app.fxOptions = fx.Options(
		// stop timeout: program gives this time to background tasks to complete their work before terminate everything and stop
		fx.StopTimeout(time.Minute*10),
		app.fxOptions,
		fx.NopLogger,

		fx.Populate(&app.httpServer),
		fx.Invoke(func(lc fx.Lifecycle) {
			lc.Append(fx.Hook{
				OnStart: app.onStart,
				OnStop:  app.onStop,
			})
		}),
	)

	app.fxApp = fx.New(app.fxOptions)

	return nil
}

func (app *App) onStart(ctx context.Context) error {
	app.httpServer.Start()
	return nil
}

func (app *App) onStop(ctx context.Context) error {
	app.httpServer.GracefulShutdown()
	return nil
}

func (app *App) Run() error {
	if err := app.fxApp.Err(); err != nil {
		return err
	}

	app.fxApp.Run()
	return nil
}
