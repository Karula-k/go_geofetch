package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go_geofetch/cmd/config"
	"github.com/go_geofetch/cmd/routes"
	"github.com/go_geofetch/db"
	_ "github.com/go_geofetch/docs"
	"github.com/go_geofetch/internal/mqtt"
	"github.com/go_geofetch/internal/mqtt/subscriptions"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func startServer(app *fiber.App, port string) {
	log.Fatal(app.Listen(port))
}

// @title			Order Api
// @version		1.0
// @description	This is an Boilerplate for Backend
// @termsOfService	http://swagger.io/terms/
func main() {
	ctx := context.Background()
	env, err := config.EnvConfig()
	if err != nil {
		log.Fatal(err)
	}
	conn, queries, err := db.InitDB(ctx, &env)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	// Initialize MQTT
	mqttClient, err := mqtt.InitMQTT(&env)
	if err != nil {
		log.Fatal(err)
	}

	subscriptions.MQTTSubscription(ctx, queries, &env, mqttClient)

	fmt.Println(`
	 $$$$$$\   $$$$$$\         $$$$$$\  $$$$$$$\ $$$$$$\ 
	$$  __$$\ $$  __$$\       $$  __$$\ $$  __$$\\_$$  _|
	$$ /  \__|$$ /  $$ |      $$ /  $$ |$$ |  $$ | $$ |  
	$$ |$$$$\ $$ |  $$ |      $$$$$$$$ |$$$$$$$  | $$ |  
	$$ |\_$$ |$$ |  $$ |      $$  __$$ |$$  ____/  $$ |  
	$$ |  $$ |$$ |  $$ |      $$ |  $$ |$$ |       $$ |  
	\$$$$$$  | $$$$$$  |      $$ |  $$ |$$ |     $$$$$$\ 
	\______/  \______/       \__|  \__|\__|     \______|
	`)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin,Content-Type,Accept",
	}))
	routes.Routes(app, ctx, queries, &env, mqttClient)
	startServer(app, ":"+env.Port)
}
