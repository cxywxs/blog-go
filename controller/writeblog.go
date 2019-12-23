package controller

import (
	"blog-go/dao"
	"blog-go/entity"
	"blog-go/tools"
	"encoding/json"
	"github.com/kataras/iris/v12"
	"log"
)

func WriteBlogController(ctx iris.Context) {
	body, err := ctx.GetBody()
	if err != nil {
		log.Println(err)
	}
	article := entity.Article{}
	_ = json.Unmarshal(body, &article)
	article.Time = tools.GetTime()
	article.Readnum = 0
	dao.InsertArticle(article)
	_, err = ctx.Writef(tools.JsonSendOkString("博客插入完成"))
	if err != nil {
		log.Println(err)
	}
}

func GetBlogTimeSortController(ctx iris.Context) {
	_, err := ctx.Writef(tools.JsonSendOkStruct(dao.GetArticleSort("time")))
	if err != nil {
		log.Println(err)
	}
}

func GetblogReadNumsortController(ctx iris.Context) {
	_, err := ctx.Writef(tools.JsonSendOkStruct(dao.GetArticleSort("readnum")))
	if err != nil {
		log.Println(err)
	}
}
