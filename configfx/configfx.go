package configfx

import (
	"go.uber.org/fx"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

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

	data, err := ioutil.ReadFile("config/base.yaml")
	if err != nil {
		panic("Could not find or read from config/base.yaml file") // stopping because we cannot start without proper config
	}

	if err := yaml.Unmarshal([]byte(data), &conf); err != nil {
		panic("Could Unmarshal config/base.yaml file, please make sure if it's a valid Yaml file")
	}

	return &conf
}

var Module = fx.Options(
	fx.Provide(ProvideConfig),
)
