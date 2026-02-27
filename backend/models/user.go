package models

import "gorm.io/gorm"

type User struct {
	gorm.Model        //嵌入的gorm结构体
	Username   string `gorm:"unique"`
	Password   string
}
