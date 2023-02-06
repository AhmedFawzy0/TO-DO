package models

import "gorm.io/gorm"

type Task struct 
{
	gorm.Model
	Finished  bool   `json:"finished"`
	Detail    string  `json:"detail"`
	UserID    uint
}

