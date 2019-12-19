package controller

import (
	"blog-go/tools"
	"github.com/kataras/iris/v12"
	"log"
)

//验证token是否通过
func VerifyToken(ctx iris.Context) {
	token, err := ctx.GetBody()
	if err != nil {
		log.Println(err)
	}
	isVerify := tools.ValidToken(string(token))
	if isVerify == true {

		ctx.Writef(tools.JsonSendOkString("验证成功"))
	} else {
		ctx.Writef(tools.JsonSendOkString("验证失败"))
	}
	log.Println("Post verify 完成")

}