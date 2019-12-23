package dao

import (
	"blog-go/entity"
	"errors"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
)

//查询所有的文章
func GetAllArticle() []entity.Article {
	session, err := mgo.Dial("")
	if err != nil {
		log.Fatal("mongoDB连接失败")
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("cxytest")
	collection := db.C("article")
	var result []entity.Article
	iter := collection.Find(nil).Iter()
	_ = iter.All(&result)
	return result

}

//根据一个条件查询一篇文章
func GetArticle(condition entity.SelectCondition) entity.Article {
	session, err := mgo.Dial("")
	if err != nil {
		log.Fatal("mongoDB连接失败")
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("cxytest")
	collection := db.C("article")
	article := entity.Article{}
	_ = collection.Find(bson.M{condition.Field: bson.M{condition.Key: condition.Value}}).One(&article)
	return article
}

//根据单个条件查询出多篇文章
func GetArticles(condition entity.SelectCondition) []entity.Article {
	session, err := mgo.Dial("")
	if err != nil {
		log.Fatal("mongoDB连接失败")
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("cxytest")
	collection := db.C("article")
	var articles []entity.Article
	_ = collection.Find(bson.M{condition.Field: bson.M{condition.Key: condition.Value}}).All(&articles)
	return articles
}

//根据字段排序（-field为倒序）
func GetArticleSort(field string) []entity.Article {
	session, err := mgo.Dial("")
	if err != nil {
		log.Fatal("mongoDB连接失败")
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("cxytest")
	collection := db.C("article")
	var articles []entity.Article
	_ = collection.Find(nil).Sort(field).All(&articles)
	return articles
}

//插入博客
func InsertArticle(article entity.Article) {
	session, err := mgo.Dial("")
	if err != nil {
		log.Fatal("mongoDB连接失败")
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("cxytest")
	collection := db.C("article")
	err = collection.Insert(article)
	if err != nil {
		log.Fatalln(err)
	}

}

//更新博客
func UpdateArticle(article entity.Article) error {
	if article.Articleid == "" {
		return errors.New("文章id缺失")
	}
	session, err := mgo.Dial("")
	if err != nil {
		log.Fatal("mongoDB连接失败")
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("cxytest")
	collection := db.C("article")
	err = collection.Update(bson.M{"articleid": article.Articleid}, bson.M{"content": article.Content})
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}
