// Bundle all infra code in this package in order to move out Business/logic code
package infrafx

import (
	"context"
	"github.com/medmes/fx-go/configfx"
	"github.com/medmes/fx-go/httpfx"
	"github.com/medmes/fx-go/zaploggerfx"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
)

// Register Hooks (Managing fx Life Cycle:
func RegisterHooks( // function Parameters:
	lifecycle fx.Lifecycle,
	logger *zap.SugaredLogger,
	conf *configfx.Config,
	mux *http.ServeMux,
) { // function Body:
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go http.ListenAndServe(conf.ApplicationConfig.Address, mux)
				logger.Info("HTTP Listening on port: ", conf.ApplicationConfig.Address)
				return nil
			},
			OnStop: func(ctx context.Context) error {
				logger.Info("Tearing Down!")
				return logger.Sync()
			},
		},
	)
}

var Module = fx.Options(
	configfx.Module,
	zaploggerfx.Module,
	httpfx.Module,
	fx.Invoke(RegisterHooks),
)
