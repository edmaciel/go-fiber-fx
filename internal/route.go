package internal

import (
	"net/http"

	"github.com/edmaciel/go-fiber-fx/internal/controller"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

// var Routes = fx.Provide(provideRoutes)

type Route struct {
	fx.Out
	Path    string
	Method  string
	Handler fiber.Handler
}

func ProvideRoutes() []Route {
	return []Route{
		{
			Path:    "/users",
			Method:  http.MethodGet,
			Handler: controller.FindUser,
		},
	}
}
