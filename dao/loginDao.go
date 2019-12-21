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

//返回值为空说明成功，same表示用户名重复，failed数据库插入错误
func SignUpDao(user entity.IrisUser, db *gorm.DB) string {
	var user1 entity.IrisUser
	db.Where("username = ? and password = ?", user.Username, user.Password).First(&user1)
	if user1.Username != "" {
		return "same"
	}
	db.Create(&user)
	db.Where("username = ? and password = ?", user.Username, user.Password).First(&user1)
	if user1.Username == user.Username {
		return ""
	}
	return "failed"
}
