package models

import "gorm.io/gorm"


type User struct {
	gorm.Model
	Username   string `json:"username" gorm:"text;not null;default:null"`
	Password   string `json:"-" gorm:"text;not null;default:null"`
	Tasks      []Task `json:"tasks" gorm:"foreignKey:UserID"`
}

type UserDTO struct {
	Username   string `json:"username" gorm:"text;not null;default:null"`
	Password   string `json:"password" gorm:"text;not null;default:null"`
}

type Task struct 
{
	gorm.Model
	Finished  bool   `json:"finished"`
	Detail    string  `json:"detail"`
	UserID    uint
}

