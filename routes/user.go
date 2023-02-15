package routes

import (
	"github.com/AhmedFawzy0/TO-DO/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	app.Get("/", controllers.MainPage)
	app.Post("/logIn", controllers.LogInLogic)
	app.Post("/user", controllers.CreateUser)
}
