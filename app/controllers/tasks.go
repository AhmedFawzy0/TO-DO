package controllers

import (
	"errors"

	"github.com/AhmedFawzy0/TO-DO/app/middleware"
	"github.com/AhmedFawzy0/TO-DO/app/models"
	"github.com/AhmedFawzy0/TO-DO/app/repos"
	"github.com/AhmedFawzy0/TO-DO/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func HandleTaskPage(c *fiber.Ctx) error {

	sess, err := middleware.Store.Get(c)
	if err != nil {
		panic(err)
	}
	name := sess.Get("username")
	if name == nil {
		return c.SendString("Unauthenticated, please Sign In!")
	}

	userLoggedIn := new(models.User)
	repos.FindUser(userLoggedIn, name.(string), database.DB.Db)

	var userResponse models.UserResponse
	repos.FindUserTasks(userLoggedIn, &userResponse.Tasks, database.DB.Db)
	userResponse.Tasks = repos.SortById(userResponse.Tasks)
	return c.Render("taskPage", userResponse.Tasks)
}

func SignOut(c *fiber.Ctx) error {

	sess, err := middleware.Store.Get(c)
	if err != nil {
		panic(err)
	}

	if err := sess.Destroy(); err != nil {
		panic(err)
	}

	return c.Redirect("/")
}

func AddTask(c *fiber.Ctx) error {

	task_new := new(models.TaskDTO)
	if err := c.BodyParser(task_new); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	sess, err := middleware.Store.Get(c)
	if err != nil {
		panic(err)
	}
	name := sess.Get("username")
	if name == nil {
		return c.SendString("Unauthenticated, please Sign In!")
	}

	userLoggedIn := new(models.User)

	repos.FindUser(userLoggedIn, name.(string), database.DB.Db)

	newtask := &models.Task{Finished: task_new.Finished, Detail: task_new.Detail}
	taskrespond := &models.TaskResponse{Finished: task_new.Finished, Detail: task_new.Detail}
	repos.AddTask(userLoggedIn, newtask, database.DB.Db)
	taskrespond.ID = newtask.ID

	return c.JSON(fiber.Map{
		"success": true,
		"Task":    taskrespond,
	})
}

func DeleteTask(c *fiber.Ctx) error {

	task_get := new(models.TaskDTO)
	if err := c.BodyParser(task_get); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	task_del := &models.Task{Finished: task_get.Finished, Detail: task_get.Detail}
	task_del.ID = task_get.ID

	sess, err := middleware.Store.Get(c)
	if err != nil {
		panic(err)
	}
	name := sess.Get("username")
	if name == nil {
		return c.SendString("Unauthenticated, please Sign In!")
	}

	error := repos.TaskDelete(task_del.ID, database.DB.Db)

	if errors.Is(error, gorm.ErrRecordNotFound) {
		return errors.New("task not found")
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"deleted": true,
	})

}

func UpdateTask(c *fiber.Ctx) error {
	task_get := new(models.TaskDTO)
	if err := c.BodyParser(task_get); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	task_up := &models.Task{Finished: task_get.Finished, Detail: task_get.Detail}
	task_up.ID = task_get.ID

	sess, err := middleware.Store.Get(c)
	if err != nil {
		panic(err)
	}
	name := sess.Get("username")
	if name == nil {
		return c.SendString("Unauthenticated, please Sign In!")
	}
	update := true

	if repos.UpdateTask(task_up, database.DB.Db) != nil {
		update = false
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"updated": update,
	})

}
