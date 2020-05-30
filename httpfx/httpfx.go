package httpfx

import (
	"go.uber.org/fx"
	"net/http"
)

var Module = fx.Options(
	fx.Provide(http.NewServeMux),
)
