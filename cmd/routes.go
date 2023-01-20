package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/AhmedFawzy0/TO-DO/handlers"

)

func setupRoutes(app *fiber.App) {

	app.Get("/", handlers.ListUsers)
	app.Get("/logIn", handlers.LogInLogic)
	app.Post("/user", handlers.CreateUser)
	app.Get("/taskPage", handlers.HandleTaskPage)
	app.Post("/addTask", handlers.AddTask)
	app.Get("/signOut",handlers.SignOut)
	app.Delete("/deleteTask",handlers.DeleteTask)
	app.Put("/updateTask",handlers.UpdateTask)

}
