package main

import (
	"blog-go/controller"
	"github.com/kataras/iris/v12"
	"log"
	"os"
)

func init() {
	file := "./" + "blog-go" + ".log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetFlags(log.LstdFlags)
	return
}

func main() {
	//启动iris默认配置
	app := iris.Default()
	//登录post
	app.Post("/login", controller.LoginController)
	//验证token post
	app.Post("/gettoken", controller.VerifyToken)

	app.Post("/signup", controller.SignUpController)

	app.Post("/writeblog", controller.WriteBlogController)

	app.Get("/timesort", controller.GetBlogTimeSortController)

	app.Get("readnumsort", controller.GetblogReadNumsortController)

	//启动端口为8085的blog-go项目服务
	app.Run(iris.Addr(":8085"))
}
