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
	ctx.Writef(tools.JsonSendOkString("博客插入完成"))
}

func GetBlogTimeSortController(ctx iris.Context) {
	ctx.Writef(tools.JsonSendOkStruct(dao.GetArticleSort("time")))
}

func GetblogReadNumsortController(ctx iris.Context) {
	ctx.Writef(tools.JsonSendOkStruct(dao.GetArticleSort("readnum")))
}
