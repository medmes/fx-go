package main

import (
	"github.com/medmes/fx-go/httphandler"
	"github.com/medmes/fx-go/infrafx"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		infrafx.Module,
		fx.Invoke(httphandler.New),
	).Run()
}
