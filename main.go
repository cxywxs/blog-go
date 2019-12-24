package main

import (
	"blog-go/controller"
	"github.com/iris-contrib/middleware/cors"
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
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	v1 := app.Party("/", crs).AllowMethods(iris.MethodOptions)
	{
		//登录post
		v1.Post("/login", controller.LoginController)
		//验证token post
		v1.Post("/gettoken", controller.VerifyToken)
		//注册
		v1.Post("/signup", controller.SignUpController)
		//写入博客内容
		v1.Post("/writeblog", controller.WriteBlogController)
		//时间排序
		v1.Get("/timesort", controller.GetBlogTimeSortController)
		//阅读量排序排序
		v1.Get("readnumsort", controller.GetblogReadNumsortController)
		//更新博客
	}

	//启动端口为8085的blog-go项目服务
	_ = app.Run(iris.Addr(":8085"))
}
