package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/AhmedFawzy0/TO-DO/handlers"

)

func setupRoutes(app *fiber.App) {

	app.Get("/", handlers.ListUsers)
	app.Post("/logIn", handlers.LogInLogic)
	app.Post("/user", handlers.CreateUser)
	app.Get("/taskPage", handlers.HandleTaskPage)
	app.Post("/addTask", handlers.AddTask)
	app.Get("/signOut",handlers.SignOut)
	app.Post("/addTask",handlers.AddTask)
	app.Post("/deleteTask",handlers.DeleteTask)
	app.Post("/updateTask",handlers.UpdateTask)

}
