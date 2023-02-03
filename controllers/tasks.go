package controllers

import (
	"errors"
	//"fmt"
	"sort"
	//"strconv"
	//"time"
	"github.com/AhmedFawzy0/TO-DO/database"
	"github.com/AhmedFawzy0/TO-DO/models"
	"github.com/gofiber/fiber/v2"
	//"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)



func HandleTaskPage(c *fiber.Ctx) error {

	sess, err := database.Store.Get(c) //redis authentication
				if err != nil {
					panic(err)
				}
	name := sess.Get("username")
	if(name==nil){ return c.SendString("Unauthenticated, please Sign In!")}
	
	userLoggedIn:=new(models.User)
	database.DB.Db.Where("Username = ?", name).First(&userLoggedIn)

	var UserTasks []models.Task
	database.DB.Db.Model(&userLoggedIn).Association("Tasks").Find(&UserTasks)
	

	sort.Slice(UserTasks, func(i, j int) bool {
		return UserTasks[i].ID < UserTasks[j].ID
	})

	return c.Render("taskPage", UserTasks)
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

	task_new := new(models.Task)
	if err := c.BodyParser(task_new); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	sess, err := database.Store.Get(c) //redis authentication
				if err != nil {
					panic(err)
				}
	name := sess.Get("username")
	if(name==nil){ return c.SendString("Unauthenticated, please Sign In!")}
	
	userLoggedIn:=new(models.User)
	database.DB.Db.Where("Username = ?", name).First(&userLoggedIn)

	database.DB.Db.Model(&userLoggedIn).Association("Tasks").Append(task_new)

	var task_temp = []models.Task{}
	database.DB.Db.Model(&userLoggedIn).Association("Tasks").Find(&task_temp)

	return c.JSON(fiber.Map{
		"success": true,
		"Task":    task_new,
	})
}

func DeleteTask(c *fiber.Ctx) error {

	task_del := new(models.Task)
	if err := c.BodyParser(task_del); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}


	sess, err := database.Store.Get(c) //redis authentication
				if err != nil {
					panic(err)
				}
	name := sess.Get("username")
	if(name==nil){ return c.SendString("Unauthenticated, please Sign In!")}
	
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
	task_up := new(models.Task)
	if err := c.BodyParser(task_up); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	sess, err := database.Store.Get(c) //redis authentication
	if err != nil {
		panic(err)
	}
name := sess.Get("username")
if(name==nil){ return c.SendString("Unauthenticated, please Sign In!")}

	database.DB.Db.Model(&task_up).Select("Finished", "Detail").Where("id = ?", task_up.ID).Updates(models.Task{Finished: !task_up.Finished, Detail: task_up.Detail})

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"updated": true,
	})

}
