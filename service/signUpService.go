package service

import (
	"blog-go/dao"
	"blog-go/entity"
)

func SignUpService(user entity.IrisUser) string {
	db := dao.UserDao()
	createbool := dao.SignUpDao(user, db)
	if createbool != "" {
		if createbool == "same" {
			return "用户名重复"
		}
		return "插入数据错误"
	}
	return "创建成功"
}
