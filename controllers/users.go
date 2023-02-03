package controllers

import (
	"github.com/AhmedFawzy0/TO-DO/database"
	"github.com/AhmedFawzy0/TO-DO/models"
	"github.com/gofiber/fiber/v2"
	//"github.com/golang-jwt/jwt/v4"
)

const SecretKey = "secret"

func MainPage(c *fiber.Ctx) error {
		return c.Render("mainPage",nil)
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.UserDTO)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	
	userTemp:= new(models.User)
	database.DB.Db.Where("Username = ?", user.Username).First(&userTemp)
	if(userTemp.Username==user.Username){
		return c.JSON(fiber.Map{
			"UserExists": true,
			"UserCreated": false,
		})
	}else{
		userTemp.Username=user.Username
		userTemp.Password=user.Password
		err:=database.DB.Db.Create(&userTemp).Error
		if err!=nil{
			return c.JSON(fiber.Map{
				"UserExists": true,
			    "UserCreated": false,
			})
			}

		return c.JSON(fiber.Map{
			"UserExists": false,
			"UserCreated": true,
		})
	}

		
}

func LogInLogic(c *fiber.Ctx) error {

	user1 := new(models.UserDTO)
	if err := c.BodyParser(user1); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}


	userTemp:= new(models.User)
	database.DB.Db.Where("Username = ?", user1.Username).First(&userTemp)

	if(userTemp.Username==user1.Username && userTemp.Password==user1.Password){
			sess, err := database.Store.Get(c)
			if err != nil {
				panic(err)}
			sess.Set("username", user1.Username)
			if err := sess.Save(); err != nil {
				panic(err)}	

				return c.JSON(fiber.Map{
					"success":true,
					"UserExists":true,})
	}else{
		return c.JSON(fiber.Map{
			"success":false,
			"UserExists":userTemp.Username==user1.Username,})
	}
	}



