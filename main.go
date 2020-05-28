package main

import (
	"context"
	"fmt"
	"github.com/medmes/go-tdd/fx/httphandler"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
)

func main() {
	fx.New(
		fx.Provide(ProvideConfig),
		fx.Provide(ProvideZapSugaredLogger),
		fx.Provide(http.NewServeMux),
		fx.Invoke(httphandler.New),
		fx.Invoke(RegisterHooks),
	).Run()
}

// Register Hooks (Managing fx Life Cycle:
func RegisterHooks( // function Parameters:
	lifecycle fx.Lifecycle,
	logger *zap.SugaredLogger,
	conf *Config,
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

// ApplicationConfig load, as a struct...
type ApplicationConfig struct {
	Address string `yaml:"address"`
}

// Config ...
type Config struct {
	ApplicationConfig `yaml:"application"`
}

// ProvideConfig
func ProvideConfig() *Config {
	conf := Config{}

	data, err := ioutil.ReadFile("fx/config/base.yaml")
	if err != nil {
		fmt.Println("Could not find or read from config/base.yaml file")
	}

	if err := yaml.Unmarshal([]byte(data), &conf); err != nil {
		fmt.Println("Could Unmarshal config/base.yaml file, please make sure if it's a valid Yaml file")
	}

	return &conf
}

// ProvideZapSugaredLogger
func ProvideZapSugaredLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	return logger.Sugar()
}
