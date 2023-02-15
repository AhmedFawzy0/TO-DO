package routes

import (
	"github.com/AhmedFawzy0/TO-DO/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupTaskRoutes(app *fiber.App) {

	app.Get("/taskPage", controllers.HandleTaskPage)
	app.Post("/addTask", controllers.AddTask)
	app.Get("/signOut", controllers.SignOut)
	app.Delete("/deleteTask", controllers.DeleteTask)
	app.Put("/updateTask", controllers.UpdateTask)

}
