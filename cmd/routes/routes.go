package routes

import (
	"context"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go_geofetch/cmd/controllers"
	"github.com/go_geofetch/cmd/models"
	"github.com/go_geofetch/generated"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func Routes(app *fiber.App, ctx context.Context, queries *generated.Queries, env *models.EnvModel, client mqtt.Client) {
	app.Post("/auth/login", controllers.LoginController(ctx, queries, env))
	app.Post("/auth/register", controllers.RegisterController(ctx, queries, env))
	app.Post("/auth/refresh_token", controllers.RefreshToken(ctx, queries, env))
	app.Post("/vehicles/dummy", controllers.DummyVehicleController(ctx, client, env))
	app.Get("/vehicles/:vehicle_id/location", controllers.LastLocationController(ctx, queries))
	app.Get("/vehicles/:vehicle_id/history", controllers.HistoryLocationController(ctx, queries))
	app.Get("/swagger/*", swagger.HandlerDefault)
	// protected := app.Group("/api", middlewares.MiddleWare(env))

}
