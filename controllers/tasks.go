package controllers

import (
	"errors"

	"github.com/AhmedFawzy0/TO-DO/database"
	"github.com/AhmedFawzy0/TO-DO/models"
	"github.com/AhmedFawzy0/TO-DO/repos"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func HandleTaskPage(c *fiber.Ctx) error {

	sess, err := database.Store.Get(c)
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

	sess, err := database.Store.Get(c)
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
	sess, err := database.Store.Get(c)
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

	sess, err := database.Store.Get(c)
	if err != nil {
		panic(err)
	}
	name := sess.Get("username")
	if name == nil {
		return c.SendString("Unauthenticated, please Sign In!")
	}

	var task_temp models.Task

	error := database.DB.Db.Delete(&task_temp, task_del.ID).Error
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

	sess, err := database.Store.Get(c)
	if err != nil {
		panic(err)
	}
	name := sess.Get("username")
	if name == nil {
		return c.SendString("Unauthenticated, please Sign In!")
	}

	database.DB.Db.Model(&task_up).Select("Finished", "Detail").Where("id = ?", task_up.ID).Updates(models.Task{Finished: !task_up.Finished, Detail: task_up.Detail})

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"updated": true,
	})

}
