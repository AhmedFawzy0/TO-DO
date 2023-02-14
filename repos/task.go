package repos

import (
	"github.com/AhmedFawzy0/TO-DO/models"
	"gorm.io/gorm"
)

func FindUserTasks(user *models.User, usertask *[]models.Task, db1 *gorm.DB) (err error) {
	return db1.Model(user).Association("Tasks").Find(usertask)
}

func AddTask(user *models.User, task *models.Task, db1 *gorm.DB) (err error) {
	return db1.Model(user).Association("Tasks").Append(task)

}

func TaskDelete(task_model *models.Task, task_id uint, db1 *gorm.DB) (err error) {
	return db1.Delete(&task_model, task_id).Error
}

func UpdateTask(task_up *models.Task, db1 *gorm.DB) (err error) {
	return db1.Model(task_up).Select("Finished", "Detail").Where("id = ?", task_up.ID).Updates(models.Task{Finished: !task_up.Finished, Detail: task_up.Detail}).Error
}
