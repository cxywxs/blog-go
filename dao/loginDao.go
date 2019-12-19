package dao

import (
	"blog-go/entity"
	"github.com/jinzhu/gorm"
	"log"
)

func LoginDao(user entity.IrisUser, db *gorm.DB) bool {
	var user1 entity.IrisUser
	db.Where("username = ? and password = ?", user.Username, user.Password).First(&user1)
	if user1.Username == "" {
		log.Println(user.Username, user.Password, "数据库层面验证未通过")
		return false
	}
	log.Println(user.Username, user.Password, "数据库层面验证已通过通过")
	return true
}
