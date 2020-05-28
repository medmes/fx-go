package main

import (
	"github.com/medmes/go-tdd/fx/httphandler"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
)

// ApplicationConfig load, as a struct...
type ApplicationConfig struct {
	Address string `yaml:"address"`
}

// Config ...
type Config struct {
	ApplicationConfig `yaml:"application"`
}

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	slogger := logger.Sugar()

	conf := &Config{}
	data, err := ioutil.ReadFile("fx/config/base.yaml")
	if err != nil {
		slogger.Errorf("Could not find or read from config/base.yaml file")
	}

	err = yaml.Unmarshal([]byte(data), &conf)
	if err != nil {
		slogger.Errorf("Could Unmarshal config/base.yaml file, please make sure if it's a valid Yaml file")
	}

	mux := http.NewServeMux()
	httphandler.New(mux)

	//
	http.ListenAndServe(conf.Address, mux)
}
