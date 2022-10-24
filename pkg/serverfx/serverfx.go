package serverfx

import (
	"context"

	"github.com/edmaciel/go-fiber-fx/internal"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"fiber",
	fx.Provide(provideApp),
	fx.Invoke(startServer),
)

func provideApp() *fiber.App {
	return fiber.New()
}

type Params struct {
	fx.In
	Routes []internal.Route
}

func startServer(lx fx.Lifecycle, app *fiber.App, p Params) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Server on!!!")
	})

	api := app.Group("/api")
	v1 := api.Group("v1")

	for _, route := range p.Routes {
		v1.Add(route.Method, route.Path, route.Handler)
	}

	app.Listen(":3000")
	lx.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// app.Listen(":3000")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			app.Shutdown()
			return nil
		},
	})
}
