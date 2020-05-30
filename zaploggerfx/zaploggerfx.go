package zaploggerfx

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ProvideZapSugaredLogger
func ProvideZapSugaredLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	return logger.Sugar()
}

var Module = fx.Provide(
	ProvideZapSugaredLogger,
)
