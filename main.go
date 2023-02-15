package main

import (
	"github.com/AhmedFawzy0/TO-DO/app/middleware"
	"github.com/AhmedFawzy0/TO-DO/database"
	"github.com/AhmedFawzy0/TO-DO/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
)

func main() {
	database.ConnectDb()
	middleware.ConnectRedis()

	engine := html.New("./resources/views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.SetupUserRoutes(app)
	routes.SetupTaskRoutes(app)

	app.Static("/", ".resources/assests")

	app.Listen(":3000")

}
