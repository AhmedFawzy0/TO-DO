package repos

import (
	"sort"

	"github.com/AhmedFawzy0/TO-DO/models"
	"github.com/alexedwards/argon2id"
	"gorm.io/gorm"
)

func LoginAuthorize(userTemp *models.User, user1 *models.UserDTO) bool {

	match, err := argon2id.ComparePasswordAndHash(user1.Password, userTemp.Password)

	if userTemp.Username == user1.Username && match && err == nil {
		return true
	}

	return false
}

func SortById(UserTasks []models.Task) []models.Task {

	sort.Slice(UserTasks, func(i, j int) bool {
		return UserTasks[i].ID < UserTasks[j].ID
	})

	return UserTasks
}

func FindUser(userTemp *models.User, user string, db1 *gorm.DB) (err error) {

	err1 := db1.Where("Username = ?", user).First(userTemp).Error

	return err1

}

func CreateUser(username string, password string, db1 *gorm.DB) (user *models.User, err error) {

	user1 := &models.User{Username: username, Password: password}
	err1 := db1.Create(user1).Error

	if err != nil {
		return nil, err1
	}

	return user1, err1

}
