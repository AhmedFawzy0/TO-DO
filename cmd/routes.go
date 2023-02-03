package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/AhmedFawzy0/TO-DO/controllers"
	
)

func setupRoutes(app *fiber.App) {

	app.Get("/", controllers.MainPage)
	app.Post("/logIn", controllers.LogInLogic)
	app.Post("/user", controllers.CreateUser)
	app.Get("/taskPage", controllers.HandleTaskPage)
	app.Post("/addTask", controllers.AddTask)
	app.Get("/signOut", controllers.SignOut)
	app.Delete("/deleteTask", controllers.DeleteTask)
	app.Put("/updateTask", controllers.UpdateTask)

}
