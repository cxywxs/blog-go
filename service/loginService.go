package service

import (
	"blog-go/dao"
	"blog-go/entity"
	"blog-go/tools"
)

func LoginServer(user entity.IrisUser) string {
	db := dao.UserDao()
	if dao.LoginDao(user, db) == true {
		return tools.CreateToken(user.Username, user.Password)
	}

	return "登录失败"
}
