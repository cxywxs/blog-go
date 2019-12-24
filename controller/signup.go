package controller

import (
	"blog-go/entity"
	"blog-go/service"
	"blog-go/tools"
	"encoding/json"
	"github.com/kataras/iris/v12"
	"log"
)

func SignUpController(ctx iris.Context) {
	body, err := ctx.GetBody()
	if err != nil {
		log.Println(err)
	}
	user := entity.IrisUser{}
	_ = json.Unmarshal(body, &user)
	user.Id = tools.CreateUuid()
	result := service.SignUpService(user)
	_, err = ctx.Writef(tools.JsonSendOkString(result))
	if err != nil {
		log.Println(err)
	}
	log.Println("Post login 完成")
}
