package repos

import (
	"github.com/AhmedFawzy0/TO-DO/models"
	"gorm.io/gorm"
)

func FindUserTasks(user *models.User, usertask *[]models.Task, db1 *gorm.DB) (err error) {
	return db1.Model(user).Association("Tasks").Find(usertask)
}

func AddTask(user *models.User, task *models.Task, db1 *gorm.DB) (err error) {
	err1 := db1.Model(user).Association("Tasks").Append(task)
	return err1
}
