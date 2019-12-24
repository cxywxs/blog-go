package controller

import (
	"blog-go/entity"
	"blog-go/tools"
	"encoding/json"
	"github.com/kataras/iris/v12"
	"log"
)

//验证token是否通过
func VerifyToken(ctx iris.Context) {
	body, err := ctx.GetBody()
	if err != nil {
		log.Println(err)
	}
	token := entity.Token{}
	_ = json.Unmarshal(body, &token)
	isVerify := tools.ValidToken(token.Tokenstring)
	if isVerify == true {

		_, err = ctx.Writef(tools.JsonSendOkString("验证成功"))
		if err != nil {
			log.Println(err)
		}
	} else {
		_, err = ctx.Writef(tools.JsonSendOkString("验证失败"))
		if err != nil {
			log.Println(err)
		}
	}
	log.Println("Post verify 完成")

}
